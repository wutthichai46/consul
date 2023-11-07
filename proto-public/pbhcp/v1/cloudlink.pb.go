// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: pbhcp/v1/cloudlink.proto

package hcpv1

import (
	_ "github.com/hashicorp/consul/proto-public/pbresource"
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

type Link_ACCESS_LEVEL int32

const (
	Link_ACCESS_LEVEL_UNSPECIFIED       Link_ACCESS_LEVEL = 0
	Link_ACCESS_LEVEL_GLOBAL_READ_ONLY  Link_ACCESS_LEVEL = 1
	Link_ACCESS_LEVEL_GLOBAL_READ_WRITE Link_ACCESS_LEVEL = 2
)

// Enum value maps for Link_ACCESS_LEVEL.
var (
	Link_ACCESS_LEVEL_name = map[int32]string{
		0: "ACCESS_LEVEL_UNSPECIFIED",
		1: "ACCESS_LEVEL_GLOBAL_READ_ONLY",
		2: "ACCESS_LEVEL_GLOBAL_READ_WRITE",
	}
	Link_ACCESS_LEVEL_value = map[string]int32{
		"ACCESS_LEVEL_UNSPECIFIED":       0,
		"ACCESS_LEVEL_GLOBAL_READ_ONLY":  1,
		"ACCESS_LEVEL_GLOBAL_READ_WRITE": 2,
	}
)

func (x Link_ACCESS_LEVEL) Enum() *Link_ACCESS_LEVEL {
	p := new(Link_ACCESS_LEVEL)
	*p = x
	return p
}

func (x Link_ACCESS_LEVEL) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Link_ACCESS_LEVEL) Descriptor() protoreflect.EnumDescriptor {
	return file_pbhcp_v1_cloudlink_proto_enumTypes[0].Descriptor()
}

func (Link_ACCESS_LEVEL) Type() protoreflect.EnumType {
	return &file_pbhcp_v1_cloudlink_proto_enumTypes[0]
}

func (x Link_ACCESS_LEVEL) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Link_ACCESS_LEVEL.Descriptor instead.
func (Link_ACCESS_LEVEL) EnumDescriptor() ([]byte, []int) {
	return file_pbhcp_v1_cloudlink_proto_rawDescGZIP(), []int{0, 0}
}

type LinkConfig_TLS int32

const (
	LinkConfig_ENABLED  LinkConfig_TLS = 0
	LinkConfig_INSECURE LinkConfig_TLS = 1
	LinkConfig_DISABLED LinkConfig_TLS = 2
)

// Enum value maps for LinkConfig_TLS.
var (
	LinkConfig_TLS_name = map[int32]string{
		0: "ENABLED",
		1: "INSECURE",
		2: "DISABLED",
	}
	LinkConfig_TLS_value = map[string]int32{
		"ENABLED":  0,
		"INSECURE": 1,
		"DISABLED": 2,
	}
)

func (x LinkConfig_TLS) Enum() *LinkConfig_TLS {
	p := new(LinkConfig_TLS)
	*p = x
	return p
}

func (x LinkConfig_TLS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LinkConfig_TLS) Descriptor() protoreflect.EnumDescriptor {
	return file_pbhcp_v1_cloudlink_proto_enumTypes[1].Descriptor()
}

func (LinkConfig_TLS) Type() protoreflect.EnumType {
	return &file_pbhcp_v1_cloudlink_proto_enumTypes[1]
}

func (x LinkConfig_TLS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LinkConfig_TLS.Descriptor instead.
func (LinkConfig_TLS) EnumDescriptor() ([]byte, []int) {
	return file_pbhcp_v1_cloudlink_proto_rawDescGZIP(), []int{1, 0}
}

type Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required
	ResourceId   string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	ClientId     string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret string `protobuf:"bytes,3,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	// Optional, testing only really
	AuthUrl      string `protobuf:"bytes,4,opt,name=auth_url,json=authUrl,proto3" json:"auth_url,omitempty"`
	ApiAddress   string `protobuf:"bytes,5,opt,name=api_address,json=apiAddress,proto3" json:"api_address,omitempty"`
	ScadaAddress string `protobuf:"bytes,6,opt,name=scada_address,json=scadaAddress,proto3" json:"scada_address,omitempty"`
	PortalUrl    string `protobuf:"bytes,7,opt,name=portal_url,json=portalUrl,proto3" json:"portal_url,omitempty"`
	// Info from GNM, can't be set by user
	AccessLevel Link_ACCESS_LEVEL `protobuf:"varint,8,opt,name=access_level,json=accessLevel,proto3,enum=hashicorp.consul.hcp.v1.Link_ACCESS_LEVEL" json:"access_level,omitempty"`
}

func (x *Link) Reset() {
	*x = Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbhcp_v1_cloudlink_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_pbhcp_v1_cloudlink_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_pbhcp_v1_cloudlink_proto_rawDescGZIP(), []int{0}
}

func (x *Link) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *Link) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Link) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *Link) GetAuthUrl() string {
	if x != nil {
		return x.AuthUrl
	}
	return ""
}

func (x *Link) GetApiAddress() string {
	if x != nil {
		return x.ApiAddress
	}
	return ""
}

func (x *Link) GetScadaAddress() string {
	if x != nil {
		return x.ScadaAddress
	}
	return ""
}

func (x *Link) GetPortalUrl() string {
	if x != nil {
		return x.PortalUrl
	}
	return ""
}

func (x *Link) GetAccessLevel() Link_ACCESS_LEVEL {
	if x != nil {
		return x.AccessLevel
	}
	return Link_ACCESS_LEVEL_UNSPECIFIED
}

type LinkConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId     string         `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret string         `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	AuthUrl      string         `protobuf:"bytes,3,opt,name=auth_url,json=authUrl,proto3" json:"auth_url,omitempty"`
	AuthTls      LinkConfig_TLS `protobuf:"varint,4,opt,name=auth_tls,json=authTls,proto3,enum=hashicorp.consul.hcp.v1.LinkConfig_TLS" json:"auth_tls,omitempty"`
	ApiAddress   string         `protobuf:"bytes,5,opt,name=api_address,json=apiAddress,proto3" json:"api_address,omitempty"`
	ApiTls       LinkConfig_TLS `protobuf:"varint,6,opt,name=api_tls,json=apiTls,proto3,enum=hashicorp.consul.hcp.v1.LinkConfig_TLS" json:"api_tls,omitempty"`
	ScadaAddress string         `protobuf:"bytes,7,opt,name=scada_address,json=scadaAddress,proto3" json:"scada_address,omitempty"`
	ScadaTls     LinkConfig_TLS `protobuf:"varint,8,opt,name=scada_tls,json=scadaTls,proto3,enum=hashicorp.consul.hcp.v1.LinkConfig_TLS" json:"scada_tls,omitempty"`
	PortalUrl    string         `protobuf:"bytes,9,opt,name=portal_url,json=portalUrl,proto3" json:"portal_url,omitempty"`
}

func (x *LinkConfig) Reset() {
	*x = LinkConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbhcp_v1_cloudlink_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkConfig) ProtoMessage() {}

func (x *LinkConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pbhcp_v1_cloudlink_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkConfig.ProtoReflect.Descriptor instead.
func (*LinkConfig) Descriptor() ([]byte, []int) {
	return file_pbhcp_v1_cloudlink_proto_rawDescGZIP(), []int{1}
}

func (x *LinkConfig) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *LinkConfig) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *LinkConfig) GetAuthUrl() string {
	if x != nil {
		return x.AuthUrl
	}
	return ""
}

func (x *LinkConfig) GetAuthTls() LinkConfig_TLS {
	if x != nil {
		return x.AuthTls
	}
	return LinkConfig_ENABLED
}

func (x *LinkConfig) GetApiAddress() string {
	if x != nil {
		return x.ApiAddress
	}
	return ""
}

func (x *LinkConfig) GetApiTls() LinkConfig_TLS {
	if x != nil {
		return x.ApiTls
	}
	return LinkConfig_ENABLED
}

func (x *LinkConfig) GetScadaAddress() string {
	if x != nil {
		return x.ScadaAddress
	}
	return ""
}

func (x *LinkConfig) GetScadaTls() LinkConfig_TLS {
	if x != nil {
		return x.ScadaTls
	}
	return LinkConfig_ENABLED
}

func (x *LinkConfig) GetPortalUrl() string {
	if x != nil {
		return x.PortalUrl
	}
	return ""
}

var File_pbhcp_v1_cloudlink_proto protoreflect.FileDescriptor

var file_pbhcp_v1_cloudlink_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x62, 0x68, 0x63, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x68, 0x63, 0x70,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x70, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb5, 0x03, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x75, 0x74, 0x68, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61,
	0x70, 0x69, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x63, 0x61,
	0x64, 0x61, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x63, 0x61, 0x64, 0x61, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x4d, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x68, 0x63, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x52,
	0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x73, 0x0a, 0x0c,
	0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x12, 0x1c, 0x0a, 0x18,
	0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x41, 0x43,
	0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x47, 0x4c, 0x4f, 0x42, 0x41,
	0x4c, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x01, 0x12, 0x22, 0x0a,
	0x1e, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x47, 0x4c,
	0x4f, 0x42, 0x41, 0x4c, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10,
	0x02, 0x3a, 0x06, 0xa2, 0x93, 0x04, 0x02, 0x08, 0x01, 0x22, 0xca, 0x03, 0x0a, 0x0a, 0x4c, 0x69,
	0x6e, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x75,
	0x74, 0x68, 0x55, 0x72, 0x6c, 0x12, 0x42, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6c,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x68, 0x63, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x4c, 0x53,
	0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6c, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x61, 0x70, 0x69, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x40, 0x0a, 0x07, 0x61, 0x70,
	0x69, 0x5f, 0x74, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x68, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x68,
	0x63, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x54, 0x4c, 0x53, 0x52, 0x06, 0x61, 0x70, 0x69, 0x54, 0x6c, 0x73, 0x12, 0x23, 0x0a, 0x0d,
	0x73, 0x63, 0x61, 0x64, 0x61, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x63, 0x61, 0x64, 0x61, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x44, 0x0a, 0x09, 0x73, 0x63, 0x61, 0x64, 0x61, 0x5f, 0x74, 0x6c, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x68, 0x63, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x6e, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x4c, 0x53, 0x52, 0x08, 0x73,
	0x63, 0x61, 0x64, 0x61, 0x54, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0x2e, 0x0a, 0x03, 0x54, 0x4c, 0x53, 0x12, 0x0b, 0x0a,
	0x07, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e,
	0x53, 0x45, 0x43, 0x55, 0x52, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41,
	0x42, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x42, 0xe5, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e,
	0x68, 0x63, 0x70, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x6c, 0x69, 0x6e,
	0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x2f, 0x70, 0x62, 0x68, 0x63, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x68, 0x63, 0x70, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x48, 0x43, 0x48, 0xaa, 0x02, 0x17, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x48, 0x63, 0x70, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x17, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x48, 0x63, 0x70, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x23, 0x48, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x48,
	0x63, 0x70, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x1a, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x3a, 0x3a, 0x48, 0x63, 0x70, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbhcp_v1_cloudlink_proto_rawDescOnce sync.Once
	file_pbhcp_v1_cloudlink_proto_rawDescData = file_pbhcp_v1_cloudlink_proto_rawDesc
)

func file_pbhcp_v1_cloudlink_proto_rawDescGZIP() []byte {
	file_pbhcp_v1_cloudlink_proto_rawDescOnce.Do(func() {
		file_pbhcp_v1_cloudlink_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbhcp_v1_cloudlink_proto_rawDescData)
	})
	return file_pbhcp_v1_cloudlink_proto_rawDescData
}

var file_pbhcp_v1_cloudlink_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pbhcp_v1_cloudlink_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pbhcp_v1_cloudlink_proto_goTypes = []interface{}{
	(Link_ACCESS_LEVEL)(0), // 0: hashicorp.consul.hcp.v1.Link.ACCESS_LEVEL
	(LinkConfig_TLS)(0),    // 1: hashicorp.consul.hcp.v1.LinkConfig.TLS
	(*Link)(nil),           // 2: hashicorp.consul.hcp.v1.Link
	(*LinkConfig)(nil),     // 3: hashicorp.consul.hcp.v1.LinkConfig
}
var file_pbhcp_v1_cloudlink_proto_depIdxs = []int32{
	0, // 0: hashicorp.consul.hcp.v1.Link.access_level:type_name -> hashicorp.consul.hcp.v1.Link.ACCESS_LEVEL
	1, // 1: hashicorp.consul.hcp.v1.LinkConfig.auth_tls:type_name -> hashicorp.consul.hcp.v1.LinkConfig.TLS
	1, // 2: hashicorp.consul.hcp.v1.LinkConfig.api_tls:type_name -> hashicorp.consul.hcp.v1.LinkConfig.TLS
	1, // 3: hashicorp.consul.hcp.v1.LinkConfig.scada_tls:type_name -> hashicorp.consul.hcp.v1.LinkConfig.TLS
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pbhcp_v1_cloudlink_proto_init() }
func file_pbhcp_v1_cloudlink_proto_init() {
	if File_pbhcp_v1_cloudlink_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbhcp_v1_cloudlink_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Link); i {
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
		file_pbhcp_v1_cloudlink_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkConfig); i {
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
			RawDescriptor: file_pbhcp_v1_cloudlink_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbhcp_v1_cloudlink_proto_goTypes,
		DependencyIndexes: file_pbhcp_v1_cloudlink_proto_depIdxs,
		EnumInfos:         file_pbhcp_v1_cloudlink_proto_enumTypes,
		MessageInfos:      file_pbhcp_v1_cloudlink_proto_msgTypes,
	}.Build()
	File_pbhcp_v1_cloudlink_proto = out.File
	file_pbhcp_v1_cloudlink_proto_rawDesc = nil
	file_pbhcp_v1_cloudlink_proto_goTypes = nil
	file_pbhcp_v1_cloudlink_proto_depIdxs = nil
}