// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.0
// source: rpc/pasteit.proto

package server

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PasteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string            `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string            `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Headers map[string]string `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Ttl     int64             `protobuf:"varint,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *PasteRequest) Reset() {
	*x = PasteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_pasteit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasteRequest) ProtoMessage() {}

func (x *PasteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_pasteit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasteRequest.ProtoReflect.Descriptor instead.
func (*PasteRequest) Descriptor() ([]byte, []int) {
	return file_rpc_pasteit_proto_rawDescGZIP(), []int{0}
}

func (x *PasteRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PasteRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PasteRequest) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *PasteRequest) GetTtl() int64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type PasteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	CreatedAt int64  `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *PasteResponse) Reset() {
	*x = PasteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_pasteit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasteResponse) ProtoMessage() {}

func (x *PasteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_pasteit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasteResponse.ProtoReflect.Descriptor instead.
func (*PasteResponse) Descriptor() ([]byte, []int) {
	return file_rpc_pasteit_proto_rawDescGZIP(), []int{1}
}

func (x *PasteResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *PasteResponse) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

var File_rpc_pasteit_proto protoreflect.FileDescriptor

var file_rpc_pasteit_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x61, 0x73, 0x74, 0x65, 0x69, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x61, 0x73, 0x74, 0x65, 0x69, 0x74, 0x22, 0xca, 0x01, 0x0a,
	0x0c, 0x50, 0x61, 0x73, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x3c, 0x0a,
	0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x70, 0x61, 0x73, 0x74, 0x65, 0x69, 0x74, 0x2e, 0x50, 0x61, 0x73, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x74, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x1a, 0x3a, 0x0a,
	0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x41, 0x0a, 0x0d, 0x50, 0x61, 0x73,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x40, 0x0a, 0x06,
	0x50, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x05, 0x50, 0x61, 0x73, 0x74, 0x65, 0x12,
	0x15, 0x2e, 0x70, 0x61, 0x73, 0x74, 0x65, 0x69, 0x74, 0x2e, 0x50, 0x61, 0x73, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x61, 0x73, 0x74, 0x65, 0x69, 0x74,
	0x2e, 0x50, 0x61, 0x73, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c,
	0x5a, 0x0a, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_pasteit_proto_rawDescOnce sync.Once
	file_rpc_pasteit_proto_rawDescData = file_rpc_pasteit_proto_rawDesc
)

func file_rpc_pasteit_proto_rawDescGZIP() []byte {
	file_rpc_pasteit_proto_rawDescOnce.Do(func() {
		file_rpc_pasteit_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_pasteit_proto_rawDescData)
	})
	return file_rpc_pasteit_proto_rawDescData
}

var file_rpc_pasteit_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rpc_pasteit_proto_goTypes = []interface{}{
	(*PasteRequest)(nil),  // 0: pasteit.PasteRequest
	(*PasteResponse)(nil), // 1: pasteit.PasteResponse
	nil,                   // 2: pasteit.PasteRequest.HeadersEntry
}
var file_rpc_pasteit_proto_depIdxs = []int32{
	2, // 0: pasteit.PasteRequest.headers:type_name -> pasteit.PasteRequest.HeadersEntry
	0, // 1: pasteit.Paster.Paste:input_type -> pasteit.PasteRequest
	1, // 2: pasteit.Paster.Paste:output_type -> pasteit.PasteResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_pasteit_proto_init() }
func file_rpc_pasteit_proto_init() {
	if File_rpc_pasteit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_pasteit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasteRequest); i {
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
		file_rpc_pasteit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasteResponse); i {
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
			RawDescriptor: file_rpc_pasteit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_pasteit_proto_goTypes,
		DependencyIndexes: file_rpc_pasteit_proto_depIdxs,
		MessageInfos:      file_rpc_pasteit_proto_msgTypes,
	}.Build()
	File_rpc_pasteit_proto = out.File
	file_rpc_pasteit_proto_rawDesc = nil
	file_rpc_pasteit_proto_goTypes = nil
	file_rpc_pasteit_proto_depIdxs = nil
}
