// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: talker/v1/talker.proto

package talkerv1

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

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target  string  `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Message string  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Foo     *string `protobuf:"bytes,3,opt,name=foo,proto3,oneof" json:"foo,omitempty"`
	Bar     *int32  `protobuf:"varint,4,opt,name=bar,proto3,oneof" json:"bar,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talker_v1_talker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_talker_v1_talker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_talker_v1_talker_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *HelloRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *HelloRequest) GetFoo() string {
	if x != nil && x.Foo != nil {
		return *x.Foo
	}
	return ""
}

func (x *HelloRequest) GetBar() int32 {
	if x != nil && x.Bar != nil {
		return *x.Bar
	}
	return 0
}

type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message       string         `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	EmbeddedThing *EmbeddedThing `protobuf:"bytes,2,opt,name=embedded_thing,json=embeddedThing,proto3" json:"embedded_thing,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talker_v1_talker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_talker_v1_talker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_talker_v1_talker_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *HelloResponse) GetEmbeddedThing() *EmbeddedThing {
	if x != nil {
		return x.EmbeddedThing
	}
	return nil
}

type EmbeddedThing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InternalMessage *string `protobuf:"bytes,1,opt,name=internal_message,json=internalMessage,proto3,oneof" json:"internal_message,omitempty"`
}

func (x *EmbeddedThing) Reset() {
	*x = EmbeddedThing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talker_v1_talker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmbeddedThing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmbeddedThing) ProtoMessage() {}

func (x *EmbeddedThing) ProtoReflect() protoreflect.Message {
	mi := &file_talker_v1_talker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmbeddedThing.ProtoReflect.Descriptor instead.
func (*EmbeddedThing) Descriptor() ([]byte, []int) {
	return file_talker_v1_talker_proto_rawDescGZIP(), []int{2}
}

func (x *EmbeddedThing) GetInternalMessage() string {
	if x != nil && x.InternalMessage != nil {
		return *x.InternalMessage
	}
	return ""
}

type WhisperRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target  string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *WhisperRequest) Reset() {
	*x = WhisperRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talker_v1_talker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhisperRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhisperRequest) ProtoMessage() {}

func (x *WhisperRequest) ProtoReflect() protoreflect.Message {
	mi := &file_talker_v1_talker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhisperRequest.ProtoReflect.Descriptor instead.
func (*WhisperRequest) Descriptor() ([]byte, []int) {
	return file_talker_v1_talker_proto_rawDescGZIP(), []int{3}
}

func (x *WhisperRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *WhisperRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type WhisperResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *WhisperResponse) Reset() {
	*x = WhisperResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talker_v1_talker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhisperResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhisperResponse) ProtoMessage() {}

func (x *WhisperResponse) ProtoReflect() protoreflect.Message {
	mi := &file_talker_v1_talker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhisperResponse.ProtoReflect.Descriptor instead.
func (*WhisperResponse) Descriptor() ([]byte, []int) {
	return file_talker_v1_talker_proto_rawDescGZIP(), []int{4}
}

func (x *WhisperResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ByeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target  string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ByeRequest) Reset() {
	*x = ByeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talker_v1_talker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByeRequest) ProtoMessage() {}

func (x *ByeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_talker_v1_talker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByeRequest.ProtoReflect.Descriptor instead.
func (*ByeRequest) Descriptor() ([]byte, []int) {
	return file_talker_v1_talker_proto_rawDescGZIP(), []int{5}
}

func (x *ByeRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *ByeRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ByeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ByeResponse) Reset() {
	*x = ByeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talker_v1_talker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByeResponse) ProtoMessage() {}

func (x *ByeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_talker_v1_talker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByeResponse.ProtoReflect.Descriptor instead.
func (*ByeResponse) Descriptor() ([]byte, []int) {
	return file_talker_v1_talker_proto_rawDescGZIP(), []int{6}
}

func (x *ByeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type FibonacciRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Places uint32 `protobuf:"varint,1,opt,name=places,proto3" json:"places,omitempty"`
	Target string `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *FibonacciRequest) Reset() {
	*x = FibonacciRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talker_v1_talker_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FibonacciRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FibonacciRequest) ProtoMessage() {}

func (x *FibonacciRequest) ProtoReflect() protoreflect.Message {
	mi := &file_talker_v1_talker_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FibonacciRequest.ProtoReflect.Descriptor instead.
func (*FibonacciRequest) Descriptor() ([]byte, []int) {
	return file_talker_v1_talker_proto_rawDescGZIP(), []int{7}
}

func (x *FibonacciRequest) GetPlaces() uint32 {
	if x != nil {
		return x.Places
	}
	return 0
}

func (x *FibonacciRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

type FibonacciResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position uint32 `protobuf:"varint,1,opt,name=position,proto3" json:"position,omitempty"`
	Value    uint32 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *FibonacciResponse) Reset() {
	*x = FibonacciResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talker_v1_talker_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FibonacciResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FibonacciResponse) ProtoMessage() {}

func (x *FibonacciResponse) ProtoReflect() protoreflect.Message {
	mi := &file_talker_v1_talker_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FibonacciResponse.ProtoReflect.Descriptor instead.
func (*FibonacciResponse) Descriptor() ([]byte, []int) {
	return file_talker_v1_talker_proto_rawDescGZIP(), []int{8}
}

func (x *FibonacciResponse) GetPosition() uint32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *FibonacciResponse) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type AdditionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addend uint32 `protobuf:"varint,1,opt,name=addend,proto3" json:"addend,omitempty"`
	Target string `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *AdditionRequest) Reset() {
	*x = AdditionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talker_v1_talker_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdditionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdditionRequest) ProtoMessage() {}

func (x *AdditionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_talker_v1_talker_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdditionRequest.ProtoReflect.Descriptor instead.
func (*AdditionRequest) Descriptor() ([]byte, []int) {
	return file_talker_v1_talker_proto_rawDescGZIP(), []int{9}
}

func (x *AdditionRequest) GetAddend() uint32 {
	if x != nil {
		return x.Addend
	}
	return 0
}

func (x *AdditionRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

type AdditionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum uint64 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *AdditionResponse) Reset() {
	*x = AdditionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talker_v1_talker_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdditionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdditionResponse) ProtoMessage() {}

func (x *AdditionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_talker_v1_talker_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdditionResponse.ProtoReflect.Descriptor instead.
func (*AdditionResponse) Descriptor() ([]byte, []int) {
	return file_talker_v1_talker_proto_rawDescGZIP(), []int{10}
}

func (x *AdditionResponse) GetSum() uint64 {
	if x != nil {
		return x.Sum
	}
	return 0
}

type FailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target  string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *FailRequest) Reset() {
	*x = FailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talker_v1_talker_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailRequest) ProtoMessage() {}

func (x *FailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_talker_v1_talker_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailRequest.ProtoReflect.Descriptor instead.
func (*FailRequest) Descriptor() ([]byte, []int) {
	return file_talker_v1_talker_proto_rawDescGZIP(), []int{11}
}

func (x *FailRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *FailRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type FailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *FailResponse) Reset() {
	*x = FailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talker_v1_talker_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailResponse) ProtoMessage() {}

func (x *FailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_talker_v1_talker_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailResponse.ProtoReflect.Descriptor instead.
func (*FailResponse) Descriptor() ([]byte, []int) {
	return file_talker_v1_talker_proto_rawDescGZIP(), []int{12}
}

func (x *FailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type FailOnFourRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Value  uint32 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *FailOnFourRequest) Reset() {
	*x = FailOnFourRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talker_v1_talker_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailOnFourRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailOnFourRequest) ProtoMessage() {}

func (x *FailOnFourRequest) ProtoReflect() protoreflect.Message {
	mi := &file_talker_v1_talker_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailOnFourRequest.ProtoReflect.Descriptor instead.
func (*FailOnFourRequest) Descriptor() ([]byte, []int) {
	return file_talker_v1_talker_proto_rawDescGZIP(), []int{13}
}

func (x *FailOnFourRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *FailOnFourRequest) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type FailOnFourResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *FailOnFourResponse) Reset() {
	*x = FailOnFourResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_talker_v1_talker_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailOnFourResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailOnFourResponse) ProtoMessage() {}

func (x *FailOnFourResponse) ProtoReflect() protoreflect.Message {
	mi := &file_talker_v1_talker_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailOnFourResponse.ProtoReflect.Descriptor instead.
func (*FailOnFourResponse) Descriptor() ([]byte, []int) {
	return file_talker_v1_talker_proto_rawDescGZIP(), []int{14}
}

func (x *FailOnFourResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_talker_v1_talker_proto protoreflect.FileDescriptor

var file_talker_v1_talker_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x61, 0x6c, 0x6b, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x6c, 0x6b,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x61, 0x6c, 0x6b, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x22, 0x7e, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x15, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03,
	0x62, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x03, 0x62, 0x61, 0x72,
	0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x66, 0x6f, 0x6f, 0x42, 0x06, 0x0a, 0x04, 0x5f,
	0x62, 0x61, 0x72, 0x22, 0x6a, 0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3f,
	0x0a, 0x0e, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x5f, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x61, 0x6c, 0x6b, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x54, 0x68, 0x69, 0x6e, 0x67,
	0x52, 0x0d, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x22,
	0x54, 0x0a, 0x0d, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x54, 0x68, 0x69, 0x6e, 0x67,
	0x12, 0x2e, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x13, 0x0a, 0x11, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x42, 0x0a, 0x0e, 0x57, 0x68, 0x69, 0x73, 0x70, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2b, 0x0a, 0x0f, 0x57, 0x68, 0x69,
	0x73, 0x70, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3e, 0x0a, 0x0a, 0x42, 0x79, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x27, 0x0a, 0x0b, 0x42, 0x79, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x42, 0x0a, 0x10, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x22, 0x45, 0x0a, 0x11, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x41, 0x0a, 0x0f, 0x41, 0x64,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x64, 0x64, 0x65, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61,
	0x64, 0x64, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x24, 0x0a,
	0x10, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x73, 0x75, 0x6d, 0x22, 0x3f, 0x0a, 0x0b, 0x46, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x46, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x41,
	0x0a, 0x11, 0x46, 0x61, 0x69, 0x6c, 0x4f, 0x6e, 0x46, 0x6f, 0x75, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x2e, 0x0a, 0x12, 0x46, 0x61, 0x69, 0x6c, 0x4f, 0x6e, 0x46, 0x6f, 0x75, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0x4b, 0x0a, 0x0d, 0x54, 0x61, 0x6c, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x17, 0x2e, 0x74, 0x61,
	0x6c, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x61, 0x6c, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x36,
	0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x6f, 0x6c,
	0x6f, 0x63, 0x73, 0x2f, 0x74, 0x72, 0x79, 0x2d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x74, 0x61, 0x6c, 0x6b, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x61,
	0x6c, 0x6b, 0x65, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_talker_v1_talker_proto_rawDescOnce sync.Once
	file_talker_v1_talker_proto_rawDescData = file_talker_v1_talker_proto_rawDesc
)

func file_talker_v1_talker_proto_rawDescGZIP() []byte {
	file_talker_v1_talker_proto_rawDescOnce.Do(func() {
		file_talker_v1_talker_proto_rawDescData = protoimpl.X.CompressGZIP(file_talker_v1_talker_proto_rawDescData)
	})
	return file_talker_v1_talker_proto_rawDescData
}

var file_talker_v1_talker_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_talker_v1_talker_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),       // 0: talker.v1.HelloRequest
	(*HelloResponse)(nil),      // 1: talker.v1.HelloResponse
	(*EmbeddedThing)(nil),      // 2: talker.v1.EmbeddedThing
	(*WhisperRequest)(nil),     // 3: talker.v1.WhisperRequest
	(*WhisperResponse)(nil),    // 4: talker.v1.WhisperResponse
	(*ByeRequest)(nil),         // 5: talker.v1.ByeRequest
	(*ByeResponse)(nil),        // 6: talker.v1.ByeResponse
	(*FibonacciRequest)(nil),   // 7: talker.v1.FibonacciRequest
	(*FibonacciResponse)(nil),  // 8: talker.v1.FibonacciResponse
	(*AdditionRequest)(nil),    // 9: talker.v1.AdditionRequest
	(*AdditionResponse)(nil),   // 10: talker.v1.AdditionResponse
	(*FailRequest)(nil),        // 11: talker.v1.FailRequest
	(*FailResponse)(nil),       // 12: talker.v1.FailResponse
	(*FailOnFourRequest)(nil),  // 13: talker.v1.FailOnFourRequest
	(*FailOnFourResponse)(nil), // 14: talker.v1.FailOnFourResponse
}
var file_talker_v1_talker_proto_depIdxs = []int32{
	2, // 0: talker.v1.HelloResponse.embedded_thing:type_name -> talker.v1.EmbeddedThing
	0, // 1: talker.v1.TalkerService.Hello:input_type -> talker.v1.HelloRequest
	1, // 2: talker.v1.TalkerService.Hello:output_type -> talker.v1.HelloResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_talker_v1_talker_proto_init() }
func file_talker_v1_talker_proto_init() {
	if File_talker_v1_talker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_talker_v1_talker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_talker_v1_talker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
		file_talker_v1_talker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmbeddedThing); i {
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
		file_talker_v1_talker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhisperRequest); i {
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
		file_talker_v1_talker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhisperResponse); i {
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
		file_talker_v1_talker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByeRequest); i {
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
		file_talker_v1_talker_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByeResponse); i {
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
		file_talker_v1_talker_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FibonacciRequest); i {
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
		file_talker_v1_talker_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FibonacciResponse); i {
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
		file_talker_v1_talker_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdditionRequest); i {
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
		file_talker_v1_talker_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdditionResponse); i {
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
		file_talker_v1_talker_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailRequest); i {
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
		file_talker_v1_talker_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailResponse); i {
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
		file_talker_v1_talker_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailOnFourRequest); i {
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
		file_talker_v1_talker_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailOnFourResponse); i {
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
	file_talker_v1_talker_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_talker_v1_talker_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_talker_v1_talker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_talker_v1_talker_proto_goTypes,
		DependencyIndexes: file_talker_v1_talker_proto_depIdxs,
		MessageInfos:      file_talker_v1_talker_proto_msgTypes,
	}.Build()
	File_talker_v1_talker_proto = out.File
	file_talker_v1_talker_proto_rawDesc = nil
	file_talker_v1_talker_proto_goTypes = nil
	file_talker_v1_talker_proto_depIdxs = nil
}
