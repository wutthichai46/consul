// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package hcpv1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using Link within kubernetes types, where deepcopy-gen is used.
func (in *Link) DeepCopyInto(out *Link) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Link. Required by controller-gen.
func (in *Link) DeepCopy() *Link {
	if in == nil {
		return nil
	}
	out := new(Link)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Link. Required by controller-gen.
func (in *Link) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using LinkConfig within kubernetes types, where deepcopy-gen is used.
func (in *LinkConfig) DeepCopyInto(out *LinkConfig) {
	proto.Reset(out)
	proto.Merge(out, proto.Clone(in))
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinkConfig. Required by controller-gen.
func (in *LinkConfig) DeepCopy() *LinkConfig {
	if in == nil {
		return nil
	}
	out := new(LinkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new LinkConfig. Required by controller-gen.
func (in *LinkConfig) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}