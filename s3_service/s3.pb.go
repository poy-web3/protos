// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: s3.proto

package s3_service

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

type UploadConstImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConstImage []byte `protobuf:"bytes,1,opt,name=const_image,json=constImage,proto3" json:"const_image,omitempty"`
	Yid        string `protobuf:"bytes,2,opt,name=yid,proto3" json:"yid,omitempty"`
}

func (x *UploadConstImageRequest) Reset() {
	*x = UploadConstImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadConstImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadConstImageRequest) ProtoMessage() {}

func (x *UploadConstImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_s3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadConstImageRequest.ProtoReflect.Descriptor instead.
func (*UploadConstImageRequest) Descriptor() ([]byte, []int) {
	return file_s3_proto_rawDescGZIP(), []int{0}
}

func (x *UploadConstImageRequest) GetConstImage() []byte {
	if x != nil {
		return x.ConstImage
	}
	return nil
}

func (x *UploadConstImageRequest) GetYid() string {
	if x != nil {
		return x.Yid
	}
	return ""
}

type UploadConstImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Succ         bool   `protobuf:"varint,1,opt,name=succ,proto3" json:"succ,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *UploadConstImageResponse) Reset() {
	*x = UploadConstImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s3_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadConstImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadConstImageResponse) ProtoMessage() {}

func (x *UploadConstImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_s3_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadConstImageResponse.ProtoReflect.Descriptor instead.
func (*UploadConstImageResponse) Descriptor() ([]byte, []int) {
	return file_s3_proto_rawDescGZIP(), []int{1}
}

func (x *UploadConstImageResponse) GetSucc() bool {
	if x != nil {
		return x.Succ
	}
	return false
}

func (x *UploadConstImageResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type UploadIndexImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FaceIdToImage map[string][]byte `protobuf:"bytes,1,rep,name=face_id_to_image,json=faceIdToImage,proto3" json:"face_id_to_image,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Yid           string            `protobuf:"bytes,2,opt,name=yid,proto3" json:"yid,omitempty"`
}

func (x *UploadIndexImageRequest) Reset() {
	*x = UploadIndexImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s3_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadIndexImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadIndexImageRequest) ProtoMessage() {}

func (x *UploadIndexImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_s3_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadIndexImageRequest.ProtoReflect.Descriptor instead.
func (*UploadIndexImageRequest) Descriptor() ([]byte, []int) {
	return file_s3_proto_rawDescGZIP(), []int{2}
}

func (x *UploadIndexImageRequest) GetFaceIdToImage() map[string][]byte {
	if x != nil {
		return x.FaceIdToImage
	}
	return nil
}

func (x *UploadIndexImageRequest) GetYid() string {
	if x != nil {
		return x.Yid
	}
	return ""
}

type UpdateIndexImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FaceId string `protobuf:"bytes,1,opt,name=face_id,json=faceId,proto3" json:"face_id,omitempty"`
	Image  []byte `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Yid    string `protobuf:"bytes,3,opt,name=yid,proto3" json:"yid,omitempty"`
}

func (x *UpdateIndexImageRequest) Reset() {
	*x = UpdateIndexImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s3_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateIndexImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateIndexImageRequest) ProtoMessage() {}

func (x *UpdateIndexImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_s3_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateIndexImageRequest.ProtoReflect.Descriptor instead.
func (*UpdateIndexImageRequest) Descriptor() ([]byte, []int) {
	return file_s3_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateIndexImageRequest) GetFaceId() string {
	if x != nil {
		return x.FaceId
	}
	return ""
}

func (x *UpdateIndexImageRequest) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *UpdateIndexImageRequest) GetYid() string {
	if x != nil {
		return x.Yid
	}
	return ""
}

type UploadIndexImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Succ         bool   `protobuf:"varint,1,opt,name=succ,proto3" json:"succ,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *UploadIndexImageResponse) Reset() {
	*x = UploadIndexImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s3_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadIndexImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadIndexImageResponse) ProtoMessage() {}

func (x *UploadIndexImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_s3_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadIndexImageResponse.ProtoReflect.Descriptor instead.
func (*UploadIndexImageResponse) Descriptor() ([]byte, []int) {
	return file_s3_proto_rawDescGZIP(), []int{4}
}

func (x *UploadIndexImageResponse) GetSucc() bool {
	if x != nil {
		return x.Succ
	}
	return false
}

func (x *UploadIndexImageResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type UpdateIndexImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Succ          bool   `protobuf:"varint,1,opt,name=succ,proto3" json:"succ,omitempty"`
	DeletedFaceId string `protobuf:"bytes,2,opt,name=deleted_face_id,json=deletedFaceId,proto3" json:"deleted_face_id,omitempty"`
	ErrorMessage  string `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *UpdateIndexImageResponse) Reset() {
	*x = UpdateIndexImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s3_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateIndexImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateIndexImageResponse) ProtoMessage() {}

func (x *UpdateIndexImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_s3_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateIndexImageResponse.ProtoReflect.Descriptor instead.
func (*UpdateIndexImageResponse) Descriptor() ([]byte, []int) {
	return file_s3_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateIndexImageResponse) GetSucc() bool {
	if x != nil {
		return x.Succ
	}
	return false
}

func (x *UpdateIndexImageResponse) GetDeletedFaceId() string {
	if x != nil {
		return x.DeletedFaceId
	}
	return ""
}

func (x *UpdateIndexImageResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type GetUserFacesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Yid string `protobuf:"bytes,1,opt,name=yid,proto3" json:"yid,omitempty"`
}

func (x *GetUserFacesRequest) Reset() {
	*x = GetUserFacesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s3_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserFacesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserFacesRequest) ProtoMessage() {}

func (x *GetUserFacesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_s3_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserFacesRequest.ProtoReflect.Descriptor instead.
func (*GetUserFacesRequest) Descriptor() ([]byte, []int) {
	return file_s3_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserFacesRequest) GetYid() string {
	if x != nil {
		return x.Yid
	}
	return ""
}

type GetUserFacesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FaceIds      []string `protobuf:"bytes,1,rep,name=face_ids,json=faceIds,proto3" json:"face_ids,omitempty"`
	ErrorMessage string   `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *GetUserFacesResponse) Reset() {
	*x = GetUserFacesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s3_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserFacesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserFacesResponse) ProtoMessage() {}

func (x *GetUserFacesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_s3_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserFacesResponse.ProtoReflect.Descriptor instead.
func (*GetUserFacesResponse) Descriptor() ([]byte, []int) {
	return file_s3_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserFacesResponse) GetFaceIds() []string {
	if x != nil {
		return x.FaceIds
	}
	return nil
}

func (x *GetUserFacesResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_s3_proto protoreflect.FileDescriptor

var file_s3_proto_rawDesc = []byte{
	0x0a, 0x08, 0x73, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x17, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x5f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x73,
	0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x79, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x79, 0x69, 0x64, 0x22, 0x53, 0x0a, 0x18, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x75, 0x63, 0x63, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x73, 0x75, 0x63, 0x63, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xc3, 0x01,
	0x0a, 0x17, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x54, 0x0a, 0x10, 0x66, 0x61, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x61,
	0x63, 0x65, 0x49, 0x64, 0x54, 0x6f, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0d, 0x66, 0x61, 0x63, 0x65, 0x49, 0x64, 0x54, 0x6f, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x79, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x79, 0x69,
	0x64, 0x1a, 0x40, 0x0a, 0x12, 0x46, 0x61, 0x63, 0x65, 0x49, 0x64, 0x54, 0x6f, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x5a, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x79, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x79, 0x69, 0x64, 0x22,
	0x53, 0x0a, 0x18, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x75, 0x63, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x75, 0x63, 0x63, 0x12,
	0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x7b, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x75, 0x63, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x73, 0x75, 0x63, 0x63, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f,
	0x66, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x46, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x27, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x79, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x79, 0x69, 0x64, 0x22, 0x56, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x63, 0x65, 0x49, 0x64, 0x73, 0x12, 0x23, 0x0a,
	0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0xa1, 0x02, 0x0a, 0x02, 0x53, 0x33, 0x12, 0x47, 0x0a, 0x10, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x18, 0x2e,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x43, 0x6f, 0x6e, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x47, 0x0a, 0x10, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x18, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x10, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x46,
	0x61, 0x63, 0x65, 0x73, 0x42, 0x79, 0x59, 0x49, 0x44, 0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x46, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6f, 0x79, 0x2d, 0x77, 0x65, 0x62, 0x33, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x33, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_s3_proto_rawDescOnce sync.Once
	file_s3_proto_rawDescData = file_s3_proto_rawDesc
)

func file_s3_proto_rawDescGZIP() []byte {
	file_s3_proto_rawDescOnce.Do(func() {
		file_s3_proto_rawDescData = protoimpl.X.CompressGZIP(file_s3_proto_rawDescData)
	})
	return file_s3_proto_rawDescData
}

var file_s3_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_s3_proto_goTypes = []interface{}{
	(*UploadConstImageRequest)(nil),  // 0: UploadConstImageRequest
	(*UploadConstImageResponse)(nil), // 1: UploadConstImageResponse
	(*UploadIndexImageRequest)(nil),  // 2: UploadIndexImageRequest
	(*UpdateIndexImageRequest)(nil),  // 3: UpdateIndexImageRequest
	(*UploadIndexImageResponse)(nil), // 4: UploadIndexImageResponse
	(*UpdateIndexImageResponse)(nil), // 5: UpdateIndexImageResponse
	(*GetUserFacesRequest)(nil),      // 6: GetUserFacesRequest
	(*GetUserFacesResponse)(nil),     // 7: GetUserFacesResponse
	nil,                              // 8: UploadIndexImageRequest.FaceIdToImageEntry
}
var file_s3_proto_depIdxs = []int32{
	8, // 0: UploadIndexImageRequest.face_id_to_image:type_name -> UploadIndexImageRequest.FaceIdToImageEntry
	0, // 1: S3.UploadConstImage:input_type -> UploadConstImageRequest
	2, // 2: S3.UploadIndexImage:input_type -> UploadIndexImageRequest
	3, // 3: S3.UpdateIndexImage:input_type -> UpdateIndexImageRequest
	6, // 4: S3.GetUserFacesByYID:input_type -> GetUserFacesRequest
	1, // 5: S3.UploadConstImage:output_type -> UploadConstImageResponse
	4, // 6: S3.UploadIndexImage:output_type -> UploadIndexImageResponse
	5, // 7: S3.UpdateIndexImage:output_type -> UpdateIndexImageResponse
	7, // 8: S3.GetUserFacesByYID:output_type -> GetUserFacesResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_s3_proto_init() }
func file_s3_proto_init() {
	if File_s3_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_s3_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadConstImageRequest); i {
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
		file_s3_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadConstImageResponse); i {
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
		file_s3_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadIndexImageRequest); i {
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
		file_s3_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateIndexImageRequest); i {
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
		file_s3_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadIndexImageResponse); i {
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
		file_s3_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateIndexImageResponse); i {
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
		file_s3_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserFacesRequest); i {
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
		file_s3_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserFacesResponse); i {
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
			RawDescriptor: file_s3_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_s3_proto_goTypes,
		DependencyIndexes: file_s3_proto_depIdxs,
		MessageInfos:      file_s3_proto_msgTypes,
	}.Build()
	File_s3_proto = out.File
	file_s3_proto_rawDesc = nil
	file_s3_proto_goTypes = nil
	file_s3_proto_depIdxs = nil
}
