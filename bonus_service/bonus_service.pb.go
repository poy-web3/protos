// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.0
// source: bonus_service.proto

package bonus_service

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

type ERR_CODE_BONUS_SERVICE int32

const (
	ERR_CODE_BONUS_SERVICE_CODE_NO_ERROR          ERR_CODE_BONUS_SERVICE = 0
	ERR_CODE_BONUS_SERVICE_CODE_ERR_MISSING_PARAM ERR_CODE_BONUS_SERVICE = -1
	ERR_CODE_BONUS_SERVICE_CODE_ERR_INVALID_TOKEN ERR_CODE_BONUS_SERVICE = -2
)

// Enum value maps for ERR_CODE_BONUS_SERVICE.
var (
	ERR_CODE_BONUS_SERVICE_name = map[int32]string{
		0:  "CODE_NO_ERROR",
		-1: "CODE_ERR_MISSING_PARAM",
		-2: "CODE_ERR_INVALID_TOKEN",
	}
	ERR_CODE_BONUS_SERVICE_value = map[string]int32{
		"CODE_NO_ERROR":          0,
		"CODE_ERR_MISSING_PARAM": -1,
		"CODE_ERR_INVALID_TOKEN": -2,
	}
)

func (x ERR_CODE_BONUS_SERVICE) Enum() *ERR_CODE_BONUS_SERVICE {
	p := new(ERR_CODE_BONUS_SERVICE)
	*p = x
	return p
}

func (x ERR_CODE_BONUS_SERVICE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ERR_CODE_BONUS_SERVICE) Descriptor() protoreflect.EnumDescriptor {
	return file_bonus_service_proto_enumTypes[0].Descriptor()
}

func (ERR_CODE_BONUS_SERVICE) Type() protoreflect.EnumType {
	return &file_bonus_service_proto_enumTypes[0]
}

func (x ERR_CODE_BONUS_SERVICE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ERR_CODE_BONUS_SERVICE.Descriptor instead.
func (ERR_CODE_BONUS_SERVICE) EnumDescriptor() ([]byte, []int) {
	return file_bonus_service_proto_rawDescGZIP(), []int{0}
}

type BonusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountList       []string `protobuf:"bytes,1,rep,name=account_list,json=accountList,proto3" json:"account_list,omitempty"`
	MoneyAllocateList []int32  `protobuf:"varint,2,rep,packed,name=money_allocate_list,json=moneyAllocateList,proto3" json:"money_allocate_list,omitempty"`
}

func (x *BonusRequest) Reset() {
	*x = BonusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bonus_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BonusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BonusRequest) ProtoMessage() {}

func (x *BonusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bonus_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BonusRequest.ProtoReflect.Descriptor instead.
func (*BonusRequest) Descriptor() ([]byte, []int) {
	return file_bonus_service_proto_rawDescGZIP(), []int{0}
}

func (x *BonusRequest) GetAccountList() []string {
	if x != nil {
		return x.AccountList
	}
	return nil
}

func (x *BonusRequest) GetMoneyAllocateList() []int32 {
	if x != nil {
		return x.MoneyAllocateList
	}
	return nil
}

type BonusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *BonusRspHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *BonusResponse) Reset() {
	*x = BonusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bonus_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BonusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BonusResponse) ProtoMessage() {}

func (x *BonusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bonus_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BonusResponse.ProtoReflect.Descriptor instead.
func (*BonusResponse) Descriptor() ([]byte, []int) {
	return file_bonus_service_proto_rawDescGZIP(), []int{1}
}

func (x *BonusResponse) GetHeader() *BonusRspHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

type BonusRspHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret    int32  `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *BonusRspHeader) Reset() {
	*x = BonusRspHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bonus_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BonusRspHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BonusRspHeader) ProtoMessage() {}

func (x *BonusRspHeader) ProtoReflect() protoreflect.Message {
	mi := &file_bonus_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BonusRspHeader.ProtoReflect.Descriptor instead.
func (*BonusRspHeader) Descriptor() ([]byte, []int) {
	return file_bonus_service_proto_rawDescGZIP(), []int{2}
}

func (x *BonusRspHeader) GetRet() int32 {
	if x != nil {
		return x.Ret
	}
	return 0
}

func (x *BonusRspHeader) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

var File_bonus_service_proto protoreflect.FileDescriptor

var file_bonus_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x0c, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x6f, 0x6e, 0x65,
	0x79, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x11, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x38, 0x0a, 0x0d, 0x42, 0x6f, 0x6e, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x42, 0x6f, 0x6e, 0x75,
	0x73, 0x52, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x22, 0x3a, 0x0a, 0x0e, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x52, 0x73, 0x70, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x72, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2a, 0x75,
	0x0a, 0x16, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x42, 0x4f, 0x4e, 0x55, 0x53,
	0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x4e, 0x4f, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x23, 0x0a, 0x16, 0x43,
	0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f,
	0x50, 0x41, 0x52, 0x41, 0x4d, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01,
	0x12, 0x23, 0x0a, 0x16, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0xfe, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0x01, 0x32, 0x39, 0x0a, 0x0c, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6e, 0x75,
	0x73, 0x12, 0x0d, 0x2e, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0e, 0x2e, 0x42, 0x6f, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x6f, 0x79, 0x2d, 0x77, 0x65, 0x62, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62,
	0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bonus_service_proto_rawDescOnce sync.Once
	file_bonus_service_proto_rawDescData = file_bonus_service_proto_rawDesc
)

func file_bonus_service_proto_rawDescGZIP() []byte {
	file_bonus_service_proto_rawDescOnce.Do(func() {
		file_bonus_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_bonus_service_proto_rawDescData)
	})
	return file_bonus_service_proto_rawDescData
}

var file_bonus_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_bonus_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bonus_service_proto_goTypes = []interface{}{
	(ERR_CODE_BONUS_SERVICE)(0), // 0: ERR_CODE_BONUS_SERVICE
	(*BonusRequest)(nil),        // 1: BonusRequest
	(*BonusResponse)(nil),       // 2: BonusResponse
	(*BonusRspHeader)(nil),      // 3: BonusRspHeader
}
var file_bonus_service_proto_depIdxs = []int32{
	3, // 0: BonusResponse.header:type_name -> BonusRspHeader
	1, // 1: BonusService.GetBonus:input_type -> BonusRequest
	2, // 2: BonusService.GetBonus:output_type -> BonusResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bonus_service_proto_init() }
func file_bonus_service_proto_init() {
	if File_bonus_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bonus_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BonusRequest); i {
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
		file_bonus_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BonusResponse); i {
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
		file_bonus_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BonusRspHeader); i {
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
			RawDescriptor: file_bonus_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bonus_service_proto_goTypes,
		DependencyIndexes: file_bonus_service_proto_depIdxs,
		EnumInfos:         file_bonus_service_proto_enumTypes,
		MessageInfos:      file_bonus_service_proto_msgTypes,
	}.Build()
	File_bonus_service_proto = out.File
	file_bonus_service_proto_rawDesc = nil
	file_bonus_service_proto_goTypes = nil
	file_bonus_service_proto_depIdxs = nil
}
