// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: protos/protos.proto

package tarea_sist_distribuido

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

type MessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request string `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Autor   string `protobuf:"bytes,2,opt,name=autor,proto3" json:"autor,omitempty"`
	Reloj   string `protobuf:"bytes,3,opt,name=reloj,proto3" json:"reloj,omitempty"`
	Ip      string `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *MessageRequest) Reset() {
	*x = MessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_protos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageRequest) ProtoMessage() {}

func (x *MessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_protos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageRequest.ProtoReflect.Descriptor instead.
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return file_protos_protos_proto_rawDescGZIP(), []int{0}
}

func (x *MessageRequest) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

func (x *MessageRequest) GetAutor() string {
	if x != nil {
		return x.Autor
	}
	return ""
}

func (x *MessageRequest) GetReloj() string {
	if x != nil {
		return x.Reloj
	}
	return ""
}

func (x *MessageRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type MessageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply string `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *MessageReply) Reset() {
	*x = MessageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_protos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageReply) ProtoMessage() {}

func (x *MessageReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_protos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageReply.ProtoReflect.Descriptor instead.
func (*MessageReply) Descriptor() ([]byte, []int) {
	return file_protos_protos_proto_rawDescGZIP(), []int{1}
}

func (x *MessageReply) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

// Se solicita a los servidores que envíen sus logs
type CoordinacionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request string `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *CoordinacionRequest) Reset() {
	*x = CoordinacionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_protos_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoordinacionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoordinacionRequest) ProtoMessage() {}

func (x *CoordinacionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_protos_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoordinacionRequest.ProtoReflect.Descriptor instead.
func (*CoordinacionRequest) Descriptor() ([]byte, []int) {
	return file_protos_protos_proto_rawDescGZIP(), []int{2}
}

func (x *CoordinacionRequest) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

// Envío de logs, uno por cada planeta
// Formato:
// {
//  planeta1: [log1, log2,.. logn],
//  planeta2: [...]
// }
type CoordinacionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planetas string `protobuf:"bytes,1,opt,name=planetas,proto3" json:"planetas,omitempty"`
	Logs     string `protobuf:"bytes,2,opt,name=logs,proto3" json:"logs,omitempty"`
	Vector   string `protobuf:"bytes,3,opt,name=vector,proto3" json:"vector,omitempty"`
}

func (x *CoordinacionReply) Reset() {
	*x = CoordinacionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_protos_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoordinacionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoordinacionReply) ProtoMessage() {}

func (x *CoordinacionReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_protos_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoordinacionReply.ProtoReflect.Descriptor instead.
func (*CoordinacionReply) Descriptor() ([]byte, []int) {
	return file_protos_protos_proto_rawDescGZIP(), []int{3}
}

func (x *CoordinacionReply) GetPlanetas() string {
	if x != nil {
		return x.Planetas
	}
	return ""
}

func (x *CoordinacionReply) GetLogs() string {
	if x != nil {
		return x.Logs
	}
	return ""
}

func (x *CoordinacionReply) GetVector() string {
	if x != nil {
		return x.Vector
	}
	return ""
}

type ReestructuracionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planetas    string `protobuf:"bytes,1,opt,name=planetas,proto3" json:"planetas,omitempty"`
	Vectores    string `protobuf:"bytes,2,opt,name=vectores,proto3" json:"vectores,omitempty"`
	Registrotxt string `protobuf:"bytes,3,opt,name=registrotxt,proto3" json:"registrotxt,omitempty"`
}

func (x *ReestructuracionRequest) Reset() {
	*x = ReestructuracionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_protos_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReestructuracionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReestructuracionRequest) ProtoMessage() {}

func (x *ReestructuracionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_protos_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReestructuracionRequest.ProtoReflect.Descriptor instead.
func (*ReestructuracionRequest) Descriptor() ([]byte, []int) {
	return file_protos_protos_proto_rawDescGZIP(), []int{4}
}

func (x *ReestructuracionRequest) GetPlanetas() string {
	if x != nil {
		return x.Planetas
	}
	return ""
}

func (x *ReestructuracionRequest) GetVectores() string {
	if x != nil {
		return x.Vectores
	}
	return ""
}

func (x *ReestructuracionRequest) GetRegistrotxt() string {
	if x != nil {
		return x.Registrotxt
	}
	return ""
}

type ReestructuracionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply string `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *ReestructuracionReply) Reset() {
	*x = ReestructuracionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_protos_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReestructuracionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReestructuracionReply) ProtoMessage() {}

func (x *ReestructuracionReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_protos_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReestructuracionReply.ProtoReflect.Descriptor instead.
func (*ReestructuracionReply) Descriptor() ([]byte, []int) {
	return file_protos_protos_proto_rawDescGZIP(), []int{5}
}

func (x *ReestructuracionReply) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

type RelojRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request string `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *RelojRequest) Reset() {
	*x = RelojRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_protos_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelojRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelojRequest) ProtoMessage() {}

func (x *RelojRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_protos_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelojRequest.ProtoReflect.Descriptor instead.
func (*RelojRequest) Descriptor() ([]byte, []int) {
	return file_protos_protos_proto_rawDescGZIP(), []int{6}
}

func (x *RelojRequest) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

type RelojReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply string `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *RelojReply) Reset() {
	*x = RelojReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_protos_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelojReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelojReply) ProtoMessage() {}

func (x *RelojReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_protos_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelojReply.ProtoReflect.Descriptor instead.
func (*RelojReply) Descriptor() ([]byte, []int) {
	return file_protos_protos_proto_rawDescGZIP(), []int{7}
}

func (x *RelojReply) GetReply() string {
	if x != nil {
		return x.Reply
	}
	return ""
}

var File_protos_protos_proto protoreflect.FileDescriptor

var file_protos_protos_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x66, 0x0a,
	0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x75, 0x74,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x75, 0x74, 0x6f, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x24, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2f, 0x0a, 0x13, 0x43,
	0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5b, 0x0a, 0x11,
	0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x73, 0x0a, 0x17, 0x52, 0x65, 0x65,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x6f, 0x74, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x6f, 0x74, 0x78, 0x74, 0x22, 0x2d,
	0x0a, 0x15, 0x52, 0x65, 0x65, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x61, 0x63, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x28, 0x0a,
	0x0c, 0x52, 0x65, 0x6c, 0x6f, 0x6a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x22, 0x0a, 0x0a, 0x52, 0x65, 0x6c, 0x6f, 0x6a,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xa9, 0x02, 0x0a, 0x12,
	0x4d, 0x61, 0x6e, 0x65, 0x6a, 0x6f, 0x43, 0x6f, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x63, 0x69,
	0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x72, 0x12,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x45, 0x0a, 0x09, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x72, 0x12, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x63, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0d, 0x52, 0x65, 0x65, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x61, 0x72, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x52, 0x65, 0x65, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x61, 0x63, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x52, 0x65, 0x65, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x61, 0x63, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0e, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x74, 0x61, 0x72, 0x52, 0x65, 0x6c, 0x6f, 0x6a, 0x12, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x6c, 0x6f, 0x6a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52, 0x65, 0x6c, 0x6f, 0x6a,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x65, 0x6c, 0x69, 0x70, 0x65, 0x2d, 0x61, 0x67, 0x75,
	0x69, 0x72, 0x72, 0x65, 0x2f, 0x74, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x73, 0x69, 0x73, 0x74, 0x5f,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x69, 0x64, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x3b, 0x74, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x73, 0x69, 0x73, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x69, 0x64, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_protos_proto_rawDescOnce sync.Once
	file_protos_protos_proto_rawDescData = file_protos_protos_proto_rawDesc
)

func file_protos_protos_proto_rawDescGZIP() []byte {
	file_protos_protos_proto_rawDescOnce.Do(func() {
		file_protos_protos_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_protos_proto_rawDescData)
	})
	return file_protos_protos_proto_rawDescData
}

var file_protos_protos_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protos_protos_proto_goTypes = []interface{}{
	(*MessageRequest)(nil),          // 0: protos.MessageRequest
	(*MessageReply)(nil),            // 1: protos.MessageReply
	(*CoordinacionRequest)(nil),     // 2: protos.CoordinacionRequest
	(*CoordinacionReply)(nil),       // 3: protos.CoordinacionReply
	(*ReestructuracionRequest)(nil), // 4: protos.ReestructuracionRequest
	(*ReestructuracionReply)(nil),   // 5: protos.ReestructuracionReply
	(*RelojRequest)(nil),            // 6: protos.RelojRequest
	(*RelojReply)(nil),              // 7: protos.RelojReply
}
var file_protos_protos_proto_depIdxs = []int32{
	0, // 0: protos.ManejoComunicacion.Comunicar:input_type -> protos.MessageRequest
	2, // 1: protos.ManejoComunicacion.Coordinar:input_type -> protos.CoordinacionRequest
	4, // 2: protos.ManejoComunicacion.Reestructurar:input_type -> protos.ReestructuracionRequest
	6, // 3: protos.ManejoComunicacion.ConsultarReloj:input_type -> protos.RelojRequest
	1, // 4: protos.ManejoComunicacion.Comunicar:output_type -> protos.MessageReply
	3, // 5: protos.ManejoComunicacion.Coordinar:output_type -> protos.CoordinacionReply
	5, // 6: protos.ManejoComunicacion.Reestructurar:output_type -> protos.ReestructuracionReply
	7, // 7: protos.ManejoComunicacion.ConsultarReloj:output_type -> protos.RelojReply
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_protos_proto_init() }
func file_protos_protos_proto_init() {
	if File_protos_protos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_protos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageRequest); i {
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
		file_protos_protos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageReply); i {
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
		file_protos_protos_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoordinacionRequest); i {
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
		file_protos_protos_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoordinacionReply); i {
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
		file_protos_protos_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReestructuracionRequest); i {
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
		file_protos_protos_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReestructuracionReply); i {
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
		file_protos_protos_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelojRequest); i {
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
		file_protos_protos_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelojReply); i {
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
			RawDescriptor: file_protos_protos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_protos_proto_goTypes,
		DependencyIndexes: file_protos_protos_proto_depIdxs,
		MessageInfos:      file_protos_protos_proto_msgTypes,
	}.Build()
	File_protos_protos_proto = out.File
	file_protos_protos_proto_rawDesc = nil
	file_protos_protos_proto_goTypes = nil
	file_protos_protos_proto_depIdxs = nil
}
