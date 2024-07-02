// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: ambassador.proto

package ambassador_service

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

type ERR_CODE_AMBASSADOR_SERVICE int32

const (
	ERR_CODE_AMBASSADOR_SERVICE_NO_ERROR      ERR_CODE_AMBASSADOR_SERVICE = 0
	ERR_CODE_AMBASSADOR_SERVICE_INVALID_PARAM ERR_CODE_AMBASSADOR_SERVICE = -1
)

// Enum value maps for ERR_CODE_AMBASSADOR_SERVICE.
var (
	ERR_CODE_AMBASSADOR_SERVICE_name = map[int32]string{
		0:  "NO_ERROR",
		-1: "INVALID_PARAM",
	}
	ERR_CODE_AMBASSADOR_SERVICE_value = map[string]int32{
		"NO_ERROR":      0,
		"INVALID_PARAM": -1,
	}
)

func (x ERR_CODE_AMBASSADOR_SERVICE) Enum() *ERR_CODE_AMBASSADOR_SERVICE {
	p := new(ERR_CODE_AMBASSADOR_SERVICE)
	*p = x
	return p
}

func (x ERR_CODE_AMBASSADOR_SERVICE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ERR_CODE_AMBASSADOR_SERVICE) Descriptor() protoreflect.EnumDescriptor {
	return file_ambassador_proto_enumTypes[0].Descriptor()
}

func (ERR_CODE_AMBASSADOR_SERVICE) Type() protoreflect.EnumType {
	return &file_ambassador_proto_enumTypes[0]
}

func (x ERR_CODE_AMBASSADOR_SERVICE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ERR_CODE_AMBASSADOR_SERVICE.Descriptor instead.
func (ERR_CODE_AMBASSADOR_SERVICE) EnumDescriptor() ([]byte, []int) {
	return file_ambassador_proto_rawDescGZIP(), []int{0}
}

type AmbassadorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inviter uint32 `protobuf:"varint,1,opt,name=inviter,proto3" json:"inviter,omitempty"`
	Invitee uint32 `protobuf:"varint,2,opt,name=invitee,proto3" json:"invitee,omitempty"`
}

func (x *AmbassadorRequest) Reset() {
	*x = AmbassadorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ambassador_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AmbassadorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AmbassadorRequest) ProtoMessage() {}

func (x *AmbassadorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ambassador_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AmbassadorRequest.ProtoReflect.Descriptor instead.
func (*AmbassadorRequest) Descriptor() ([]byte, []int) {
	return file_ambassador_proto_rawDescGZIP(), []int{0}
}

func (x *AmbassadorRequest) GetInviter() uint32 {
	if x != nil {
		return x.Inviter
	}
	return 0
}

func (x *AmbassadorRequest) GetInvitee() uint32 {
	if x != nil {
		return x.Invitee
	}
	return 0
}

type AmbassadorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *AmbassadorRspHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *AmbassadorResponse) Reset() {
	*x = AmbassadorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ambassador_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AmbassadorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AmbassadorResponse) ProtoMessage() {}

func (x *AmbassadorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ambassador_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AmbassadorResponse.ProtoReflect.Descriptor instead.
func (*AmbassadorResponse) Descriptor() ([]byte, []int) {
	return file_ambassador_proto_rawDescGZIP(), []int{1}
}

func (x *AmbassadorResponse) GetHeader() *AmbassadorRspHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

type AmbassadorRspHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret    int32  `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *AmbassadorRspHeader) Reset() {
	*x = AmbassadorRspHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ambassador_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AmbassadorRspHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AmbassadorRspHeader) ProtoMessage() {}

func (x *AmbassadorRspHeader) ProtoReflect() protoreflect.Message {
	mi := &file_ambassador_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AmbassadorRspHeader.ProtoReflect.Descriptor instead.
func (*AmbassadorRspHeader) Descriptor() ([]byte, []int) {
	return file_ambassador_proto_rawDescGZIP(), []int{2}
}

func (x *AmbassadorRspHeader) GetRet() int32 {
	if x != nil {
		return x.Ret
	}
	return 0
}

func (x *AmbassadorRspHeader) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

var File_ambassador_proto protoreflect.FileDescriptor

var file_ambassador_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x6d, 0x62, 0x61, 0x73, 0x73, 0x61, 0x64, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x47, 0x0a, 0x11, 0x41, 0x6d, 0x62, 0x61, 0x73, 0x73, 0x61, 0x64, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x22, 0x42, 0x0a, 0x12, 0x41,
	0x6d, 0x62, 0x61, 0x73, 0x73, 0x61, 0x64, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2c, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x41, 0x6d, 0x62, 0x61, 0x73, 0x73, 0x61, 0x64, 0x6f, 0x72, 0x52, 0x73,
	0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22,
	0x3f, 0x0a, 0x13, 0x41, 0x6d, 0x62, 0x61, 0x73, 0x73, 0x61, 0x64, 0x6f, 0x72, 0x52, 0x73, 0x70,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x2a, 0x47, 0x0a, 0x1b, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x4d, 0x42,
	0x41, 0x53, 0x53, 0x41, 0x44, 0x4f, 0x52, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x12,
	0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x1a, 0x0a,
	0x0d, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x10, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x32, 0x46, 0x0a, 0x0a, 0x41, 0x6d, 0x62,
	0x61, 0x73, 0x73, 0x61, 0x64, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x41, 0x6d, 0x62, 0x61, 0x73,
	0x73, 0x61, 0x64, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x41,
	0x6d, 0x62, 0x61, 0x73, 0x73, 0x61, 0x64, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x6f, 0x79, 0x2d, 0x77, 0x65, 0x62, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x61, 0x6d, 0x62, 0x61, 0x73, 0x73, 0x61, 0x64, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ambassador_proto_rawDescOnce sync.Once
	file_ambassador_proto_rawDescData = file_ambassador_proto_rawDesc
)

func file_ambassador_proto_rawDescGZIP() []byte {
	file_ambassador_proto_rawDescOnce.Do(func() {
		file_ambassador_proto_rawDescData = protoimpl.X.CompressGZIP(file_ambassador_proto_rawDescData)
	})
	return file_ambassador_proto_rawDescData
}

var file_ambassador_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ambassador_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ambassador_proto_goTypes = []interface{}{
	(ERR_CODE_AMBASSADOR_SERVICE)(0), // 0: ERR_CODE_AMBASSADOR_SERVICE
	(*AmbassadorRequest)(nil),        // 1: AmbassadorRequest
	(*AmbassadorResponse)(nil),       // 2: AmbassadorResponse
	(*AmbassadorRspHeader)(nil),      // 3: AmbassadorRspHeader
}
var file_ambassador_proto_depIdxs = []int32{
	3, // 0: AmbassadorResponse.header:type_name -> AmbassadorRspHeader
	1, // 1: Ambassador.AddConnection:input_type -> AmbassadorRequest
	2, // 2: Ambassador.AddConnection:output_type -> AmbassadorResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ambassador_proto_init() }
func file_ambassador_proto_init() {
	if File_ambassador_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ambassador_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AmbassadorRequest); i {
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
		file_ambassador_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AmbassadorResponse); i {
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
		file_ambassador_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AmbassadorRspHeader); i {
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
			RawDescriptor: file_ambassador_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ambassador_proto_goTypes,
		DependencyIndexes: file_ambassador_proto_depIdxs,
		EnumInfos:         file_ambassador_proto_enumTypes,
		MessageInfos:      file_ambassador_proto_msgTypes,
	}.Build()
	File_ambassador_proto = out.File
	file_ambassador_proto_rawDesc = nil
	file_ambassador_proto_goTypes = nil
	file_ambassador_proto_depIdxs = nil
}