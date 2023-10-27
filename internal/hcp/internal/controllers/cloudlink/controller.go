package cloudlink

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hashicorp/consul/agent/hcp"
	"github.com/hashicorp/consul/internal/controller"
	pbhcp "github.com/hashicorp/consul/proto-public/pbhcp/v1alpha1"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

const (
	StatusKey = "consul.io/cloudlink"
)

func HCPCloudLinkController(manager *hcp.Manager) controller.Controller {
	return controller.ForType(pbhcp.CloudLinkType).
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

	rt.Logger.Trace("reconciling cloud link")

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
	var cl pbhcp.CloudLink
	if err := res.Data.UnmarshalTo(&cl); err != nil {
		rt.Logger.Error("error unmarshalling cloud link data", "error", err)
		return err
	}

	// Start or restart manager with cloudlink config.
	return nil
}
