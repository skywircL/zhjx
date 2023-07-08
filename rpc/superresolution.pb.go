// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: rpc/superresolution.proto

package rpc

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

type SuperResolutionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageUri string `protobuf:"bytes,1,opt,name=image_uri,json=imageUri,proto3" json:"image_uri,omitempty"`
}

func (x *SuperResolutionRequest) Reset() {
	*x = SuperResolutionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_superresolution_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuperResolutionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuperResolutionRequest) ProtoMessage() {}

func (x *SuperResolutionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_superresolution_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuperResolutionRequest.ProtoReflect.Descriptor instead.
func (*SuperResolutionRequest) Descriptor() ([]byte, []int) {
	return file_rpc_superresolution_proto_rawDescGZIP(), []int{0}
}

func (x *SuperResolutionRequest) GetImageUri() string {
	if x != nil {
		return x.ImageUri
	}
	return ""
}

type SuperResolutionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SuperResolutionResponse) Reset() {
	*x = SuperResolutionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_superresolution_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuperResolutionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuperResolutionResponse) ProtoMessage() {}

func (x *SuperResolutionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_superresolution_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuperResolutionResponse.ProtoReflect.Descriptor instead.
func (*SuperResolutionResponse) Descriptor() ([]byte, []int) {
	return file_rpc_superresolution_proto_rawDescGZIP(), []int{1}
}

func (x *SuperResolutionResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *SuperResolutionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PersonBankRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PersonBankRequest) Reset() {
	*x = PersonBankRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_superresolution_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonBankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonBankRequest) ProtoMessage() {}

func (x *PersonBankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_superresolution_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonBankRequest.ProtoReflect.Descriptor instead.
func (*PersonBankRequest) Descriptor() ([]byte, []int) {
	return file_rpc_superresolution_proto_rawDescGZIP(), []int{2}
}

type PersonBankResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PersonBankResponse) Reset() {
	*x = PersonBankResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_superresolution_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonBankResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonBankResponse) ProtoMessage() {}

func (x *PersonBankResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_superresolution_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonBankResponse.ProtoReflect.Descriptor instead.
func (*PersonBankResponse) Descriptor() ([]byte, []int) {
	return file_rpc_superresolution_proto_rawDescGZIP(), []int{3}
}

func (x *PersonBankResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *PersonBankResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_rpc_superresolution_proto protoreflect.FileDescriptor

var file_rpc_superresolution_proto_rawDesc = []byte{
	0x0a, 0x19, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x16, 0x53,
	0x75, 0x70, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75,
	0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55,
	0x72, 0x69, 0x22, 0x49, 0x0a, 0x17, 0x53, 0x75, 0x70, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x13, 0x0a,
	0x11, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x44, 0x0a, 0x12, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x42, 0x61, 0x6e, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x96, 0x01, 0x0a, 0x0f, 0x53, 0x75, 0x70,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x13,
	0x53, 0x75, 0x70, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x46,
	0x75, 0x6e, 0x63, 0x12, 0x17, 0x2e, 0x53, 0x75, 0x70, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x53,
	0x75, 0x70, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x12, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x42,
	0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_rpc_superresolution_proto_rawDescOnce sync.Once
	file_rpc_superresolution_proto_rawDescData = file_rpc_superresolution_proto_rawDesc
)

func file_rpc_superresolution_proto_rawDescGZIP() []byte {
	file_rpc_superresolution_proto_rawDescOnce.Do(func() {
		file_rpc_superresolution_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_superresolution_proto_rawDescData)
	})
	return file_rpc_superresolution_proto_rawDescData
}

var file_rpc_superresolution_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rpc_superresolution_proto_goTypes = []interface{}{
	(*SuperResolutionRequest)(nil),  // 0: SuperResolutionRequest
	(*SuperResolutionResponse)(nil), // 1: SuperResolutionResponse
	(*PersonBankRequest)(nil),       // 2: PersonBankRequest
	(*PersonBankResponse)(nil),      // 3: PersonBankResponse
}
var file_rpc_superresolution_proto_depIdxs = []int32{
	0, // 0: SuperResolution.SuperResolutionFunc:input_type -> SuperResolutionRequest
	2, // 1: SuperResolution.PersonBank:input_type -> PersonBankRequest
	1, // 2: SuperResolution.SuperResolutionFunc:output_type -> SuperResolutionResponse
	3, // 3: SuperResolution.PersonBank:output_type -> PersonBankResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_superresolution_proto_init() }
func file_rpc_superresolution_proto_init() {
	if File_rpc_superresolution_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_superresolution_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuperResolutionRequest); i {
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
		file_rpc_superresolution_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuperResolutionResponse); i {
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
		file_rpc_superresolution_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonBankRequest); i {
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
		file_rpc_superresolution_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonBankResponse); i {
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
			RawDescriptor: file_rpc_superresolution_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_superresolution_proto_goTypes,
		DependencyIndexes: file_rpc_superresolution_proto_depIdxs,
		MessageInfos:      file_rpc_superresolution_proto_msgTypes,
	}.Build()
	File_rpc_superresolution_proto = out.File
	file_rpc_superresolution_proto_rawDesc = nil
	file_rpc_superresolution_proto_goTypes = nil
	file_rpc_superresolution_proto_depIdxs = nil
}
