// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: local_query.proto

package local_query_service

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

type ERR_CODE_LOCAL_QUERY_SERVICE int32

const (
	ERR_CODE_LOCAL_QUERY_SERVICE_CODE_NO_ERROR_QUERY_SERVICE                ERR_CODE_LOCAL_QUERY_SERVICE = 0
	ERR_CODE_LOCAL_QUERY_SERVICE_CODE_ERR_MISSING_PARAM_LOCAL_QUERY_SERVICE ERR_CODE_LOCAL_QUERY_SERVICE = -1
	ERR_CODE_LOCAL_QUERY_SERVICE_CODE_ERR_INVALID_PARAM_LOCAL_QUERY_SERVICE ERR_CODE_LOCAL_QUERY_SERVICE = -2
	ERR_CODE_LOCAL_QUERY_SERVICE_CODE_ERR_UNKNOWN_LOCAL_QUERY_SERVICE       ERR_CODE_LOCAL_QUERY_SERVICE = -1001
)

// Enum value maps for ERR_CODE_LOCAL_QUERY_SERVICE.
var (
	ERR_CODE_LOCAL_QUERY_SERVICE_name = map[int32]string{
		0:     "CODE_NO_ERROR_QUERY_SERVICE",
		-1:    "CODE_ERR_MISSING_PARAM_LOCAL_QUERY_SERVICE",
		-2:    "CODE_ERR_INVALID_PARAM_LOCAL_QUERY_SERVICE",
		-1001: "CODE_ERR_UNKNOWN_LOCAL_QUERY_SERVICE",
	}
	ERR_CODE_LOCAL_QUERY_SERVICE_value = map[string]int32{
		"CODE_NO_ERROR_QUERY_SERVICE":                0,
		"CODE_ERR_MISSING_PARAM_LOCAL_QUERY_SERVICE": -1,
		"CODE_ERR_INVALID_PARAM_LOCAL_QUERY_SERVICE": -2,
		"CODE_ERR_UNKNOWN_LOCAL_QUERY_SERVICE":       -1001,
	}
)

func (x ERR_CODE_LOCAL_QUERY_SERVICE) Enum() *ERR_CODE_LOCAL_QUERY_SERVICE {
	p := new(ERR_CODE_LOCAL_QUERY_SERVICE)
	*p = x
	return p
}

func (x ERR_CODE_LOCAL_QUERY_SERVICE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ERR_CODE_LOCAL_QUERY_SERVICE) Descriptor() protoreflect.EnumDescriptor {
	return file_local_query_proto_enumTypes[0].Descriptor()
}

func (ERR_CODE_LOCAL_QUERY_SERVICE) Type() protoreflect.EnumType {
	return &file_local_query_proto_enumTypes[0]
}

func (x ERR_CODE_LOCAL_QUERY_SERVICE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ERR_CODE_LOCAL_QUERY_SERVICE.Descriptor instead.
func (ERR_CODE_LOCAL_QUERY_SERVICE) EnumDescriptor() ([]byte, []int) {
	return file_local_query_proto_rawDescGZIP(), []int{0}
}

type UserValidityCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey string `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"` // identifier
	Token     string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UserValidityCheckRequest) Reset() {
	*x = UserValidityCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_local_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserValidityCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserValidityCheckRequest) ProtoMessage() {}

func (x *UserValidityCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_local_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserValidityCheckRequest.ProtoReflect.Descriptor instead.
func (*UserValidityCheckRequest) Descriptor() ([]byte, []int) {
	return file_local_query_proto_rawDescGZIP(), []int{0}
}

func (x *UserValidityCheckRequest) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *UserValidityCheckRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserValidityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *CommonRspHeaderLocalQueryService `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Valid  bool                              `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *UserValidityResponse) Reset() {
	*x = UserValidityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_local_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserValidityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserValidityResponse) ProtoMessage() {}

func (x *UserValidityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_local_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserValidityResponse.ProtoReflect.Descriptor instead.
func (*UserValidityResponse) Descriptor() ([]byte, []int) {
	return file_local_query_proto_rawDescGZIP(), []int{1}
}

func (x *UserValidityResponse) GetHeader() *CommonRspHeaderLocalQueryService {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *UserValidityResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

type CommonRspHeaderLocalQueryService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret    int32  `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *CommonRspHeaderLocalQueryService) Reset() {
	*x = CommonRspHeaderLocalQueryService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_local_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonRspHeaderLocalQueryService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonRspHeaderLocalQueryService) ProtoMessage() {}

func (x *CommonRspHeaderLocalQueryService) ProtoReflect() protoreflect.Message {
	mi := &file_local_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonRspHeaderLocalQueryService.ProtoReflect.Descriptor instead.
func (*CommonRspHeaderLocalQueryService) Descriptor() ([]byte, []int) {
	return file_local_query_proto_rawDescGZIP(), []int{2}
}

func (x *CommonRspHeaderLocalQueryService) GetRet() int32 {
	if x != nil {
		return x.Ret
	}
	return 0
}

func (x *CommonRspHeaderLocalQueryService) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

var File_local_query_proto protoreflect.FileDescriptor

var file_local_query_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x18, 0x55, 0x73, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x67, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x22, 0x4c, 0x0a,
	0x20, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x72, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2a, 0xe4, 0x01, 0x0a, 0x1c,
	0x45, 0x52, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x5f, 0x51,
	0x55, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x12, 0x1f, 0x0a, 0x1b,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x51, 0x55,
	0x45, 0x52, 0x59, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x00, 0x12, 0x37, 0x0a,
	0x2a, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e,
	0x47, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x5f, 0x51, 0x55,
	0x45, 0x52, 0x59, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x37, 0x0a, 0x2a, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45,
	0x52, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d,
	0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x5f, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x45, 0x52,
	0x56, 0x49, 0x43, 0x45, 0x10, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12,
	0x31, 0x0a, 0x24, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x5f, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x97, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0x01, 0x32, 0x5a, 0x0a, 0x11, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x19, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30,
	0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x79,
	0x2d, 0x77, 0x65, 0x62, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_local_query_proto_rawDescOnce sync.Once
	file_local_query_proto_rawDescData = file_local_query_proto_rawDesc
)

func file_local_query_proto_rawDescGZIP() []byte {
	file_local_query_proto_rawDescOnce.Do(func() {
		file_local_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_local_query_proto_rawDescData)
	})
	return file_local_query_proto_rawDescData
}

var file_local_query_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_local_query_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_local_query_proto_goTypes = []interface{}{
	(ERR_CODE_LOCAL_QUERY_SERVICE)(0),        // 0: ERR_CODE_LOCAL_QUERY_SERVICE
	(*UserValidityCheckRequest)(nil),         // 1: UserValidityCheckRequest
	(*UserValidityResponse)(nil),             // 2: UserValidityResponse
	(*CommonRspHeaderLocalQueryService)(nil), // 3: CommonRspHeaderLocalQueryService
}
var file_local_query_proto_depIdxs = []int32{
	3, // 0: UserValidityResponse.header:type_name -> CommonRspHeaderLocalQueryService
	1, // 1: LocalQueryService.UserValidityCheck:input_type -> UserValidityCheckRequest
	2, // 2: LocalQueryService.UserValidityCheck:output_type -> UserValidityResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_local_query_proto_init() }
func file_local_query_proto_init() {
	if File_local_query_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_local_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserValidityCheckRequest); i {
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
		file_local_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserValidityResponse); i {
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
		file_local_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonRspHeaderLocalQueryService); i {
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
			RawDescriptor: file_local_query_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_local_query_proto_goTypes,
		DependencyIndexes: file_local_query_proto_depIdxs,
		EnumInfos:         file_local_query_proto_enumTypes,
		MessageInfos:      file_local_query_proto_msgTypes,
	}.Build()
	File_local_query_proto = out.File
	file_local_query_proto_rawDesc = nil
	file_local_query_proto_goTypes = nil
	file_local_query_proto_depIdxs = nil
}