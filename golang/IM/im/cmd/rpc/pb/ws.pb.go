// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: ws.proto

package pb

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

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ver       int64  `protobuf:"varint,1,opt,name=Ver,proto3" json:"Ver,omitempty"`             // protocol version
	Operation int64  `protobuf:"varint,2,opt,name=Operation,proto3" json:"Operation,omitempty"` // operation for request
	SeqId     string `protobuf:"bytes,3,opt,name=SeqId,proto3" json:"SeqId,omitempty"`          // sequence number chosen by client
	Body      []byte `protobuf:"bytes,4,opt,name=Body,proto3" json:"Body,omitempty"`            // binary body bytes
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ws_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_ws_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_ws_proto_rawDescGZIP(), []int{0}
}

func (x *Msg) GetVer() int64 {
	if x != nil {
		return x.Ver
	}
	return 0
}

func (x *Msg) GetOperation() int64 {
	if x != nil {
		return x.Operation
	}
	return 0
}

func (x *Msg) GetSeqId() string {
	if x != nil {
		return x.SeqId
	}
	return ""
}

func (x *Msg) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type PushMsgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Msg    *Msg  `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *PushMsgRequest) Reset() {
	*x = PushMsgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ws_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMsgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsgRequest) ProtoMessage() {}

func (x *PushMsgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ws_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsgRequest.ProtoReflect.Descriptor instead.
func (*PushMsgRequest) Descriptor() ([]byte, []int) {
	return file_ws_proto_rawDescGZIP(), []int{1}
}

func (x *PushMsgRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PushMsgRequest) GetMsg() *Msg {
	if x != nil {
		return x.Msg
	}
	return nil
}

type PushRoomMsgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId int64 `protobuf:"varint,1,opt,name=RoomId,proto3" json:"RoomId,omitempty"`
	Msg    *Msg  `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *PushRoomMsgRequest) Reset() {
	*x = PushRoomMsgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ws_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushRoomMsgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRoomMsgRequest) ProtoMessage() {}

func (x *PushRoomMsgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ws_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRoomMsgRequest.ProtoReflect.Descriptor instead.
func (*PushRoomMsgRequest) Descriptor() ([]byte, []int) {
	return file_ws_proto_rawDescGZIP(), []int{2}
}

func (x *PushRoomMsgRequest) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *PushRoomMsgRequest) GetMsg() *Msg {
	if x != nil {
		return x.Msg
	}
	return nil
}

type PushRoomCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId int64 `protobuf:"varint,1,opt,name=RoomId,proto3" json:"RoomId,omitempty"`
	Count  int64 `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (x *PushRoomCountRequest) Reset() {
	*x = PushRoomCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ws_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushRoomCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRoomCountRequest) ProtoMessage() {}

func (x *PushRoomCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ws_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRoomCountRequest.ProtoReflect.Descriptor instead.
func (*PushRoomCountRequest) Descriptor() ([]byte, []int) {
	return file_ws_proto_rawDescGZIP(), []int{3}
}

func (x *PushRoomCountRequest) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *PushRoomCountRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type SuccessReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *SuccessReply) Reset() {
	*x = SuccessReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ws_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessReply) ProtoMessage() {}

func (x *SuccessReply) ProtoReflect() protoreflect.Message {
	mi := &file_ws_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessReply.ProtoReflect.Descriptor instead.
func (*SuccessReply) Descriptor() ([]byte, []int) {
	return file_ws_proto_rawDescGZIP(), []int{4}
}

func (x *SuccessReply) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SuccessReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthToken string `protobuf:"bytes,1,opt,name=AuthToken,proto3" json:"AuthToken,omitempty"`
	RoomId    int64  `protobuf:"varint,2,opt,name=RoomId,proto3" json:"RoomId,omitempty"`
	ServerId  string `protobuf:"bytes,3,opt,name=ServerId,proto3" json:"ServerId,omitempty"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ws_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ws_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_ws_proto_rawDescGZIP(), []int{5}
}

func (x *ConnectRequest) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

func (x *ConnectRequest) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *ConnectRequest) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

type ConnectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *ConnectReply) Reset() {
	*x = ConnectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ws_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectReply) ProtoMessage() {}

func (x *ConnectReply) ProtoReflect() protoreflect.Message {
	mi := &file_ws_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectReply.ProtoReflect.Descriptor instead.
func (*ConnectReply) Descriptor() ([]byte, []int) {
	return file_ws_proto_rawDescGZIP(), []int{6}
}

func (x *ConnectReply) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DisConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId int64 `protobuf:"varint,1,opt,name=RoomId,proto3" json:"RoomId,omitempty"`
	UserId int64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *DisConnectRequest) Reset() {
	*x = DisConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ws_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisConnectRequest) ProtoMessage() {}

func (x *DisConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ws_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisConnectRequest.ProtoReflect.Descriptor instead.
func (*DisConnectRequest) Descriptor() ([]byte, []int) {
	return file_ws_proto_rawDescGZIP(), []int{7}
}

func (x *DisConnectRequest) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *DisConnectRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DisConnectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Has bool `protobuf:"varint,1,opt,name=Has,proto3" json:"Has,omitempty"`
}

func (x *DisConnectReply) Reset() {
	*x = DisConnectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ws_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisConnectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisConnectReply) ProtoMessage() {}

func (x *DisConnectReply) ProtoReflect() protoreflect.Message {
	mi := &file_ws_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisConnectReply.ProtoReflect.Descriptor instead.
func (*DisConnectReply) Descriptor() ([]byte, []int) {
	return file_ws_proto_rawDescGZIP(), []int{8}
}

func (x *DisConnectReply) GetHas() bool {
	if x != nil {
		return x.Has
	}
	return false
}

var File_ws_proto protoreflect.FileDescriptor

var file_ws_proto_rawDesc = []byte{
	0x0a, 0x08, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x5f,
	0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x56, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x56, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x65, 0x71, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x65, 0x71, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x42,
	0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x22,
	0x43, 0x0a, 0x0e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x03, 0x4d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x73, 0x67, 0x52,
	0x03, 0x4d, 0x73, 0x67, 0x22, 0x47, 0x0a, 0x12, 0x50, 0x75, 0x73, 0x68, 0x52, 0x6f, 0x6f, 0x6d,
	0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f,
	0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x52, 0x6f, 0x6f, 0x6d,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0x44, 0x0a,
	0x14, 0x50, 0x75, 0x73, 0x68, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x34, 0x0a, 0x0c, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0x62, 0x0a, 0x0e, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x41,
	0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6f,
	0x6d, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x26, 0x0a,
	0x0c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x11, 0x44, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f,
	0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x52, 0x6f, 0x6f, 0x6d,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x23, 0x0a, 0x0f, 0x44, 0x69,
	0x73, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x48, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x48, 0x61, 0x73, 0x32,
	0xaa, 0x01, 0x0a, 0x02, 0x77, 0x73, 0x12, 0x39, 0x0a, 0x0d, 0x50, 0x75, 0x73, 0x68, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x73,
	0x68, 0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x2f, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x12, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x38, 0x0a, 0x0a, 0x44, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x69, 0x73,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ws_proto_rawDescOnce sync.Once
	file_ws_proto_rawDescData = file_ws_proto_rawDesc
)

func file_ws_proto_rawDescGZIP() []byte {
	file_ws_proto_rawDescOnce.Do(func() {
		file_ws_proto_rawDescData = protoimpl.X.CompressGZIP(file_ws_proto_rawDescData)
	})
	return file_ws_proto_rawDescData
}

var file_ws_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_ws_proto_goTypes = []interface{}{
	(*Msg)(nil),                  // 0: pb.Msg
	(*PushMsgRequest)(nil),       // 1: pb.PushMsgRequest
	(*PushRoomMsgRequest)(nil),   // 2: pb.PushRoomMsgRequest
	(*PushRoomCountRequest)(nil), // 3: pb.PushRoomCountRequest
	(*SuccessReply)(nil),         // 4: pb.SuccessReply
	(*ConnectRequest)(nil),       // 5: pb.ConnectRequest
	(*ConnectReply)(nil),         // 6: pb.ConnectReply
	(*DisConnectRequest)(nil),    // 7: pb.DisConnectRequest
	(*DisConnectReply)(nil),      // 8: pb.DisConnectReply
}
var file_ws_proto_depIdxs = []int32{
	0, // 0: pb.PushMsgRequest.Msg:type_name -> pb.Msg
	0, // 1: pb.PushRoomMsgRequest.Msg:type_name -> pb.Msg
	2, // 2: pb.ws.PushSingleMsg:input_type -> pb.PushRoomMsgRequest
	5, // 3: pb.ws.Connect:input_type -> pb.ConnectRequest
	7, // 4: pb.ws.DisConnect:input_type -> pb.DisConnectRequest
	4, // 5: pb.ws.PushSingleMsg:output_type -> pb.SuccessReply
	6, // 6: pb.ws.Connect:output_type -> pb.ConnectReply
	8, // 7: pb.ws.DisConnect:output_type -> pb.DisConnectReply
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ws_proto_init() }
func file_ws_proto_init() {
	if File_ws_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ws_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
		file_ws_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMsgRequest); i {
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
		file_ws_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushRoomMsgRequest); i {
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
		file_ws_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushRoomCountRequest); i {
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
		file_ws_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessReply); i {
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
		file_ws_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
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
		file_ws_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectReply); i {
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
		file_ws_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisConnectRequest); i {
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
		file_ws_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisConnectReply); i {
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
			RawDescriptor: file_ws_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ws_proto_goTypes,
		DependencyIndexes: file_ws_proto_depIdxs,
		MessageInfos:      file_ws_proto_msgTypes,
	}.Build()
	File_ws_proto = out.File
	file_ws_proto_rawDesc = nil
	file_ws_proto_goTypes = nil
	file_ws_proto_depIdxs = nil
}
