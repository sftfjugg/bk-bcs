// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: feed_server.proto

package pbfs

import (
	base "bscp.io/pkg/protocol/core/base"
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

// HandshakeMessage defines the handshake message from sidecar to feed server.
type HandshakeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiVersion *base.Versioning `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	Spec       *SidecarSpec     `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *HandshakeMessage) Reset() {
	*x = HandshakeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandshakeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandshakeMessage) ProtoMessage() {}

func (x *HandshakeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_feed_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandshakeMessage.ProtoReflect.Descriptor instead.
func (*HandshakeMessage) Descriptor() ([]byte, []int) {
	return file_feed_server_proto_rawDescGZIP(), []int{0}
}

func (x *HandshakeMessage) GetApiVersion() *base.Versioning {
	if x != nil {
		return x.ApiVersion
	}
	return nil
}

func (x *HandshakeMessage) GetSpec() *SidecarSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

// SidecarSpec defines a sidecar's specifics.
type SidecarSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BizId uint32 `protobuf:"varint,1,opt,name=biz_id,json=bizId,proto3" json:"biz_id,omitempty"`
	// version is sidecar's version
	Version *base.Versioning  `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Metas   []*SidecarAppMeta `protobuf:"bytes,3,rep,name=metas,proto3" json:"metas,omitempty"`
}

func (x *SidecarSpec) Reset() {
	*x = SidecarSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SidecarSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SidecarSpec) ProtoMessage() {}

func (x *SidecarSpec) ProtoReflect() protoreflect.Message {
	mi := &file_feed_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SidecarSpec.ProtoReflect.Descriptor instead.
func (*SidecarSpec) Descriptor() ([]byte, []int) {
	return file_feed_server_proto_rawDescGZIP(), []int{1}
}

func (x *SidecarSpec) GetBizId() uint32 {
	if x != nil {
		return x.BizId
	}
	return 0
}

func (x *SidecarSpec) GetVersion() *base.Versioning {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *SidecarSpec) GetMetas() []*SidecarAppMeta {
	if x != nil {
		return x.Metas
	}
	return nil
}

// SidecarMeta define a sidecar's basic metadata information.
type SidecarAppMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId uint32 `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Uid   string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *SidecarAppMeta) Reset() {
	*x = SidecarAppMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SidecarAppMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SidecarAppMeta) ProtoMessage() {}

func (x *SidecarAppMeta) ProtoReflect() protoreflect.Message {
	mi := &file_feed_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SidecarAppMeta.ProtoReflect.Descriptor instead.
func (*SidecarAppMeta) Descriptor() ([]byte, []int) {
	return file_feed_server_proto_rawDescGZIP(), []int{2}
}

func (x *SidecarAppMeta) GetAppId() uint32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *SidecarAppMeta) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

// HandshakeResp defines handshake resp.
type HandshakeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiVersion *base.Versioning `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// payload is the message's details information which is a json raw bytes.
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *HandshakeResp) Reset() {
	*x = HandshakeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandshakeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandshakeResp) ProtoMessage() {}

func (x *HandshakeResp) ProtoReflect() protoreflect.Message {
	mi := &file_feed_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandshakeResp.ProtoReflect.Descriptor instead.
func (*HandshakeResp) Descriptor() ([]byte, []int) {
	return file_feed_server_proto_rawDescGZIP(), []int{3}
}

func (x *HandshakeResp) GetApiVersion() *base.Versioning {
	if x != nil {
		return x.ApiVersion
	}
	return nil
}

func (x *HandshakeResp) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

// SideWatchMeta defines watch messages send from sidecar to feed server.
type SideWatchMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// version defines message's protocol version from sidecar
	ApiVersion *base.Versioning `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// payload is the message's details information which is a json raw bytes.
	// refer to sfs.SideWatchPayload.
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *SideWatchMeta) Reset() {
	*x = SideWatchMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SideWatchMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SideWatchMeta) ProtoMessage() {}

func (x *SideWatchMeta) ProtoReflect() protoreflect.Message {
	mi := &file_feed_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SideWatchMeta.ProtoReflect.Descriptor instead.
func (*SideWatchMeta) Descriptor() ([]byte, []int) {
	return file_feed_server_proto_rawDescGZIP(), []int{4}
}

func (x *SideWatchMeta) GetApiVersion() *base.Versioning {
	if x != nil {
		return x.ApiVersion
	}
	return nil
}

func (x *SideWatchMeta) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

// FeedWatchMessage defines watch messages send from feed server to sidecar.
type FeedWatchMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// version defines message's protocol version from feed server
	ApiVersion *base.Versioning `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// rid is the message's request id
	Rid string `protobuf:"bytes,2,opt,name=rid,proto3" json:"rid,omitempty"`
	// type is an enum type, it's an substitute of sfs.FeedMessageType.
	Type uint32 `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	// payload is the message's details information which is a json raw bytes.
	Payload []byte `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *FeedWatchMessage) Reset() {
	*x = FeedWatchMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedWatchMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedWatchMessage) ProtoMessage() {}

func (x *FeedWatchMessage) ProtoReflect() protoreflect.Message {
	mi := &file_feed_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedWatchMessage.ProtoReflect.Descriptor instead.
func (*FeedWatchMessage) Descriptor() ([]byte, []int) {
	return file_feed_server_proto_rawDescGZIP(), []int{5}
}

func (x *FeedWatchMessage) GetApiVersion() *base.Versioning {
	if x != nil {
		return x.ApiVersion
	}
	return nil
}

func (x *FeedWatchMessage) GetRid() string {
	if x != nil {
		return x.Rid
	}
	return ""
}

func (x *FeedWatchMessage) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *FeedWatchMessage) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

// MessagingMeta defines the message metadata send from sidecar to upstream server.
type MessagingMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// version defines message's protocol version from feed server
	ApiVersion *base.Versioning `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// rid is the message's request id
	Rid string `protobuf:"bytes,2,opt,name=rid,proto3" json:"rid,omitempty"`
	// type is an enum type, it's an substitute of scs.MessagingType.
	Type uint32 `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	// payload is the message's details information which is a json raw bytes.
	Payload []byte `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *MessagingMeta) Reset() {
	*x = MessagingMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessagingMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessagingMeta) ProtoMessage() {}

func (x *MessagingMeta) ProtoReflect() protoreflect.Message {
	mi := &file_feed_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessagingMeta.ProtoReflect.Descriptor instead.
func (*MessagingMeta) Descriptor() ([]byte, []int) {
	return file_feed_server_proto_rawDescGZIP(), []int{6}
}

func (x *MessagingMeta) GetApiVersion() *base.Versioning {
	if x != nil {
		return x.ApiVersion
	}
	return nil
}

func (x *MessagingMeta) GetRid() string {
	if x != nil {
		return x.Rid
	}
	return ""
}

func (x *MessagingMeta) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *MessagingMeta) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

// MessagingResp defines response from upstream server to sidecar,
// which is empty.
type MessagingResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MessagingResp) Reset() {
	*x = MessagingResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feed_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessagingResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessagingResp) ProtoMessage() {}

func (x *MessagingResp) ProtoReflect() protoreflect.Message {
	mi := &file_feed_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessagingResp.ProtoReflect.Descriptor instead.
func (*MessagingResp) Descriptor() ([]byte, []int) {
	return file_feed_server_proto_rawDescGZIP(), []int{7}
}

var File_feed_server_proto protoreflect.FileDescriptor

var file_feed_server_proto_rawDesc = []byte{
	0x0a, 0x11, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x70, 0x62, 0x66, 0x73, 0x1a, 0x29, 0x62, 0x73, 0x63, 0x70, 0x2e,
	0x69, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x10, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x62, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e,
	0x67, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62,
	0x66, 0x73, 0x2e, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x22, 0x7e, 0x0a, 0x0b, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69, 0x7a, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x62, 0x69, 0x7a, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x05, 0x6d, 0x65, 0x74, 0x61,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x66, 0x73, 0x2e, 0x53,
	0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x41, 0x70, 0x70, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x05, 0x6d,
	0x65, 0x74, 0x61, 0x73, 0x22, 0x39, 0x0a, 0x0e, 0x53, 0x69, 0x64, 0x65, 0x63, 0x61, 0x72, 0x41,
	0x70, 0x70, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22,
	0x5e, 0x0a, 0x0d, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x33, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0x5e, 0x0a, 0x0d, 0x53, 0x69, 0x64, 0x65, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x65, 0x74, 0x61,
	0x12, 0x33, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22,
	0x87, 0x01, 0x0a, 0x10, 0x46, 0x65, 0x65, 0x64, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x61,
	0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x84, 0x01, 0x0a, 0x0d, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x33, 0x0a, 0x0b, 0x61,
	0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x62, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x72, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x22, 0x0f, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x32, 0xb9, 0x01, 0x0a, 0x08, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x3a,
	0x0a, 0x09, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x62,
	0x66, 0x73, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x66, 0x73, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73,
	0x68, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x05, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x66, 0x73, 0x2e, 0x53, 0x69, 0x64, 0x65, 0x57,
	0x61, 0x74, 0x63, 0x68, 0x4d, 0x65, 0x74, 0x61, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x66, 0x73, 0x2e,
	0x46, 0x65, 0x65, 0x64, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x66, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x66, 0x73, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x27, 0x5a,
	0x25, 0x62, 0x73, 0x63, 0x70, 0x2e, 0x69, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x3b, 0x70, 0x62, 0x66, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feed_server_proto_rawDescOnce sync.Once
	file_feed_server_proto_rawDescData = file_feed_server_proto_rawDesc
)

func file_feed_server_proto_rawDescGZIP() []byte {
	file_feed_server_proto_rawDescOnce.Do(func() {
		file_feed_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_feed_server_proto_rawDescData)
	})
	return file_feed_server_proto_rawDescData
}

var file_feed_server_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_feed_server_proto_goTypes = []interface{}{
	(*HandshakeMessage)(nil), // 0: pbfs.HandshakeMessage
	(*SidecarSpec)(nil),      // 1: pbfs.SidecarSpec
	(*SidecarAppMeta)(nil),   // 2: pbfs.SidecarAppMeta
	(*HandshakeResp)(nil),    // 3: pbfs.HandshakeResp
	(*SideWatchMeta)(nil),    // 4: pbfs.SideWatchMeta
	(*FeedWatchMessage)(nil), // 5: pbfs.FeedWatchMessage
	(*MessagingMeta)(nil),    // 6: pbfs.MessagingMeta
	(*MessagingResp)(nil),    // 7: pbfs.MessagingResp
	(*base.Versioning)(nil),  // 8: pbbase.Versioning
}
var file_feed_server_proto_depIdxs = []int32{
	8,  // 0: pbfs.HandshakeMessage.api_version:type_name -> pbbase.Versioning
	1,  // 1: pbfs.HandshakeMessage.spec:type_name -> pbfs.SidecarSpec
	8,  // 2: pbfs.SidecarSpec.version:type_name -> pbbase.Versioning
	2,  // 3: pbfs.SidecarSpec.metas:type_name -> pbfs.SidecarAppMeta
	8,  // 4: pbfs.HandshakeResp.api_version:type_name -> pbbase.Versioning
	8,  // 5: pbfs.SideWatchMeta.api_version:type_name -> pbbase.Versioning
	8,  // 6: pbfs.FeedWatchMessage.api_version:type_name -> pbbase.Versioning
	8,  // 7: pbfs.MessagingMeta.api_version:type_name -> pbbase.Versioning
	0,  // 8: pbfs.Upstream.Handshake:input_type -> pbfs.HandshakeMessage
	4,  // 9: pbfs.Upstream.Watch:input_type -> pbfs.SideWatchMeta
	6,  // 10: pbfs.Upstream.Messaging:input_type -> pbfs.MessagingMeta
	3,  // 11: pbfs.Upstream.Handshake:output_type -> pbfs.HandshakeResp
	5,  // 12: pbfs.Upstream.Watch:output_type -> pbfs.FeedWatchMessage
	7,  // 13: pbfs.Upstream.Messaging:output_type -> pbfs.MessagingResp
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_feed_server_proto_init() }
func file_feed_server_proto_init() {
	if File_feed_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_feed_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandshakeMessage); i {
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
		file_feed_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SidecarSpec); i {
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
		file_feed_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SidecarAppMeta); i {
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
		file_feed_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandshakeResp); i {
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
		file_feed_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SideWatchMeta); i {
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
		file_feed_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedWatchMessage); i {
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
		file_feed_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessagingMeta); i {
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
		file_feed_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessagingResp); i {
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
			RawDescriptor: file_feed_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_feed_server_proto_goTypes,
		DependencyIndexes: file_feed_server_proto_depIdxs,
		MessageInfos:      file_feed_server_proto_msgTypes,
	}.Build()
	File_feed_server_proto = out.File
	file_feed_server_proto_rawDesc = nil
	file_feed_server_proto_goTypes = nil
	file_feed_server_proto_depIdxs = nil
}
