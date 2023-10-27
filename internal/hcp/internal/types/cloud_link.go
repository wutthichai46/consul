package types

import (
	"errors"

	"github.com/hashicorp/consul/internal/resource"
	pbhcp "github.com/hashicorp/consul/proto-public/pbhcp/v1alpha1"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

const (
	CloudLinkKind = "CloudLink"
)

var (
	linkConfigurationNameError = errors.New("Only a single HCP CloudLink resource is allowed and it must be named hcp-link")
)

func RegisterCloudLink(r resource.Registry) {
	r.Register(resource.Registration{
		Type:     pbhcp.CloudLinkType,
		Proto:    &pbhcp.CloudLink{},
		Scope:    resource.ScopeCluster,
		Mutate:   MutateCloudLink,
		Validate: ValidateCloudLink,
	})
}

func MutateCloudLink(res *pbresource.Resource) error {
	return nil
}

func ValidateCloudLink(res *pbresource.Resource) error {
	var link pbhcp.CloudLink

	if err := res.Data.UnmarshalTo(&link); err != nil {
		return resource.NewErrDataParse(&link, err)
	}

	if res.Id.Name != "hcp-link" {
		return resource.ErrInvalidField{
			Name:    "name",
			Wrapped: linkConfigurationNameError,
		}
	}

	return nil
}
