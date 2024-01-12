// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package resource

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"

	"github.com/hashicorp/consul/command/resource/client"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

const (
	HeaderConsulToken = "x-consul-token"
)

type ResourceGRPC struct {
	C *client.GRPCClient
}

func (resource *ResourceGRPC) Apply(parsedResource *pbresource.Resource) (*pbresource.Resource, error) {
	token, err := resource.C.Config.GetToken()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	if token != "" {
		ctx = metadata.AppendToOutgoingContext(context.Background(), HeaderConsulToken, token)
	}

	defer resource.C.Conn.Close()
	writeRsp, err := resource.C.Client.Write(ctx, &pbresource.WriteRequest{Resource: parsedResource})
	if err != nil {
		return nil, fmt.Errorf("error writing resource: %+v", err)
	}

	return writeRsp.Resource, err
}

func (resource *ResourceGRPC) Read(resourceType *pbresource.Type, resourceTenancy *pbresource.Tenancy, resourceName string, stale bool) (*pbresource.Resource, error) {
	token, err := resource.C.Config.GetToken()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	if !stale {
		ctx = metadata.AppendToOutgoingContext(ctx, "x-consul-consistency-mode", "consistent")
	}
	if token != "" {
		ctx = metadata.AppendToOutgoingContext(context.Background(), HeaderConsulToken, token)
	}

	defer resource.C.Conn.Close()
	readRsp, err := resource.C.Client.Read(ctx, &pbresource.ReadRequest{
		Id: &pbresource.ID{
			Type:    resourceType,
			Tenancy: resourceTenancy,
			Name:    resourceName,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("error writing resource: %+v", err)
	}

	return readRsp.Resource, err
}
