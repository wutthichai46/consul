// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent
// +build !consulent

package types

import (
	"github.com/hashicorp/consul/internal/resource/resourcetest"
	multiclusterv1alpha1 "github.com/hashicorp/consul/proto-public/pbmulticluster/v1alpha1"
	"github.com/hashicorp/consul/proto-public/pbresource"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExportedServicesValidation(t *testing.T) {
	type testcase struct {
		Resource *pbresource.Resource
	}
	run := func(t *testing.T, tc testcase) {
		err := MutateExportedServices(tc.Resource)
		require.NoError(t, err)

		err = ValidateExportedServices(tc.Resource)
		require.NoError(t, err)
	}

	cases := map[string]testcase{
		"exported services with peer": {
			Resource: resourcetest.Resource(multiclusterv1alpha1.ExportedServicesType, "exported-services").
				WithData(t, validExportedServicesWithPeer()).
				Build(),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			run(t, tc)
		})
	}
}

func TestExportedServicesValidation_Error(t *testing.T) {
	type testcase struct {
		Resource *pbresource.Resource
	}
	run := func(t *testing.T, tc testcase) {
		err := MutateExportedServices(tc.Resource)
		require.NoError(t, err)

		err = ValidateExportedServices(tc.Resource)
		require.Error(t, err)
	}

	cases := map[string]testcase{
		"exported services with partition": {
			Resource: resourcetest.Resource(multiclusterv1alpha1.ExportedServicesType, "exported-services").
				WithData(t, validExportedServicesWithPartition()).
				Build(),
		},
		"exported services with sameness_group": {
			Resource: resourcetest.Resource(multiclusterv1alpha1.ExportedServicesType, "exported-services").
				WithData(t, validExportedServicesWithSamenessGroup()).
				Build(),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			run(t, tc)
		})
	}
}
