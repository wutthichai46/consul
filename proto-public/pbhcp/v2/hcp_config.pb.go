// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: pbhcp/v2/hcp_config.proto

package hcpv2

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// HCPConfig is used to configure the HCP SDK for communicating with
// the HashiCorp Cloud Platform
type HCPConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthUrl               string `protobuf:"bytes,1,opt,name=auth_url,json=authUrl,proto3" json:"auth_url,omitempty"`
	ApiAddress            string `protobuf:"bytes,2,opt,name=api_address,json=apiAddress,proto3" json:"api_address,omitempty"`
	ScadaAddress          string `protobuf:"bytes,3,opt,name=scada_address,json=scadaAddress,proto3" json:"scada_address,omitempty"`
	TlsInsecureSkipVerify bool   `protobuf:"varint,4,opt,name=tls_insecure_skip_verify,json=tlsInsecureSkipVerify,proto3" json:"tls_insecure_skip_verify,omitempty"`
}

func (x *HCPConfig) Reset() {
	*x = HCPConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbhcp_v2_hcp_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HCPConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HCPConfig) ProtoMessage() {}

func (x *HCPConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pbhcp_v2_hcp_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HCPConfig.ProtoReflect.Descriptor instead.
func (*HCPConfig) Descriptor() ([]byte, []int) {
	return file_pbhcp_v2_hcp_config_proto_rawDescGZIP(), []int{0}
}

func (x *HCPConfig) GetAuthUrl() string {
	if x != nil {
		return x.AuthUrl
	}
	return ""
}

func (x *HCPConfig) GetApiAddress() string {
	if x != nil {
		return x.ApiAddress
	}
	return ""
}

func (x *HCPConfig) GetScadaAddress() string {
	if x != nil {
		return x.ScadaAddress
	}
	return ""
}

func (x *HCPConfig) GetTlsInsecureSkipVerify() bool {
	if x != nil {
		return x.TlsInsecureSkipVerify
	}
	return false
}

var File_pbhcp_v2_hcp_config_proto protoreflect.FileDescriptor

var file_pbhcp_v2_hcp_config_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x62, 0x68, 0x63, 0x70, 0x2f, 0x76, 0x32, 0x2f, 0x68, 0x63, 0x70, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x68, 0x63,
	0x70, 0x2e, 0x76, 0x32, 0x22, 0xa5, 0x01, 0x0a, 0x09, 0x48, 0x43, 0x50, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a,
	0x0b, 0x61, 0x70, 0x69, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x23,
	0x0a, 0x0d, 0x73, 0x63, 0x61, 0x64, 0x61, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x63, 0x61, 0x64, 0x61, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x37, 0x0a, 0x18, 0x74, 0x6c, 0x73, 0x5f, 0x69, 0x6e, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x65, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x74, 0x6c, 0x73, 0x49, 0x6e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x65, 0x53, 0x6b, 0x69, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x42, 0xe5, 0x01, 0x0a,
	0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x68, 0x63, 0x70, 0x2e, 0x76, 0x32, 0x42, 0x0e, 0x48, 0x63,
	0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x62, 0x68, 0x63, 0x70, 0x2f, 0x76,
	0x32, 0x3b, 0x68, 0x63, 0x70, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x48, 0x43, 0x48, 0xaa, 0x02, 0x17,
	0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x2e, 0x48, 0x63, 0x70, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x17, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x48, 0x63, 0x70, 0x5c, 0x56,
	0x32, 0xe2, 0x02, 0x23, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x48, 0x63, 0x70, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x3a, 0x3a, 0x48, 0x63, 0x70,
	0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbhcp_v2_hcp_config_proto_rawDescOnce sync.Once
	file_pbhcp_v2_hcp_config_proto_rawDescData = file_pbhcp_v2_hcp_config_proto_rawDesc
)

func file_pbhcp_v2_hcp_config_proto_rawDescGZIP() []byte {
	file_pbhcp_v2_hcp_config_proto_rawDescOnce.Do(func() {
		file_pbhcp_v2_hcp_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbhcp_v2_hcp_config_proto_rawDescData)
	})
	return file_pbhcp_v2_hcp_config_proto_rawDescData
}

var file_pbhcp_v2_hcp_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pbhcp_v2_hcp_config_proto_goTypes = []interface{}{
	(*HCPConfig)(nil), // 0: hashicorp.consul.hcp.v2.HCPConfig
}
var file_pbhcp_v2_hcp_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pbhcp_v2_hcp_config_proto_init() }
func file_pbhcp_v2_hcp_config_proto_init() {
	if File_pbhcp_v2_hcp_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbhcp_v2_hcp_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HCPConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pbhcp_v2_hcp_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbhcp_v2_hcp_config_proto_goTypes,
		DependencyIndexes: file_pbhcp_v2_hcp_config_proto_depIdxs,
		MessageInfos:      file_pbhcp_v2_hcp_config_proto_msgTypes,
	}.Build()
	File_pbhcp_v2_hcp_config_proto = out.File
	file_pbhcp_v2_hcp_config_proto_rawDesc = nil
	file_pbhcp_v2_hcp_config_proto_goTypes = nil
	file_pbhcp_v2_hcp_config_proto_depIdxs = nil
}
