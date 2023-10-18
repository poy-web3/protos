// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: login.proto

package login_service

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

type ERR_CODE_LOGIN_SERVICE int32

const (
	ERR_CODE_LOGIN_SERVICE_CODE_NO_ERROR_LOGIN_SERVICE          ERR_CODE_LOGIN_SERVICE = 0
	ERR_CODE_LOGIN_SERVICE_CODE_ERR_MISSING_PARAM_LOGIN_SERVICE ERR_CODE_LOGIN_SERVICE = -1
	ERR_CODE_LOGIN_SERVICE_CODE_ERR_INVALID_PARAM_LOGIN_SERVICE ERR_CODE_LOGIN_SERVICE = -2
	ERR_CODE_LOGIN_SERVICE_CODE_ERR_UNKNOWN_LOGIN_SERVICE       ERR_CODE_LOGIN_SERVICE = -1001
)

// Enum value maps for ERR_CODE_LOGIN_SERVICE.
var (
	ERR_CODE_LOGIN_SERVICE_name = map[int32]string{
		0:     "CODE_NO_ERROR_LOGIN_SERVICE",
		-1:    "CODE_ERR_MISSING_PARAM_LOGIN_SERVICE",
		-2:    "CODE_ERR_INVALID_PARAM_LOGIN_SERVICE",
		-1001: "CODE_ERR_UNKNOWN_LOGIN_SERVICE",
	}
	ERR_CODE_LOGIN_SERVICE_value = map[string]int32{
		"CODE_NO_ERROR_LOGIN_SERVICE":          0,
		"CODE_ERR_MISSING_PARAM_LOGIN_SERVICE": -1,
		"CODE_ERR_INVALID_PARAM_LOGIN_SERVICE": -2,
		"CODE_ERR_UNKNOWN_LOGIN_SERVICE":       -1001,
	}
)

func (x ERR_CODE_LOGIN_SERVICE) Enum() *ERR_CODE_LOGIN_SERVICE {
	p := new(ERR_CODE_LOGIN_SERVICE)
	*p = x
	return p
}

func (x ERR_CODE_LOGIN_SERVICE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ERR_CODE_LOGIN_SERVICE) Descriptor() protoreflect.EnumDescriptor {
	return file_login_proto_enumTypes[0].Descriptor()
}

func (ERR_CODE_LOGIN_SERVICE) Type() protoreflect.EnumType {
	return &file_login_proto_enumTypes[0]
}

func (x ERR_CODE_LOGIN_SERVICE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ERR_CODE_LOGIN_SERVICE.Descriptor instead.
func (ERR_CODE_LOGIN_SERVICE) EnumDescriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{0}
}

type BioKeyLoginRequset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FacialImage []byte `protobuf:"bytes,1,opt,name=facial_image,json=facialImage,proto3" json:"facial_image,omitempty"`
	PublicKey   string `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *BioKeyLoginRequset) Reset() {
	*x = BioKeyLoginRequset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BioKeyLoginRequset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BioKeyLoginRequset) ProtoMessage() {}

func (x *BioKeyLoginRequset) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BioKeyLoginRequset.ProtoReflect.Descriptor instead.
func (*BioKeyLoginRequset) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{0}
}

func (x *BioKeyLoginRequset) GetFacialImage() []byte {
	if x != nil {
		return x.FacialImage
	}
	return nil
}

func (x *BioKeyLoginRequset) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

type BioKeyLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *CommonRspHeaderLoginService `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Token  string                       `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *BioKeyLoginResponse) Reset() {
	*x = BioKeyLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BioKeyLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BioKeyLoginResponse) ProtoMessage() {}

func (x *BioKeyLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BioKeyLoginResponse.ProtoReflect.Descriptor instead.
func (*BioKeyLoginResponse) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{1}
}

func (x *BioKeyLoginResponse) GetHeader() *CommonRspHeaderLoginService {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *BioKeyLoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CommonRspHeaderLoginService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret    int32  `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *CommonRspHeaderLoginService) Reset() {
	*x = CommonRspHeaderLoginService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonRspHeaderLoginService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonRspHeaderLoginService) ProtoMessage() {}

func (x *CommonRspHeaderLoginService) ProtoReflect() protoreflect.Message {
	mi := &file_login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonRspHeaderLoginService.ProtoReflect.Descriptor instead.
func (*CommonRspHeaderLoginService) Descriptor() ([]byte, []int) {
	return file_login_proto_rawDescGZIP(), []int{2}
}

func (x *CommonRspHeaderLoginService) GetRet() int32 {
	if x != nil {
		return x.Ret
	}
	return 0
}

func (x *CommonRspHeaderLoginService) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

var File_login_proto protoreflect.FileDescriptor

var file_login_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a,
	0x12, 0x42, 0x69, 0x6f, 0x4b, 0x65, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x73, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x61, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x66, 0x61, 0x63, 0x69, 0x61,
	0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x61, 0x0a, 0x13, 0x42, 0x69, 0x6f, 0x4b, 0x65, 0x79, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x47, 0x0a, 0x1b, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x52, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x2a, 0xcc, 0x01, 0x0a, 0x16, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4c,
	0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x12, 0x1f, 0x0a, 0x1b,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x4c, 0x4f,
	0x47, 0x49, 0x4e, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x00, 0x12, 0x31, 0x0a,
	0x24, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e,
	0x47, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x53, 0x45,
	0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01,
	0x12, 0x31, 0x0a, 0x24, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e,
	0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0x01, 0x12, 0x2b, 0x0a, 0x1e, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x53, 0x45,
	0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x97, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01,
	0x32, 0x4c, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3c, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x42, 0x69, 0x6f,
	0x4b, 0x65, 0x79, 0x12, 0x13, 0x2e, 0x42, 0x69, 0x6f, 0x4b, 0x65, 0x79, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x73, 0x65, 0x74, 0x1a, 0x14, 0x2e, 0x42, 0x69, 0x6f, 0x4b, 0x65,
	0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11,
	0x5a, 0x0f, 0x2e, 0x3b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_login_proto_rawDescOnce sync.Once
	file_login_proto_rawDescData = file_login_proto_rawDesc
)

func file_login_proto_rawDescGZIP() []byte {
	file_login_proto_rawDescOnce.Do(func() {
		file_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_login_proto_rawDescData)
	})
	return file_login_proto_rawDescData
}

var file_login_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_login_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_login_proto_goTypes = []interface{}{
	(ERR_CODE_LOGIN_SERVICE)(0),         // 0: ERR_CODE_LOGIN_SERVICE
	(*BioKeyLoginRequset)(nil),          // 1: BioKeyLoginRequset
	(*BioKeyLoginResponse)(nil),         // 2: BioKeyLoginResponse
	(*CommonRspHeaderLoginService)(nil), // 3: CommonRspHeaderLoginService
}
var file_login_proto_depIdxs = []int32{
	3, // 0: BioKeyLoginResponse.header:type_name -> CommonRspHeaderLoginService
	1, // 1: LoginService.LoginWithBioKey:input_type -> BioKeyLoginRequset
	2, // 2: LoginService.LoginWithBioKey:output_type -> BioKeyLoginResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_login_proto_init() }
func file_login_proto_init() {
	if File_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BioKeyLoginRequset); i {
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
		file_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BioKeyLoginResponse); i {
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
		file_login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonRspHeaderLoginService); i {
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
			RawDescriptor: file_login_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_login_proto_goTypes,
		DependencyIndexes: file_login_proto_depIdxs,
		EnumInfos:         file_login_proto_enumTypes,
		MessageInfos:      file_login_proto_msgTypes,
	}.Build()
	File_login_proto = out.File
	file_login_proto_rawDesc = nil
	file_login_proto_goTypes = nil
	file_login_proto_depIdxs = nil
}
