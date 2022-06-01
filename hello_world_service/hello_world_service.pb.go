// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: protos/hello_world_service/hello_world_service.proto

package hello_world_service

import (
	common "github.com/solost23/my_interface/common"
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

type CreateHelloWorldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header     *common.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	HelloWorld string                `protobuf:"bytes,2,opt,name=helloWorld,proto3" json:"helloWorld,omitempty"`
}

func (x *CreateHelloWorldRequest) Reset() {
	*x = CreateHelloWorldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hello_world_service_hello_world_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHelloWorldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHelloWorldRequest) ProtoMessage() {}

func (x *CreateHelloWorldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hello_world_service_hello_world_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHelloWorldRequest.ProtoReflect.Descriptor instead.
func (*CreateHelloWorldRequest) Descriptor() ([]byte, []int) {
	return file_protos_hello_world_service_hello_world_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateHelloWorldRequest) GetHeader() *common.RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *CreateHelloWorldRequest) GetHelloWorld() string {
	if x != nil {
		return x.HelloWorld
	}
	return ""
}

type CreateHelloWorldResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorInfo  *common.ErrorInfo `protobuf:"bytes,1,opt,name=error_info,json=errorInfo,proto3" json:"error_info,omitempty"`
	HelloWorld string            `protobuf:"bytes,2,opt,name=helloWorld,proto3" json:"helloWorld,omitempty"`
}

func (x *CreateHelloWorldResponse) Reset() {
	*x = CreateHelloWorldResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hello_world_service_hello_world_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHelloWorldResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHelloWorldResponse) ProtoMessage() {}

func (x *CreateHelloWorldResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hello_world_service_hello_world_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHelloWorldResponse.ProtoReflect.Descriptor instead.
func (*CreateHelloWorldResponse) Descriptor() ([]byte, []int) {
	return file_protos_hello_world_service_hello_world_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateHelloWorldResponse) GetErrorInfo() *common.ErrorInfo {
	if x != nil {
		return x.ErrorInfo
	}
	return nil
}

func (x *CreateHelloWorldResponse) GetHelloWorld() string {
	if x != nil {
		return x.HelloWorld
	}
	return ""
}

type DeleteHelloWorldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header     *common.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	HelloWorld string                `protobuf:"bytes,2,opt,name=helloWorld,proto3" json:"helloWorld,omitempty"`
}

func (x *DeleteHelloWorldRequest) Reset() {
	*x = DeleteHelloWorldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hello_world_service_hello_world_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHelloWorldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHelloWorldRequest) ProtoMessage() {}

func (x *DeleteHelloWorldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hello_world_service_hello_world_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHelloWorldRequest.ProtoReflect.Descriptor instead.
func (*DeleteHelloWorldRequest) Descriptor() ([]byte, []int) {
	return file_protos_hello_world_service_hello_world_service_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteHelloWorldRequest) GetHeader() *common.RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *DeleteHelloWorldRequest) GetHelloWorld() string {
	if x != nil {
		return x.HelloWorld
	}
	return ""
}

type DeleteHelloWorldResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorInfo  *common.ErrorInfo `protobuf:"bytes,1,opt,name=error_info,json=errorInfo,proto3" json:"error_info,omitempty"`
	HelloWorld string            `protobuf:"bytes,2,opt,name=helloWorld,proto3" json:"helloWorld,omitempty"`
}

func (x *DeleteHelloWorldResponse) Reset() {
	*x = DeleteHelloWorldResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hello_world_service_hello_world_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHelloWorldResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHelloWorldResponse) ProtoMessage() {}

func (x *DeleteHelloWorldResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hello_world_service_hello_world_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHelloWorldResponse.ProtoReflect.Descriptor instead.
func (*DeleteHelloWorldResponse) Descriptor() ([]byte, []int) {
	return file_protos_hello_world_service_hello_world_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteHelloWorldResponse) GetErrorInfo() *common.ErrorInfo {
	if x != nil {
		return x.ErrorInfo
	}
	return nil
}

func (x *DeleteHelloWorldResponse) GetHelloWorld() string {
	if x != nil {
		return x.HelloWorld
	}
	return ""
}

type UpdateHelloWorldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header     *common.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	HelloWorld string                `protobuf:"bytes,2,opt,name=helloWorld,proto3" json:"helloWorld,omitempty"`
}

func (x *UpdateHelloWorldRequest) Reset() {
	*x = UpdateHelloWorldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hello_world_service_hello_world_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHelloWorldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHelloWorldRequest) ProtoMessage() {}

func (x *UpdateHelloWorldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hello_world_service_hello_world_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHelloWorldRequest.ProtoReflect.Descriptor instead.
func (*UpdateHelloWorldRequest) Descriptor() ([]byte, []int) {
	return file_protos_hello_world_service_hello_world_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateHelloWorldRequest) GetHeader() *common.RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *UpdateHelloWorldRequest) GetHelloWorld() string {
	if x != nil {
		return x.HelloWorld
	}
	return ""
}

type UpdateHelloWorldResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorInfo  *common.ErrorInfo `protobuf:"bytes,1,opt,name=error_info,json=errorInfo,proto3" json:"error_info,omitempty"`
	HelloWorld string            `protobuf:"bytes,2,opt,name=helloWorld,proto3" json:"helloWorld,omitempty"`
}

func (x *UpdateHelloWorldResponse) Reset() {
	*x = UpdateHelloWorldResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hello_world_service_hello_world_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHelloWorldResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHelloWorldResponse) ProtoMessage() {}

func (x *UpdateHelloWorldResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hello_world_service_hello_world_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHelloWorldResponse.ProtoReflect.Descriptor instead.
func (*UpdateHelloWorldResponse) Descriptor() ([]byte, []int) {
	return file_protos_hello_world_service_hello_world_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateHelloWorldResponse) GetErrorInfo() *common.ErrorInfo {
	if x != nil {
		return x.ErrorInfo
	}
	return nil
}

func (x *UpdateHelloWorldResponse) GetHelloWorld() string {
	if x != nil {
		return x.HelloWorld
	}
	return ""
}

type ListHelloWorldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header     *common.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	HelloWorld string                `protobuf:"bytes,2,opt,name=helloWorld,proto3" json:"helloWorld,omitempty"`
}

func (x *ListHelloWorldRequest) Reset() {
	*x = ListHelloWorldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hello_world_service_hello_world_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHelloWorldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHelloWorldRequest) ProtoMessage() {}

func (x *ListHelloWorldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hello_world_service_hello_world_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHelloWorldRequest.ProtoReflect.Descriptor instead.
func (*ListHelloWorldRequest) Descriptor() ([]byte, []int) {
	return file_protos_hello_world_service_hello_world_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListHelloWorldRequest) GetHeader() *common.RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ListHelloWorldRequest) GetHelloWorld() string {
	if x != nil {
		return x.HelloWorld
	}
	return ""
}

type ListHelloWorldResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorInfo  *common.ErrorInfo `protobuf:"bytes,1,opt,name=error_info,json=errorInfo,proto3" json:"error_info,omitempty"`
	HelloWorld string            `protobuf:"bytes,2,opt,name=helloWorld,proto3" json:"helloWorld,omitempty"`
}

func (x *ListHelloWorldResponse) Reset() {
	*x = ListHelloWorldResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_hello_world_service_hello_world_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHelloWorldResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHelloWorldResponse) ProtoMessage() {}

func (x *ListHelloWorldResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_hello_world_service_hello_world_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHelloWorldResponse.ProtoReflect.Descriptor instead.
func (*ListHelloWorldResponse) Descriptor() ([]byte, []int) {
	return file_protos_hello_world_service_hello_world_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListHelloWorldResponse) GetErrorInfo() *common.ErrorInfo {
	if x != nil {
		return x.ErrorInfo
	}
	return nil
}

func (x *ListHelloWorldResponse) GetHelloWorld() string {
	if x != nil {
		return x.HelloWorld
	}
	return ""
}

var File_protos_hello_world_service_hello_world_service_proto protoreflect.FileDescriptor

var file_protos_hello_world_service_hello_world_service_proto_rawDesc = []byte{
	0x0a, 0x34, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x6d, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x6d, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x79, 0x0a, 0x18, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x79,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57,
	0x6f, 0x72, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x75, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x0a,
	0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x79, 0x0a,
	0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x6d, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x75, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x1e, 0x0a, 0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x22,
	0x79, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f,
	0x72, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x6d, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x73, 0x0a, 0x15, 0x4c, 0x69,
	0x73, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x1e, 0x0a, 0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x22,
	0x77, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x6d, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x32, 0xc5, 0x04, 0x0a, 0x11, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8b,
	0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f,
	0x72, 0x6c, 0x64, 0x12, 0x39, 0x2e, 0x6d, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a,
	0x2e, 0x6d, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72,
	0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8b, 0x01, 0x0a,
	0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c,
	0x64, 0x12, 0x39, 0x2e, 0x6d, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x6d,
	0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8b, 0x01, 0x0a, 0x10, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x12,
	0x39, 0x2e, 0x6d, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f,
	0x72, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x6d, 0x79, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x85, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73,
	0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x37, 0x2e, 0x6d, 0x79,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x6d, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x73, 0x74, 0x32, 0x33, 0x2f, 0x6d, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x5f, 0x67, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_hello_world_service_hello_world_service_proto_rawDescOnce sync.Once
	file_protos_hello_world_service_hello_world_service_proto_rawDescData = file_protos_hello_world_service_hello_world_service_proto_rawDesc
)

func file_protos_hello_world_service_hello_world_service_proto_rawDescGZIP() []byte {
	file_protos_hello_world_service_hello_world_service_proto_rawDescOnce.Do(func() {
		file_protos_hello_world_service_hello_world_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_hello_world_service_hello_world_service_proto_rawDescData)
	})
	return file_protos_hello_world_service_hello_world_service_proto_rawDescData
}

var file_protos_hello_world_service_hello_world_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protos_hello_world_service_hello_world_service_proto_goTypes = []interface{}{
	(*CreateHelloWorldRequest)(nil),  // 0: my_interface.hello_world_service.CreateHelloWorldRequest
	(*CreateHelloWorldResponse)(nil), // 1: my_interface.hello_world_service.CreateHelloWorldResponse
	(*DeleteHelloWorldRequest)(nil),  // 2: my_interface.hello_world_service.DeleteHelloWorldRequest
	(*DeleteHelloWorldResponse)(nil), // 3: my_interface.hello_world_service.DeleteHelloWorldResponse
	(*UpdateHelloWorldRequest)(nil),  // 4: my_interface.hello_world_service.UpdateHelloWorldRequest
	(*UpdateHelloWorldResponse)(nil), // 5: my_interface.hello_world_service.UpdateHelloWorldResponse
	(*ListHelloWorldRequest)(nil),    // 6: my_interface.hello_world_service.ListHelloWorldRequest
	(*ListHelloWorldResponse)(nil),   // 7: my_interface.hello_world_service.ListHelloWorldResponse
	(*common.RequestHeader)(nil),     // 8: my_interface.common.requestHeader
	(*common.ErrorInfo)(nil),         // 9: my_interface.common.errorInfo
}
var file_protos_hello_world_service_hello_world_service_proto_depIdxs = []int32{
	8,  // 0: my_interface.hello_world_service.CreateHelloWorldRequest.header:type_name -> my_interface.common.requestHeader
	9,  // 1: my_interface.hello_world_service.CreateHelloWorldResponse.error_info:type_name -> my_interface.common.errorInfo
	8,  // 2: my_interface.hello_world_service.DeleteHelloWorldRequest.header:type_name -> my_interface.common.requestHeader
	9,  // 3: my_interface.hello_world_service.DeleteHelloWorldResponse.error_info:type_name -> my_interface.common.errorInfo
	8,  // 4: my_interface.hello_world_service.UpdateHelloWorldRequest.header:type_name -> my_interface.common.requestHeader
	9,  // 5: my_interface.hello_world_service.UpdateHelloWorldResponse.error_info:type_name -> my_interface.common.errorInfo
	8,  // 6: my_interface.hello_world_service.ListHelloWorldRequest.header:type_name -> my_interface.common.requestHeader
	9,  // 7: my_interface.hello_world_service.ListHelloWorldResponse.error_info:type_name -> my_interface.common.errorInfo
	0,  // 8: my_interface.hello_world_service.HelloWorldService.CreateHelloWorld:input_type -> my_interface.hello_world_service.CreateHelloWorldRequest
	2,  // 9: my_interface.hello_world_service.HelloWorldService.DeleteHelloWorld:input_type -> my_interface.hello_world_service.DeleteHelloWorldRequest
	4,  // 10: my_interface.hello_world_service.HelloWorldService.UpdateHelloWorld:input_type -> my_interface.hello_world_service.UpdateHelloWorldRequest
	6,  // 11: my_interface.hello_world_service.HelloWorldService.ListHelloWorld:input_type -> my_interface.hello_world_service.ListHelloWorldRequest
	1,  // 12: my_interface.hello_world_service.HelloWorldService.CreateHelloWorld:output_type -> my_interface.hello_world_service.CreateHelloWorldResponse
	3,  // 13: my_interface.hello_world_service.HelloWorldService.DeleteHelloWorld:output_type -> my_interface.hello_world_service.DeleteHelloWorldResponse
	5,  // 14: my_interface.hello_world_service.HelloWorldService.UpdateHelloWorld:output_type -> my_interface.hello_world_service.UpdateHelloWorldResponse
	7,  // 15: my_interface.hello_world_service.HelloWorldService.ListHelloWorld:output_type -> my_interface.hello_world_service.ListHelloWorldResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_protos_hello_world_service_hello_world_service_proto_init() }
func file_protos_hello_world_service_hello_world_service_proto_init() {
	if File_protos_hello_world_service_hello_world_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_hello_world_service_hello_world_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHelloWorldRequest); i {
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
		file_protos_hello_world_service_hello_world_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHelloWorldResponse); i {
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
		file_protos_hello_world_service_hello_world_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHelloWorldRequest); i {
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
		file_protos_hello_world_service_hello_world_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHelloWorldResponse); i {
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
		file_protos_hello_world_service_hello_world_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHelloWorldRequest); i {
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
		file_protos_hello_world_service_hello_world_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHelloWorldResponse); i {
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
		file_protos_hello_world_service_hello_world_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHelloWorldRequest); i {
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
		file_protos_hello_world_service_hello_world_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHelloWorldResponse); i {
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
			RawDescriptor: file_protos_hello_world_service_hello_world_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_hello_world_service_hello_world_service_proto_goTypes,
		DependencyIndexes: file_protos_hello_world_service_hello_world_service_proto_depIdxs,
		MessageInfos:      file_protos_hello_world_service_hello_world_service_proto_msgTypes,
	}.Build()
	File_protos_hello_world_service_hello_world_service_proto = out.File
	file_protos_hello_world_service_hello_world_service_proto_rawDesc = nil
	file_protos_hello_world_service_hello_world_service_proto_goTypes = nil
	file_protos_hello_world_service_hello_world_service_proto_depIdxs = nil
}
