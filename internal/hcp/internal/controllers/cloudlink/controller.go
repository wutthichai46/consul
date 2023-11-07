package cloudlink

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hashicorp/consul/agent/hcp"
	"github.com/hashicorp/consul/internal/controller"
	"github.com/hashicorp/consul/internal/resource"
	pbhcp "github.com/hashicorp/consul/proto-public/pbhcp/v1"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

const (
	StatusKey = "consul.io/hcp-link"
)

func HCPLinkController(manager *hcp.Manager) controller.Controller {
	return controller.ForType(pbhcp.LinkType).
		WithReconciler(&hcpLinkReconciler{
			manager: manager,
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

	conditions := []*pbresource.Condition{
		{
			Type:    "initiated",
			State:   pbresource.Condition_STATE_TRUE,
			Reason:  "LINK_INITIATED",
			Message: fmt.Sprintf("Link has been initiated to '%s'", cl.ResourceId),
		},
		{
			Type:    "linked",
			State:   pbresource.Condition_STATE_TRUE,
			Reason:  "LINK_SUCCESSFUL",
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

	// Start or restart manager with link config.
	return nil
}
