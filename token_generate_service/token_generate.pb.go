// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: token_generate.proto

package token_generate_service

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

type ERR_CODE_TOKEN_GENERATE_SERVICE int32

const (
	ERR_CODE_TOKEN_GENERATE_SERVICE_CODE_NO_ERROR                ERR_CODE_TOKEN_GENERATE_SERVICE = 0
	ERR_CODE_TOKEN_GENERATE_SERVICE_CODE_ERR_MISSING_PARAM       ERR_CODE_TOKEN_GENERATE_SERVICE = -1
	ERR_CODE_TOKEN_GENERATE_SERVICE_CODE_ERR_INVALID_PARAM       ERR_CODE_TOKEN_GENERATE_SERVICE = -2
	ERR_CODE_TOKEN_GENERATE_SERVICE_CODE_ERR_INVALID_PRIVATE_KEY ERR_CODE_TOKEN_GENERATE_SERVICE = -3
)

// Enum value maps for ERR_CODE_TOKEN_GENERATE_SERVICE.
var (
	ERR_CODE_TOKEN_GENERATE_SERVICE_name = map[int32]string{
		0:  "CODE_NO_ERROR",
		-1: "CODE_ERR_MISSING_PARAM",
		-2: "CODE_ERR_INVALID_PARAM",
		-3: "CODE_ERR_INVALID_PRIVATE_KEY",
	}
	ERR_CODE_TOKEN_GENERATE_SERVICE_value = map[string]int32{
		"CODE_NO_ERROR":                0,
		"CODE_ERR_MISSING_PARAM":       -1,
		"CODE_ERR_INVALID_PARAM":       -2,
		"CODE_ERR_INVALID_PRIVATE_KEY": -3,
	}
)

func (x ERR_CODE_TOKEN_GENERATE_SERVICE) Enum() *ERR_CODE_TOKEN_GENERATE_SERVICE {
	p := new(ERR_CODE_TOKEN_GENERATE_SERVICE)
	*p = x
	return p
}

func (x ERR_CODE_TOKEN_GENERATE_SERVICE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ERR_CODE_TOKEN_GENERATE_SERVICE) Descriptor() protoreflect.EnumDescriptor {
	return file_token_generate_proto_enumTypes[0].Descriptor()
}

func (ERR_CODE_TOKEN_GENERATE_SERVICE) Type() protoreflect.EnumType {
	return &file_token_generate_proto_enumTypes[0]
}

func (x ERR_CODE_TOKEN_GENERATE_SERVICE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ERR_CODE_TOKEN_GENERATE_SERVICE.Descriptor instead.
func (ERR_CODE_TOKEN_GENERATE_SERVICE) EnumDescriptor() ([]byte, []int) {
	return file_token_generate_proto_rawDescGZIP(), []int{0}
}

type GenerateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletAccount string `protobuf:"bytes,1,opt,name=wallet_account,json=walletAccount,proto3" json:"wallet_account,omitempty"`
	Yid           string `protobuf:"bytes,2,opt,name=yid,proto3" json:"yid,omitempty"`
}

func (x *GenerateTokenRequest) Reset() {
	*x = GenerateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_generate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenRequest) ProtoMessage() {}

func (x *GenerateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_token_generate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenRequest.ProtoReflect.Descriptor instead.
func (*GenerateTokenRequest) Descriptor() ([]byte, []int) {
	return file_token_generate_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateTokenRequest) GetWalletAccount() string {
	if x != nil {
		return x.WalletAccount
	}
	return ""
}

func (x *GenerateTokenRequest) GetYid() string {
	if x != nil {
		return x.Yid
	}
	return ""
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssociateAccount string `protobuf:"bytes,1,opt,name=associate_account,json=associateAccount,proto3" json:"associate_account,omitempty"`
	Yid              string `protobuf:"bytes,2,opt,name=yid,proto3" json:"yid,omitempty"`
	ExpireHeight     int64  `protobuf:"varint,3,opt,name=expire_height,json=expireHeight,proto3" json:"expire_height,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_generate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_token_generate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_token_generate_proto_rawDescGZIP(), []int{1}
}

func (x *Token) GetAssociateAccount() string {
	if x != nil {
		return x.AssociateAccount
	}
	return ""
}

func (x *Token) GetYid() string {
	if x != nil {
		return x.Yid
	}
	return ""
}

func (x *Token) GetExpireHeight() int64 {
	if x != nil {
		return x.ExpireHeight
	}
	return 0
}

type GenerateTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header         *TokenGenerateRspHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Token          *Token                  `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	TokenSignature []byte                  `protobuf:"bytes,3,opt,name=token_signature,json=tokenSignature,proto3" json:"token_signature,omitempty"`
}

func (x *GenerateTokenResponse) Reset() {
	*x = GenerateTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_generate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokenResponse) ProtoMessage() {}

func (x *GenerateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_token_generate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokenResponse.ProtoReflect.Descriptor instead.
func (*GenerateTokenResponse) Descriptor() ([]byte, []int) {
	return file_token_generate_proto_rawDescGZIP(), []int{2}
}

func (x *GenerateTokenResponse) GetHeader() *TokenGenerateRspHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GenerateTokenResponse) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *GenerateTokenResponse) GetTokenSignature() []byte {
	if x != nil {
		return x.TokenSignature
	}
	return nil
}

type TokenGenerateRspHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret    int32  `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *TokenGenerateRspHeader) Reset() {
	*x = TokenGenerateRspHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_generate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenGenerateRspHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenGenerateRspHeader) ProtoMessage() {}

func (x *TokenGenerateRspHeader) ProtoReflect() protoreflect.Message {
	mi := &file_token_generate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenGenerateRspHeader.ProtoReflect.Descriptor instead.
func (*TokenGenerateRspHeader) Descriptor() ([]byte, []int) {
	return file_token_generate_proto_rawDescGZIP(), []int{3}
}

func (x *TokenGenerateRspHeader) GetRet() int32 {
	if x != nil {
		return x.Ret
	}
	return 0
}

func (x *TokenGenerateRspHeader) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

var File_token_generate_proto protoreflect.FileDescriptor

var file_token_generate_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x79, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x79, 0x69, 0x64, 0x22, 0x6b, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x2b, 0x0a, 0x11, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x73, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x79, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x79, 0x69, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x48, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x22, 0x8f, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f,
	0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x73,
	0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x1c, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06,
	0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x27, 0x0a,
	0x0f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x42, 0x0a, 0x16, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72,
	0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2a, 0xa9, 0x01, 0x0a, 0x1f, 0x45,
	0x52, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x47, 0x45,
	0x4e, 0x45, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x12, 0x11,
	0x0a, 0x0d, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x00, 0x12, 0x23, 0x0a, 0x16, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x4d, 0x49,
	0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x10, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x23, 0x0a, 0x16, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45,
	0x52, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d,
	0x10, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x29, 0x0a, 0x1c, 0x43,
	0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x5f, 0x4b, 0x45, 0x59, 0x10, 0xfd, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x32, 0x56, 0x0a, 0x14, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e,
	0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x15, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x79,
	0x2d, 0x77, 0x65, 0x62, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_token_generate_proto_rawDescOnce sync.Once
	file_token_generate_proto_rawDescData = file_token_generate_proto_rawDesc
)

func file_token_generate_proto_rawDescGZIP() []byte {
	file_token_generate_proto_rawDescOnce.Do(func() {
		file_token_generate_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_generate_proto_rawDescData)
	})
	return file_token_generate_proto_rawDescData
}

var file_token_generate_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_token_generate_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_token_generate_proto_goTypes = []interface{}{
	(ERR_CODE_TOKEN_GENERATE_SERVICE)(0), // 0: ERR_CODE_TOKEN_GENERATE_SERVICE
	(*GenerateTokenRequest)(nil),         // 1: GenerateTokenRequest
	(*Token)(nil),                        // 2: Token
	(*GenerateTokenResponse)(nil),        // 3: GenerateTokenResponse
	(*TokenGenerateRspHeader)(nil),       // 4: TokenGenerateRspHeader
}
var file_token_generate_proto_depIdxs = []int32{
	4, // 0: GenerateTokenResponse.header:type_name -> TokenGenerateRspHeader
	2, // 1: GenerateTokenResponse.token:type_name -> Token
	1, // 2: TokenGenerateService.GenerateToken:input_type -> GenerateTokenRequest
	3, // 3: TokenGenerateService.GenerateToken:output_type -> GenerateTokenResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_token_generate_proto_init() }
func file_token_generate_proto_init() {
	if File_token_generate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_token_generate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTokenRequest); i {
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
		file_token_generate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_token_generate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTokenResponse); i {
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
		file_token_generate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenGenerateRspHeader); i {
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
			RawDescriptor: file_token_generate_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_token_generate_proto_goTypes,
		DependencyIndexes: file_token_generate_proto_depIdxs,
		EnumInfos:         file_token_generate_proto_enumTypes,
		MessageInfos:      file_token_generate_proto_msgTypes,
	}.Build()
	File_token_generate_proto = out.File
	file_token_generate_proto_rawDesc = nil
	file_token_generate_proto_goTypes = nil
	file_token_generate_proto_depIdxs = nil
}
