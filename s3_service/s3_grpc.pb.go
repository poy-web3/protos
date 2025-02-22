// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.26.1
// source: s3.proto

package s3_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	S3_UploadConstImage_FullMethodName   = "/S3/UploadConstImage"
	S3_UploadIndexImageV2_FullMethodName = "/S3/UploadIndexImageV2"
	S3_UpdateIndexImageV2_FullMethodName = "/S3/UpdateIndexImageV2"
	S3_GetUserFacesByYID_FullMethodName  = "/S3/GetUserFacesByYID"
	S3_GetConstImageByYID_FullMethodName = "/S3/GetConstImageByYID"
)

// S3Client is the client API for S3 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type S3Client interface {
	UploadConstImage(ctx context.Context, in *UploadConstImageRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	UploadIndexImageV2(ctx context.Context, in *UploadIndexImageRequestV2, opts ...grpc.CallOption) (*CommonResponse, error)
	UpdateIndexImageV2(ctx context.Context, in *UpdateIndexImageRequestV2, opts ...grpc.CallOption) (*CommonResponse, error)
	GetUserFacesByYID(ctx context.Context, in *GetUserFacesRequest, opts ...grpc.CallOption) (*GetUserFacesResponseV2, error)
	GetConstImageByYID(ctx context.Context, in *GetConstImageRequest, opts ...grpc.CallOption) (*GetConstImageResponse, error)
}

type s3Client struct {
	cc grpc.ClientConnInterface
}

func NewS3Client(cc grpc.ClientConnInterface) S3Client {
	return &s3Client{cc}
}

func (c *s3Client) UploadConstImage(ctx context.Context, in *UploadConstImageRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, S3_UploadConstImage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Client) UploadIndexImageV2(ctx context.Context, in *UploadIndexImageRequestV2, opts ...grpc.CallOption) (*CommonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, S3_UploadIndexImageV2_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Client) UpdateIndexImageV2(ctx context.Context, in *UpdateIndexImageRequestV2, opts ...grpc.CallOption) (*CommonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, S3_UpdateIndexImageV2_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Client) GetUserFacesByYID(ctx context.Context, in *GetUserFacesRequest, opts ...grpc.CallOption) (*GetUserFacesResponseV2, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserFacesResponseV2)
	err := c.cc.Invoke(ctx, S3_GetUserFacesByYID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Client) GetConstImageByYID(ctx context.Context, in *GetConstImageRequest, opts ...grpc.CallOption) (*GetConstImageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetConstImageResponse)
	err := c.cc.Invoke(ctx, S3_GetConstImageByYID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// S3Server is the server API for S3 service.
// All implementations must embed UnimplementedS3Server
// for forward compatibility
type S3Server interface {
	UploadConstImage(context.Context, *UploadConstImageRequest) (*CommonResponse, error)
	UploadIndexImageV2(context.Context, *UploadIndexImageRequestV2) (*CommonResponse, error)
	UpdateIndexImageV2(context.Context, *UpdateIndexImageRequestV2) (*CommonResponse, error)
	GetUserFacesByYID(context.Context, *GetUserFacesRequest) (*GetUserFacesResponseV2, error)
	GetConstImageByYID(context.Context, *GetConstImageRequest) (*GetConstImageResponse, error)
	mustEmbedUnimplementedS3Server()
}

// UnimplementedS3Server must be embedded to have forward compatible implementations.
type UnimplementedS3Server struct {
}

func (UnimplementedS3Server) UploadConstImage(context.Context, *UploadConstImageRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadConstImage not implemented")
}
func (UnimplementedS3Server) UploadIndexImageV2(context.Context, *UploadIndexImageRequestV2) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadIndexImageV2 not implemented")
}
func (UnimplementedS3Server) UpdateIndexImageV2(context.Context, *UpdateIndexImageRequestV2) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIndexImageV2 not implemented")
}
func (UnimplementedS3Server) GetUserFacesByYID(context.Context, *GetUserFacesRequest) (*GetUserFacesResponseV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFacesByYID not implemented")
}
func (UnimplementedS3Server) GetConstImageByYID(context.Context, *GetConstImageRequest) (*GetConstImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConstImageByYID not implemented")
}
func (UnimplementedS3Server) mustEmbedUnimplementedS3Server() {}

// UnsafeS3Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to S3Server will
// result in compilation errors.
type UnsafeS3Server interface {
	mustEmbedUnimplementedS3Server()
}

func RegisterS3Server(s grpc.ServiceRegistrar, srv S3Server) {
	s.RegisterService(&S3_ServiceDesc, srv)
}

func _S3_UploadConstImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadConstImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3Server).UploadConstImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3_UploadConstImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3Server).UploadConstImage(ctx, req.(*UploadConstImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3_UploadIndexImageV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadIndexImageRequestV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3Server).UploadIndexImageV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3_UploadIndexImageV2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3Server).UploadIndexImageV2(ctx, req.(*UploadIndexImageRequestV2))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3_UpdateIndexImageV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIndexImageRequestV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3Server).UpdateIndexImageV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3_UpdateIndexImageV2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3Server).UpdateIndexImageV2(ctx, req.(*UpdateIndexImageRequestV2))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3_GetUserFacesByYID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserFacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3Server).GetUserFacesByYID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3_GetUserFacesByYID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3Server).GetUserFacesByYID(ctx, req.(*GetUserFacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3_GetConstImageByYID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConstImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3Server).GetConstImageByYID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3_GetConstImageByYID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3Server).GetConstImageByYID(ctx, req.(*GetConstImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// S3_ServiceDesc is the grpc.ServiceDesc for S3 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var S3_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "S3",
	HandlerType: (*S3Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadConstImage",
			Handler:    _S3_UploadConstImage_Handler,
		},
		{
			MethodName: "UploadIndexImageV2",
			Handler:    _S3_UploadIndexImageV2_Handler,
		},
		{
			MethodName: "UpdateIndexImageV2",
			Handler:    _S3_UpdateIndexImageV2_Handler,
		},
		{
			MethodName: "GetUserFacesByYID",
			Handler:    _S3_GetUserFacesByYID_Handler,
		},
		{
			MethodName: "GetConstImageByYID",
			Handler:    _S3_GetConstImageByYID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "s3.proto",
}
