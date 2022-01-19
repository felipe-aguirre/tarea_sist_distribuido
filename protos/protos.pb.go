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

var File_protos_protos_proto protoreflect.FileDescriptor

var file_protos_protos_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x56, 0x0a,
	0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x75, 0x74,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x75, 0x74, 0x6f, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x22, 0x24, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x51, 0x0a, 0x12, 0x4d,
	0x61, 0x6e, 0x65, 0x6a, 0x6f, 0x43, 0x6f, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x63, 0x69, 0x6f,
	0x6e, 0x12, 0x3b, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x72, 0x12, 0x16,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x4e,
	0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x65, 0x6c,
	0x69, 0x70, 0x65, 0x2d, 0x61, 0x67, 0x75, 0x69, 0x72, 0x72, 0x65, 0x2f, 0x74, 0x61, 0x72, 0x65,
	0x61, 0x5f, 0x73, 0x69, 0x73, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x69,
	0x64, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x3b, 0x74, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x73, 0x69,
	0x73, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x69, 0x64, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_protos_protos_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_protos_proto_goTypes = []interface{}{
	(*MessageRequest)(nil), // 0: protos.MessageRequest
	(*MessageReply)(nil),   // 1: protos.MessageReply
}
var file_protos_protos_proto_depIdxs = []int32{
	0, // 0: protos.ManejoComunicacion.Comunicar:input_type -> protos.MessageRequest
	1, // 1: protos.ManejoComunicacion.Comunicar:output_type -> protos.MessageReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_protos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
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
