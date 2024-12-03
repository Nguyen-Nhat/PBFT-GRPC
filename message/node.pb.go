// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.13.0
// source: node.proto

package message

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

// CommitRequest - <v, n, d, m>
type PrePrepareRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ViewId - v
	ViewId int32 `protobuf:"varint,1,opt,name=view_id,json=viewId,proto3" json:"view_id,omitempty"`
	// SequenceId - n
	SequenceId int32 `protobuf:"varint,2,opt,name=sequence_id,json=sequenceId,proto3" json:"sequence_id,omitempty"`
	// Digest - d
	Digest string `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
	// Data - m
	Data *Data `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PrePrepareRequest) Reset() {
	*x = PrePrepareRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrePrepareRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrePrepareRequest) ProtoMessage() {}

func (x *PrePrepareRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrePrepareRequest.ProtoReflect.Descriptor instead.
func (*PrePrepareRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{0}
}

func (x *PrePrepareRequest) GetViewId() int32 {
	if x != nil {
		return x.ViewId
	}
	return 0
}

func (x *PrePrepareRequest) GetSequenceId() int32 {
	if x != nil {
		return x.SequenceId
	}
	return 0
}

func (x *PrePrepareRequest) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *PrePrepareRequest) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

// PrePrepareResponse
type PrePrepareResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PrePrepareResponse) Reset() {
	*x = PrePrepareResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrePrepareResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrePrepareResponse) ProtoMessage() {}

func (x *PrePrepareResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrePrepareResponse.ProtoReflect.Descriptor instead.
func (*PrePrepareResponse) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{1}
}

// PrepareRequest - <v, n, d, i>
type PrepareRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ViewId - v
	ViewId int32 `protobuf:"varint,1,opt,name=view_id,json=viewId,proto3" json:"view_id,omitempty"`
	// SequenceId - n
	SequenceId int32 `protobuf:"varint,2,opt,name=sequence_id,json=sequenceId,proto3" json:"sequence_id,omitempty"`
	// Digest - d
	Digest string `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
	// Node Id - i
	NodeId int32 `protobuf:"varint,4,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (x *PrepareRequest) Reset() {
	*x = PrepareRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepareRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepareRequest) ProtoMessage() {}

func (x *PrepareRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepareRequest.ProtoReflect.Descriptor instead.
func (*PrepareRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{2}
}

func (x *PrepareRequest) GetViewId() int32 {
	if x != nil {
		return x.ViewId
	}
	return 0
}

func (x *PrepareRequest) GetSequenceId() int32 {
	if x != nil {
		return x.SequenceId
	}
	return 0
}

func (x *PrepareRequest) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *PrepareRequest) GetNodeId() int32 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

// PrepareResponse
type PrepareResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PrepareResponse) Reset() {
	*x = PrepareResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepareResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepareResponse) ProtoMessage() {}

func (x *PrepareResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepareResponse.ProtoReflect.Descriptor instead.
func (*PrepareResponse) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{3}
}

// CommitRequest - <v, n, d, i>
type CommitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ViewId - v
	ViewId int32 `protobuf:"varint,1,opt,name=view_id,json=viewId,proto3" json:"view_id,omitempty"`
	// SequenceId - n
	SequenceId int32 `protobuf:"varint,2,opt,name=sequence_id,json=sequenceId,proto3" json:"sequence_id,omitempty"`
	// Digest - d
	Digest string `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
	// Node Id - i
	NodeId int32 `protobuf:"varint,4,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
}

func (x *CommitRequest) Reset() {
	*x = CommitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitRequest) ProtoMessage() {}

func (x *CommitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitRequest.ProtoReflect.Descriptor instead.
func (*CommitRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{4}
}

func (x *CommitRequest) GetViewId() int32 {
	if x != nil {
		return x.ViewId
	}
	return 0
}

func (x *CommitRequest) GetSequenceId() int32 {
	if x != nil {
		return x.SequenceId
	}
	return 0
}

func (x *CommitRequest) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *CommitRequest) GetNodeId() int32 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

// CommitResponse
type CommitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommitResponse) Reset() {
	*x = CommitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitResponse) ProtoMessage() {}

func (x *CommitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitResponse.ProtoReflect.Descriptor instead.
func (*CommitResponse) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{5}
}

// RequestRequest
type RequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Data `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RequestRequest) Reset() {
	*x = RequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestRequest) ProtoMessage() {}

func (x *RequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestRequest.ProtoReflect.Descriptor instead.
func (*RequestRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{6}
}

func (x *RequestRequest) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

// RequestResponse
type RequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RequestResponse) Reset() {
	*x = RequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestResponse) ProtoMessage() {}

func (x *RequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestResponse.ProtoReflect.Descriptor instead.
func (*RequestResponse) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{7}
}

// AddNodeRequest
type AddNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Address
	Port int32 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *AddNodeRequest) Reset() {
	*x = AddNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNodeRequest) ProtoMessage() {}

func (x *AddNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNodeRequest.ProtoReflect.Descriptor instead.
func (*AddNodeRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{8}
}

func (x *AddNodeRequest) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

// AddNodeResponse
type AddNodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// NodeId
	NodeId int32 `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// OtherNodeAddr
	OtherNodePorts []int32 `protobuf:"varint,2,rep,packed,name=other_node_ports,json=otherNodePorts,proto3" json:"other_node_ports,omitempty"`
}

func (x *AddNodeResponse) Reset() {
	*x = AddNodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNodeResponse) ProtoMessage() {}

func (x *AddNodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNodeResponse.ProtoReflect.Descriptor instead.
func (*AddNodeResponse) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{9}
}

func (x *AddNodeResponse) GetNodeId() int32 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *AddNodeResponse) GetOtherNodePorts() []int32 {
	if x != nil {
		return x.OtherNodePorts
	}
	return nil
}

// ListNodePortRequest
type GetListOtherPortRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetListOtherPortRequest) Reset() {
	*x = GetListOtherPortRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListOtherPortRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListOtherPortRequest) ProtoMessage() {}

func (x *GetListOtherPortRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListOtherPortRequest.ProtoReflect.Descriptor instead.
func (*GetListOtherPortRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{10}
}

// ListNodePortResponse
type GetListOtherPortResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ports []int32 `protobuf:"varint,1,rep,packed,name=ports,proto3" json:"ports,omitempty"`
}

func (x *GetListOtherPortResponse) Reset() {
	*x = GetListOtherPortResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListOtherPortResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListOtherPortResponse) ProtoMessage() {}

func (x *GetListOtherPortResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListOtherPortResponse.ProtoReflect.Descriptor instead.
func (*GetListOtherPortResponse) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{11}
}

func (x *GetListOtherPortResponse) GetPorts() []int32 {
	if x != nil {
		return x.Ports
	}
	return nil
}

// Data
type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// PreviousBlockHash
	PreviousBlockHash string `protobuf:"bytes,1,opt,name=previous_block_hash,json=previousBlockHash,proto3" json:"previous_block_hash,omitempty"`
	// BlockHash
	BlockHash string `protobuf:"bytes,2,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	// BlockHeight
	BlockHeight string `protobuf:"bytes,3,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{12}
}

func (x *Data) GetPreviousBlockHash() string {
	if x != nil {
		return x.PreviousBlockHash
	}
	return ""
}

func (x *Data) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *Data) GetBlockHeight() string {
	if x != nil {
		return x.BlockHeight
	}
	return ""
}

var File_node_proto protoreflect.FileDescriptor

var file_node_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x62,
	0x66, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x11, 0x50,
	0x72, 0x65, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x62, 0x66, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x14, 0x0a, 0x12, 0x50, 0x72,
	0x65, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x7b, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0x11, 0x0a,
	0x0f, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x7a, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0x10, 0x0a, 0x0e,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38,
	0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x62, 0x66, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x11, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x0a, 0x0e, 0x41,
	0x64, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x22, 0x54, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a,
	0x10, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x4e, 0x6f,
	0x64, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x19, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x30, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x74, 0x68,
	0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x22, 0x78, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x13,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x6f, 0x75, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x32, 0xde,
	0x03, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f,
	0x0a, 0x0a, 0x50, 0x72, 0x65, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x12, 0x1f, 0x2e, 0x70,
	0x62, 0x66, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x50,
	0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x70, 0x62, 0x66, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x65,
	0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x46, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x66,
	0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x66, 0x74, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x66, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x70, 0x62, 0x66, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x66, 0x74, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x66, 0x74, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x1c, 0x2e, 0x70, 0x62, 0x66, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x41,
	0x64, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x70, 0x62, 0x66, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x64,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x25, 0x2e, 0x70, 0x62, 0x66, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x62, 0x66, 0x74, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x74,
	0x68, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0e, 0x5a, 0x0c, 0x70, 0x62, 0x66, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_proto_rawDescOnce sync.Once
	file_node_proto_rawDescData = file_node_proto_rawDesc
)

func file_node_proto_rawDescGZIP() []byte {
	file_node_proto_rawDescOnce.Do(func() {
		file_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_proto_rawDescData)
	})
	return file_node_proto_rawDescData
}

var file_node_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_node_proto_goTypes = []interface{}{
	(*PrePrepareRequest)(nil),        // 0: pbft.message.PrePrepareRequest
	(*PrePrepareResponse)(nil),       // 1: pbft.message.PrePrepareResponse
	(*PrepareRequest)(nil),           // 2: pbft.message.PrepareRequest
	(*PrepareResponse)(nil),          // 3: pbft.message.PrepareResponse
	(*CommitRequest)(nil),            // 4: pbft.message.CommitRequest
	(*CommitResponse)(nil),           // 5: pbft.message.CommitResponse
	(*RequestRequest)(nil),           // 6: pbft.message.RequestRequest
	(*RequestResponse)(nil),          // 7: pbft.message.RequestResponse
	(*AddNodeRequest)(nil),           // 8: pbft.message.AddNodeRequest
	(*AddNodeResponse)(nil),          // 9: pbft.message.AddNodeResponse
	(*GetListOtherPortRequest)(nil),  // 10: pbft.message.GetListOtherPortRequest
	(*GetListOtherPortResponse)(nil), // 11: pbft.message.GetListOtherPortResponse
	(*Data)(nil),                     // 12: pbft.message.Data
}
var file_node_proto_depIdxs = []int32{
	12, // 0: pbft.message.PrePrepareRequest.data:type_name -> pbft.message.Data
	12, // 1: pbft.message.RequestRequest.data:type_name -> pbft.message.Data
	0,  // 2: pbft.message.NodeService.PrePrepare:input_type -> pbft.message.PrePrepareRequest
	2,  // 3: pbft.message.NodeService.Prepare:input_type -> pbft.message.PrepareRequest
	4,  // 4: pbft.message.NodeService.Commit:input_type -> pbft.message.CommitRequest
	6,  // 5: pbft.message.NodeService.Request:input_type -> pbft.message.RequestRequest
	8,  // 6: pbft.message.NodeService.AddNode:input_type -> pbft.message.AddNodeRequest
	10, // 7: pbft.message.NodeService.GetListOtherPort:input_type -> pbft.message.GetListOtherPortRequest
	1,  // 8: pbft.message.NodeService.PrePrepare:output_type -> pbft.message.PrePrepareResponse
	3,  // 9: pbft.message.NodeService.Prepare:output_type -> pbft.message.PrepareResponse
	5,  // 10: pbft.message.NodeService.Commit:output_type -> pbft.message.CommitResponse
	7,  // 11: pbft.message.NodeService.Request:output_type -> pbft.message.RequestResponse
	9,  // 12: pbft.message.NodeService.AddNode:output_type -> pbft.message.AddNodeResponse
	11, // 13: pbft.message.NodeService.GetListOtherPort:output_type -> pbft.message.GetListOtherPortResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_node_proto_init() }
func file_node_proto_init() {
	if File_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrePrepareRequest); i {
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
		file_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrePrepareResponse); i {
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
		file_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepareRequest); i {
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
		file_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepareResponse); i {
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
		file_node_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitRequest); i {
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
		file_node_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitResponse); i {
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
		file_node_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestRequest); i {
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
		file_node_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestResponse); i {
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
		file_node_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNodeRequest); i {
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
		file_node_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNodeResponse); i {
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
		file_node_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListOtherPortRequest); i {
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
		file_node_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListOtherPortResponse); i {
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
		file_node_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
			RawDescriptor: file_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_node_proto_goTypes,
		DependencyIndexes: file_node_proto_depIdxs,
		MessageInfos:      file_node_proto_msgTypes,
	}.Build()
	File_node_proto = out.File
	file_node_proto_rawDesc = nil
	file_node_proto_goTypes = nil
	file_node_proto_depIdxs = nil
}
