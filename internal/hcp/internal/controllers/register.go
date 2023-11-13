package controllers

import (
	"github.com/hashicorp/consul/agent/hcp"
	"github.com/hashicorp/consul/internal/controller"
	"github.com/hashicorp/consul/internal/hcp/internal/controllers/cloudlink"
)

type Dependencies struct {
	Manager    *hcp.Manager
	Configured bool
}

func Register(mgr *controller.Manager, deps Dependencies) {
	mgr.Register(cloudlink.HCPLinkController(
		deps.Manager,
		deps.Configured,
	))
}
