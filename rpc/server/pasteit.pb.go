// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.0
// source: api/protobuf/pasteit.proto

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

type CreatePasteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string            `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string            `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Headers map[string]string `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Ttl     uint64            `protobuf:"varint,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *CreatePasteRequest) Reset() {
	*x = CreatePasteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_pasteit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePasteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePasteRequest) ProtoMessage() {}

func (x *CreatePasteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_pasteit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePasteRequest.ProtoReflect.Descriptor instead.
func (*CreatePasteRequest) Descriptor() ([]byte, []int) {
	return file_api_protobuf_pasteit_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePasteRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreatePasteRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreatePasteRequest) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *CreatePasteRequest) GetTtl() uint64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type CreatePasteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	CreatedAt uint64 `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *CreatePasteResponse) Reset() {
	*x = CreatePasteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_pasteit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePasteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePasteResponse) ProtoMessage() {}

func (x *CreatePasteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_pasteit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePasteResponse.ProtoReflect.Descriptor instead.
func (*CreatePasteResponse) Descriptor() ([]byte, []int) {
	return file_api_protobuf_pasteit_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePasteResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *CreatePasteResponse) GetCreatedAt() uint64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type GetPasteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPasteRequest) Reset() {
	*x = GetPasteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_pasteit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPasteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPasteRequest) ProtoMessage() {}

func (x *GetPasteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_pasteit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPasteRequest.ProtoReflect.Descriptor instead.
func (*GetPasteRequest) Descriptor() ([]byte, []int) {
	return file_api_protobuf_pasteit_proto_rawDescGZIP(), []int{2}
}

func (x *GetPasteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPasteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string            `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string            `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Headers map[string]string `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetPasteResponse) Reset() {
	*x = GetPasteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_pasteit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPasteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPasteResponse) ProtoMessage() {}

func (x *GetPasteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_pasteit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPasteResponse.ProtoReflect.Descriptor instead.
func (*GetPasteResponse) Descriptor() ([]byte, []int) {
	return file_api_protobuf_pasteit_proto_rawDescGZIP(), []int{3}
}

func (x *GetPasteResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetPasteResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *GetPasteResponse) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

var File_api_protobuf_pasteit_proto protoreflect.FileDescriptor

var file_api_protobuf_pasteit_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70,
	0x61, 0x73, 0x74, 0x65, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x61,
	0x73, 0x74, 0x65, 0x69, 0x74, 0x22, 0xd6, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x73, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x07,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x70, 0x61, 0x73, 0x74, 0x65, 0x69, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x73, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x74,
	0x74, 0x6c, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x47,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x73, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xc0, 0x01, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x73, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x40, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x70, 0x61, 0x73, 0x74, 0x65, 0x69, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x73, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x93, 0x01,
	0x0a, 0x06, 0x50, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x73, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x70, 0x61, 0x73, 0x74, 0x65, 0x69,
	0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x61, 0x73, 0x74, 0x65, 0x69, 0x74, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x61, 0x73, 0x74, 0x65, 0x12, 0x18,
	0x2e, 0x70, 0x61, 0x73, 0x74, 0x65, 0x69, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x73, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x61, 0x73, 0x74, 0x65,
	0x69, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x73, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_protobuf_pasteit_proto_rawDescOnce sync.Once
	file_api_protobuf_pasteit_proto_rawDescData = file_api_protobuf_pasteit_proto_rawDesc
)

func file_api_protobuf_pasteit_proto_rawDescGZIP() []byte {
	file_api_protobuf_pasteit_proto_rawDescOnce.Do(func() {
		file_api_protobuf_pasteit_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_protobuf_pasteit_proto_rawDescData)
	})
	return file_api_protobuf_pasteit_proto_rawDescData
}

var file_api_protobuf_pasteit_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_protobuf_pasteit_proto_goTypes = []interface{}{
	(*CreatePasteRequest)(nil),  // 0: pasteit.CreatePasteRequest
	(*CreatePasteResponse)(nil), // 1: pasteit.CreatePasteResponse
	(*GetPasteRequest)(nil),     // 2: pasteit.GetPasteRequest
	(*GetPasteResponse)(nil),    // 3: pasteit.GetPasteResponse
	nil,                         // 4: pasteit.CreatePasteRequest.HeadersEntry
	nil,                         // 5: pasteit.GetPasteResponse.HeadersEntry
}
var file_api_protobuf_pasteit_proto_depIdxs = []int32{
	4, // 0: pasteit.CreatePasteRequest.headers:type_name -> pasteit.CreatePasteRequest.HeadersEntry
	5, // 1: pasteit.GetPasteResponse.headers:type_name -> pasteit.GetPasteResponse.HeadersEntry
	0, // 2: pasteit.Paster.CreatePaste:input_type -> pasteit.CreatePasteRequest
	2, // 3: pasteit.Paster.GetPaste:input_type -> pasteit.GetPasteRequest
	1, // 4: pasteit.Paster.CreatePaste:output_type -> pasteit.CreatePasteResponse
	3, // 5: pasteit.Paster.GetPaste:output_type -> pasteit.GetPasteResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_protobuf_pasteit_proto_init() }
func file_api_protobuf_pasteit_proto_init() {
	if File_api_protobuf_pasteit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_protobuf_pasteit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePasteRequest); i {
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
		file_api_protobuf_pasteit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePasteResponse); i {
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
		file_api_protobuf_pasteit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPasteRequest); i {
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
		file_api_protobuf_pasteit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPasteResponse); i {
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
			RawDescriptor: file_api_protobuf_pasteit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_protobuf_pasteit_proto_goTypes,
		DependencyIndexes: file_api_protobuf_pasteit_proto_depIdxs,
		MessageInfos:      file_api_protobuf_pasteit_proto_msgTypes,
	}.Build()
	File_api_protobuf_pasteit_proto = out.File
	file_api_protobuf_pasteit_proto_rawDesc = nil
	file_api_protobuf_pasteit_proto_goTypes = nil
	file_api_protobuf_pasteit_proto_depIdxs = nil
}
