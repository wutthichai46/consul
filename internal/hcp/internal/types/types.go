package types

import "github.com/hashicorp/consul/internal/resource"

const (
	GroupName      = "hcp"
	VersionV1      = "v1"
	CurrentVersion = VersionV1
)

func Register(r resource.Registry) {
	RegisterCloudLink(r)
}
