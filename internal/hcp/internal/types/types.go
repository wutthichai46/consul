package types

import "github.com/hashicorp/consul/internal/resource"

const (
	GroupName       = "hcp"
	VersionV1Alpha1 = "v1alpha1"
	CurrentVersion  = VersionV1Alpha1
)

func Register(r resource.Registry) {
	RegisterCloudLink(r)
}
