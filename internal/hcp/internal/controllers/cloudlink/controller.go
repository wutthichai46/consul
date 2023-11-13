package cloudlink

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/hashicorp/consul/agent/hcp"
	"github.com/hashicorp/consul/internal/controller"
	"github.com/hashicorp/consul/internal/resource"
	"github.com/hashicorp/consul/internal/storage"
	pbhcp "github.com/hashicorp/consul/proto-public/pbhcp/v1"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

const (
	StatusKey = "consul.io/hcp-link"
)

func HCPLinkController(manager *hcp.Manager, configured bool) controller.Controller {
	return controller.ForType(pbhcp.LinkType).
		WithReconciler(&hcpLinkReconciler{
			manager: manager,
		}).WithInitializer(&hcpLinkInitializer{
		manager:    manager,
		configured: configured,
	})
}

type hcpLinkReconciler struct {
	manager *hcp.Manager
}

func (r *hcpLinkReconciler) Reconcile(ctx context.Context, rt controller.Runtime, req controller.Request) error {
	// The runtime is passed by value so replacing it here for the remainder of this
	// reconciliation request processing will not affect future invocations.
	rt.Logger = rt.Logger.With("resource-id", req.ID, "controller", StatusKey)

	rt.Logger.Trace("reconciling link")

	// read the workload
	rsp, err := rt.Client.Read(ctx, &pbresource.ReadRequest{Id: req.ID})
	switch {
	case status.Code(err) == codes.NotFound:
		rt.Logger.Trace("cloud link has been deleted")
		// Shutdown Manager
		return nil
	case err != nil:
		rt.Logger.Error("the resource service has returned an unexpected error", "error", err)
		return err
	}

	res := rsp.Resource
	var cl pbhcp.Link
	if err := res.Data.UnmarshalTo(&cl); err != nil {
		rt.Logger.Error("error unmarshalling link data", "error", err)
		return err
	}

	// TODO: Start or restart manager with link config.

	// Set link status to successful
	conditions := []*pbresource.Condition{
		{
			Type:    "linked",
			State:   pbresource.Condition_STATE_TRUE,
			Reason:  "Success",
			Message: fmt.Sprintf("Successfully linked to '%s'", cl.ResourceId),
		},
	}
	newStatus := &pbresource.Status{
		ObservedGeneration: res.Generation,
		Conditions:         conditions,
	}

	if resource.EqualStatus(res.Status[StatusKey], newStatus, false) {
		return nil
	}
	_, err = rt.Client.WriteStatus(ctx, &pbresource.WriteStatusRequest{
		Id:     res.Id,
		Key:    StatusKey,
		Status: newStatus,
	})

	if err != nil {
		return err
	}

	return nil
}

type hcpLinkInitializer struct {
	manager    *hcp.Manager
	configured bool
}

func (i hcpLinkInitializer) Initialize(ctx context.Context, rt controller.Runtime) error {
	if !i.configured {
		return nil
	}

	// Create the link resource to reflect the configuration
	data, err := anypb.New(&pbhcp.Link{
		// TODO: Determine the configured resource ID, client ID, and client secret
	})
	if err != nil {
		return err
	}
	_, err = rt.Client.Write(ctx,
		&pbresource.WriteRequest{
			Resource: &pbresource.Resource{
				Id: &pbresource.ID{
					Name: "default",
					Type: pbhcp.LinkType,
				},
				Metadata: map[string]string{
					"source": "config",
				},
				Data: data,
			},
		},
	)
	if err != nil {
		if strings.Contains(err.Error(), storage.ErrWrongUid.Error()) {
			// Ignore wrong UID errors, which indicates the link already exists
			return nil
		}
		return err
	}

	return nil
}
