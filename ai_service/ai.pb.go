// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: ai.proto

package ai_service

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

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FacialImages [][]byte `protobuf:"bytes,1,rep,name=facial_images,json=facialImages,proto3" json:"facial_images,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ai_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_ai_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetFacialImages() [][]byte {
	if x != nil {
		return x.FacialImages
	}
	return nil
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header  *RspHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	FaceIds []string   `protobuf:"bytes,2,rep,name=face_ids,json=faceIds,proto3" json:"face_ids,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ai_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_ai_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetHeader() *RspHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *RegisterResponse) GetFaceIds() []string {
	if x != nil {
		return x.FaceIds
	}
	return nil
}

type RecoverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FacialImages [][]byte `protobuf:"bytes,1,rep,name=facial_images,json=facialImages,proto3" json:"facial_images,omitempty"`
}

func (x *RecoverRequest) Reset() {
	*x = RecoverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecoverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecoverRequest) ProtoMessage() {}

func (x *RecoverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ai_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecoverRequest.ProtoReflect.Descriptor instead.
func (*RecoverRequest) Descriptor() ([]byte, []int) {
	return file_ai_proto_rawDescGZIP(), []int{2}
}

func (x *RecoverRequest) GetFacialImages() [][]byte {
	if x != nil {
		return x.FacialImages
	}
	return nil
}

type RecoverResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header  *RspHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	FaceIds []string   `protobuf:"bytes,2,rep,name=face_ids,json=faceIds,proto3" json:"face_ids,omitempty"` // matched face ids
}

func (x *RecoverResponse) Reset() {
	*x = RecoverResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecoverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecoverResponse) ProtoMessage() {}

func (x *RecoverResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ai_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecoverResponse.ProtoReflect.Descriptor instead.
func (*RecoverResponse) Descriptor() ([]byte, []int) {
	return file_ai_proto_rawDescGZIP(), []int{3}
}

func (x *RecoverResponse) GetHeader() *RspHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *RecoverResponse) GetFaceIds() []string {
	if x != nil {
		return x.FaceIds
	}
	return nil
}

type RefreshTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FacialImage []byte   `protobuf:"bytes,1,opt,name=facial_image,json=facialImage,proto3" json:"facial_image,omitempty"`
	FaceS3Keys  []string `protobuf:"bytes,2,rep,name=face_s3_keys,json=faceS3Keys,proto3" json:"face_s3_keys,omitempty"`
}

func (x *RefreshTokenRequest) Reset() {
	*x = RefreshTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenRequest) ProtoMessage() {}

func (x *RefreshTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ai_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenRequest.ProtoReflect.Descriptor instead.
func (*RefreshTokenRequest) Descriptor() ([]byte, []int) {
	return file_ai_proto_rawDescGZIP(), []int{4}
}

func (x *RefreshTokenRequest) GetFacialImage() []byte {
	if x != nil {
		return x.FacialImage
	}
	return nil
}

func (x *RefreshTokenRequest) GetFaceS3Keys() []string {
	if x != nil {
		return x.FaceS3Keys
	}
	return nil
}

type RefreshTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header       *RspHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	NewFaceId    string     `protobuf:"bytes,2,opt,name=new_face_id,json=newFaceId,proto3" json:"new_face_id,omitempty"`
	DeleteFaceId string     `protobuf:"bytes,3,opt,name=delete_face_id,json=deleteFaceId,proto3" json:"delete_face_id,omitempty"`
}

func (x *RefreshTokenResponse) Reset() {
	*x = RefreshTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenResponse) ProtoMessage() {}

func (x *RefreshTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ai_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenResponse.ProtoReflect.Descriptor instead.
func (*RefreshTokenResponse) Descriptor() ([]byte, []int) {
	return file_ai_proto_rawDescGZIP(), []int{5}
}

func (x *RefreshTokenResponse) GetHeader() *RspHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *RefreshTokenResponse) GetNewFaceId() string {
	if x != nil {
		return x.NewFaceId
	}
	return ""
}

func (x *RefreshTokenResponse) GetDeleteFaceId() string {
	if x != nil {
		return x.DeleteFaceId
	}
	return ""
}

type RspHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Succ         bool   `protobuf:"varint,1,opt,name=succ,proto3" json:"succ,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *RspHeader) Reset() {
	*x = RspHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RspHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RspHeader) ProtoMessage() {}

func (x *RspHeader) ProtoReflect() protoreflect.Message {
	mi := &file_ai_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RspHeader.ProtoReflect.Descriptor instead.
func (*RspHeader) Descriptor() ([]byte, []int) {
	return file_ai_proto_rawDescGZIP(), []int{6}
}

func (x *RspHeader) GetSucc() bool {
	if x != nil {
		return x.Succ
	}
	return false
}

func (x *RspHeader) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type RegisterRollbackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FaceIds []string `protobuf:"bytes,2,rep,name=face_ids,json=faceIds,proto3" json:"face_ids,omitempty"`
}

func (x *RegisterRollbackRequest) Reset() {
	*x = RegisterRollbackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRollbackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRollbackRequest) ProtoMessage() {}

func (x *RegisterRollbackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ai_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRollbackRequest.ProtoReflect.Descriptor instead.
func (*RegisterRollbackRequest) Descriptor() ([]byte, []int) {
	return file_ai_proto_rawDescGZIP(), []int{7}
}

func (x *RegisterRollbackRequest) GetFaceIds() []string {
	if x != nil {
		return x.FaceIds
	}
	return nil
}

type RegisterRollbackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *RspHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *RegisterRollbackResponse) Reset() {
	*x = RegisterRollbackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRollbackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRollbackResponse) ProtoMessage() {}

func (x *RegisterRollbackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ai_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRollbackResponse.ProtoReflect.Descriptor instead.
func (*RegisterRollbackResponse) Descriptor() ([]byte, []int) {
	return file_ai_proto_rawDescGZIP(), []int{8}
}

func (x *RegisterRollbackResponse) GetHeader() *RspHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

var File_ai_proto protoreflect.FileDescriptor

var file_ai_proto_rawDesc = []byte{
	0x0a, 0x08, 0x61, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x0f, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a,
	0x0d, 0x66, 0x61, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0c, 0x52, 0x0c, 0x66, 0x61, 0x63, 0x69, 0x61, 0x6c, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x22, 0x51, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x52, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x61,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61,
	0x63, 0x65, 0x49, 0x64, 0x73, 0x22, 0x35, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x61, 0x63, 0x69, 0x61,
	0x6c, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0c,
	0x66, 0x61, 0x63, 0x69, 0x61, 0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x50, 0x0a, 0x0f,
	0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x52, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x63, 0x65, 0x49, 0x64, 0x73, 0x22, 0x5a,
	0x0a, 0x13, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x61, 0x63, 0x69, 0x61, 0x6c, 0x5f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x66, 0x61, 0x63,
	0x69, 0x61, 0x6c, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x66, 0x61, 0x63, 0x65,
	0x5f, 0x73, 0x33, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x66, 0x61, 0x63, 0x65, 0x53, 0x33, 0x4b, 0x65, 0x79, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x14, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x52, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x5f, 0x66,
	0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65,
	0x77, 0x46, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x5f, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x44, 0x0a,
	0x09, 0x52, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x75,
	0x63, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x75, 0x63, 0x63, 0x12, 0x23,
	0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x34, 0x0a, 0x17, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x66, 0x61, 0x63, 0x65, 0x49, 0x64, 0x73, 0x22, 0x3e, 0x0a, 0x18, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x52, 0x73, 0x70, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x32, 0xf7, 0x01, 0x0a, 0x02, 0x41, 0x49,
	0x12, 0x36, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x10, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x52, 0x65, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x52, 0x65,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a,
	0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x2e,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x10, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x12, 0x18,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x6f, 0x79, 0x2d, 0x77, 0x65, 0x62, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x61, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ai_proto_rawDescOnce sync.Once
	file_ai_proto_rawDescData = file_ai_proto_rawDesc
)

func file_ai_proto_rawDescGZIP() []byte {
	file_ai_proto_rawDescOnce.Do(func() {
		file_ai_proto_rawDescData = protoimpl.X.CompressGZIP(file_ai_proto_rawDescData)
	})
	return file_ai_proto_rawDescData
}

var file_ai_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_ai_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),          // 0: RegisterRequest
	(*RegisterResponse)(nil),         // 1: RegisterResponse
	(*RecoverRequest)(nil),           // 2: RecoverRequest
	(*RecoverResponse)(nil),          // 3: RecoverResponse
	(*RefreshTokenRequest)(nil),      // 4: RefreshTokenRequest
	(*RefreshTokenResponse)(nil),     // 5: RefreshTokenResponse
	(*RspHeader)(nil),                // 6: RspHeader
	(*RegisterRollbackRequest)(nil),  // 7: RegisterRollbackRequest
	(*RegisterRollbackResponse)(nil), // 8: RegisterRollbackResponse
}
var file_ai_proto_depIdxs = []int32{
	6, // 0: RegisterResponse.header:type_name -> RspHeader
	6, // 1: RecoverResponse.header:type_name -> RspHeader
	6, // 2: RefreshTokenResponse.header:type_name -> RspHeader
	6, // 3: RegisterRollbackResponse.header:type_name -> RspHeader
	0, // 4: AI.RegisterAccount:input_type -> RegisterRequest
	2, // 5: AI.RecoverAccount:input_type -> RecoverRequest
	4, // 6: AI.RefreshToken:input_type -> RefreshTokenRequest
	7, // 7: AI.RegisterRollBack:input_type -> RegisterRollbackRequest
	1, // 8: AI.RegisterAccount:output_type -> RegisterResponse
	3, // 9: AI.RecoverAccount:output_type -> RecoverResponse
	5, // 10: AI.RefreshToken:output_type -> RefreshTokenResponse
	8, // 11: AI.RegisterRollBack:output_type -> RegisterRollbackResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_ai_proto_init() }
func file_ai_proto_init() {
	if File_ai_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ai_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_ai_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_ai_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecoverRequest); i {
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
		file_ai_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecoverResponse); i {
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
		file_ai_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshTokenRequest); i {
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
		file_ai_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshTokenResponse); i {
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
		file_ai_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RspHeader); i {
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
		file_ai_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRollbackRequest); i {
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
		file_ai_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRollbackResponse); i {
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
			RawDescriptor: file_ai_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ai_proto_goTypes,
		DependencyIndexes: file_ai_proto_depIdxs,
		MessageInfos:      file_ai_proto_msgTypes,
	}.Build()
	File_ai_proto = out.File
	file_ai_proto_rawDesc = nil
	file_ai_proto_goTypes = nil
	file_ai_proto_depIdxs = nil
}
