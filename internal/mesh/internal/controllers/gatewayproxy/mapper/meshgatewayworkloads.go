package mapper

import (
	"context"

	"github.com/hashicorp/consul/internal/controller"
	"github.com/hashicorp/consul/internal/mesh/internal/controllers/gatewayproxy/fetcher"
	"github.com/hashicorp/consul/internal/resource"
	pbcatalog "github.com/hashicorp/consul/proto-public/pbcatalog/v2beta1"
	pbmesh "github.com/hashicorp/consul/proto-public/pbmesh/v2beta1"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

var _ controller.DependencyMapper = AllMeshGatewayWorkloads

// AllMeshGatewayWorkloads returns one controller.Request for each Workload
// selected by a MeshGateway.
var AllMeshGatewayWorkloads = func(ctx context.Context, rt controller.Runtime, _ *pbresource.Resource) ([]controller.Request, error) {
	fetcher := fetcher.New(rt.Client, nil)

	gateways, err := fetcher.FetchMeshGateways(ctx)
	if err != nil {
		return nil, err
	}

	var requests []controller.Request

	for _, gateway := range gateways {
		endpointsID := resource.ReplaceType(pbcatalog.ServiceEndpointsType, gateway.Id)

		endpoints, err := fetcher.FetchServiceEndpoints(ctx, endpointsID)
		if err != nil {
			continue
		}

		for _, endpoint := range endpoints.Data.Endpoints {
			requests = append(requests, controller.Request{
				ID: resource.ReplaceType(pbmesh.ProxyStateTemplateType, endpoint.TargetRef),
			})
		}
	}

	return requests, nil
}
