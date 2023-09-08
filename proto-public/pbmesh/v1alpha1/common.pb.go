// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: pbmesh/v1alpha1/common.proto

package meshv1alpha1

import (
	pbresource "github.com/hashicorp/consul/proto-public/pbresource"
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

// NOTE: roughly equivalent to structs.ResourceReference
type ParentReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// For east/west configuration, this should point to a pbcatalog.Service.
	// For north/south it should point to a gateway (TBD)
	Ref *pbresource.Reference `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	// Port is the network port this Route targets. It can be interpreted
	// differently based on the type of parent resource.
	//
	// When the parent resource is a Gateway, this targets all listeners
	// listening on the specified port that also support this kind of Route(and
	// select this Route). It’s not recommended to set Port unless the networking
	// behaviors specified in a Route must apply to a specific port as opposed to
	// a listener(s) whose port(s) may be changed. When both Port and SectionName
	// are specified, the name and port of the selected listener must match both
	// specified values.
	//
	// Implementations MAY choose to support other parent resources.
	// Implementations supporting other types of parent resources MUST clearly
	// document how/if Port is interpreted.
	//
	// For the purpose of status, an attachment is considered successful as long
	// as the parent resource accepts it partially. For example, Gateway
	// listeners can restrict which Routes can attach to them by Route kind,
	// namespace, or hostname. If 1 of 2 Gateway listeners accept attachment from
	// the referencing Route, the Route MUST be considered successfully attached.
	// If no Gateway listeners accept attachment from this Route, the Route MUST
	// be considered detached from the Gateway.
	//
	// For east/west this is the name of the consul port.
	// For north/south this is the stringified integer port expected by GAMMA.
	//
	// https://gateway-api.sigs.k8s.io/geps/gep-957/
	Port string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *ParentReference) Reset() {
	*x = ParentReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v1alpha1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParentReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParentReference) ProtoMessage() {}

func (x *ParentReference) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v1alpha1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParentReference.ProtoReflect.Descriptor instead.
func (*ParentReference) Descriptor() ([]byte, []int) {
	return file_pbmesh_v1alpha1_common_proto_rawDescGZIP(), []int{0}
}

func (x *ParentReference) GetRef() *pbresource.Reference {
	if x != nil {
		return x.Ref
	}
	return nil
}

func (x *ParentReference) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

type BackendReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// For east/west configuration, this should point to either a
	// pbcatalog.Service or ServiceSubset.
	//
	// For Partition/PeerName fields likely we could map them to ServiceImports
	// (MCS+GAMMA) when translating
	Ref *pbresource.Reference `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	// For east/west this is the name of the consul port.
	// For north/south this is the stringified integer port expected by GAMMA.
	Port string `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	// NOT IN GAMMA; multi-cluster + GWapi is still unknown
	//
	// Likely we could map this to ServiceImports (MCS+GAMMA) when translating
	// to/from k8s.
	//
	// https://gateway-api.sigs.k8s.io/geps/gep-1748/
	Datacenter string `protobuf:"bytes,3,opt,name=datacenter,proto3" json:"datacenter,omitempty"`
}

func (x *BackendReference) Reset() {
	*x = BackendReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v1alpha1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackendReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackendReference) ProtoMessage() {}

func (x *BackendReference) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v1alpha1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackendReference.ProtoReflect.Descriptor instead.
func (*BackendReference) Descriptor() ([]byte, []int) {
	return file_pbmesh_v1alpha1_common_proto_rawDescGZIP(), []int{1}
}

func (x *BackendReference) GetRef() *pbresource.Reference {
	if x != nil {
		return x.Ref
	}
	return nil
}

func (x *BackendReference) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *BackendReference) GetDatacenter() string {
	if x != nil {
		return x.Datacenter
	}
	return ""
}

var File_pbmesh_v1alpha1_common_proto protoreflect.FileDescriptor

var file_pbmesh_v1alpha1_common_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x19,
	0x70, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x0f, 0x50, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x03,
	0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x03, 0x72, 0x65, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x7e, 0x0a, 0x10, 0x42, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x03,
	0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x03, 0x72, 0x65, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x42, 0x93, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d,
	0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42,
	0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x6d, 0x65, 0x73, 0x68, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x43, 0x4d, 0xaa, 0x02, 0x1e, 0x48, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x4d,
	0x65, 0x73, 0x68, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x1e, 0x48,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c,
	0x4d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x2a,
	0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x5c, 0x4d, 0x65, 0x73, 0x68, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x21, 0x48, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x3a, 0x3a,
	0x4d, 0x65, 0x73, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbmesh_v1alpha1_common_proto_rawDescOnce sync.Once
	file_pbmesh_v1alpha1_common_proto_rawDescData = file_pbmesh_v1alpha1_common_proto_rawDesc
)

func file_pbmesh_v1alpha1_common_proto_rawDescGZIP() []byte {
	file_pbmesh_v1alpha1_common_proto_rawDescOnce.Do(func() {
		file_pbmesh_v1alpha1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbmesh_v1alpha1_common_proto_rawDescData)
	})
	return file_pbmesh_v1alpha1_common_proto_rawDescData
}

var file_pbmesh_v1alpha1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pbmesh_v1alpha1_common_proto_goTypes = []interface{}{
	(*ParentReference)(nil),      // 0: hashicorp.consul.mesh.v1alpha1.ParentReference
	(*BackendReference)(nil),     // 1: hashicorp.consul.mesh.v1alpha1.BackendReference
	(*pbresource.Reference)(nil), // 2: hashicorp.consul.resource.Reference
}
var file_pbmesh_v1alpha1_common_proto_depIdxs = []int32{
	2, // 0: hashicorp.consul.mesh.v1alpha1.ParentReference.ref:type_name -> hashicorp.consul.resource.Reference
	2, // 1: hashicorp.consul.mesh.v1alpha1.BackendReference.ref:type_name -> hashicorp.consul.resource.Reference
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pbmesh_v1alpha1_common_proto_init() }
func file_pbmesh_v1alpha1_common_proto_init() {
	if File_pbmesh_v1alpha1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbmesh_v1alpha1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParentReference); i {
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
		file_pbmesh_v1alpha1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackendReference); i {
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
			RawDescriptor: file_pbmesh_v1alpha1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbmesh_v1alpha1_common_proto_goTypes,
		DependencyIndexes: file_pbmesh_v1alpha1_common_proto_depIdxs,
		MessageInfos:      file_pbmesh_v1alpha1_common_proto_msgTypes,
	}.Build()
	File_pbmesh_v1alpha1_common_proto = out.File
	file_pbmesh_v1alpha1_common_proto_rawDesc = nil
	file_pbmesh_v1alpha1_common_proto_goTypes = nil
	file_pbmesh_v1alpha1_common_proto_depIdxs = nil
}
