// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	S3_UploadConstImage_FullMethodName = "/S3/UploadConstImage"
	S3_UploadIndexImage_FullMethodName = "/S3/UploadIndexImage"
	S3_UpdateIndexImage_FullMethodName = "/S3/UpdateIndexImage"
)

// S3Client is the client API for S3 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type S3Client interface {
	UploadConstImage(ctx context.Context, in *UploadConstImageRequest, opts ...grpc.CallOption) (*S3Response, error)
	UploadIndexImage(ctx context.Context, in *UploadIndexImageRequest, opts ...grpc.CallOption) (*S3Response, error)
	UpdateIndexImage(ctx context.Context, in *UpdateIndexImageRequest, opts ...grpc.CallOption) (*UpdateIndexImageResponse, error)
}

type s3Client struct {
	cc grpc.ClientConnInterface
}

func NewS3Client(cc grpc.ClientConnInterface) S3Client {
	return &s3Client{cc}
}

func (c *s3Client) UploadConstImage(ctx context.Context, in *UploadConstImageRequest, opts ...grpc.CallOption) (*S3Response, error) {
	out := new(S3Response)
	err := c.cc.Invoke(ctx, S3_UploadConstImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Client) UploadIndexImage(ctx context.Context, in *UploadIndexImageRequest, opts ...grpc.CallOption) (*S3Response, error) {
	out := new(S3Response)
	err := c.cc.Invoke(ctx, S3_UploadIndexImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3Client) UpdateIndexImage(ctx context.Context, in *UpdateIndexImageRequest, opts ...grpc.CallOption) (*UpdateIndexImageResponse, error) {
	out := new(UpdateIndexImageResponse)
	err := c.cc.Invoke(ctx, S3_UpdateIndexImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// S3Server is the server API for S3 service.
// All implementations must embed UnimplementedS3Server
// for forward compatibility
type S3Server interface {
	UploadConstImage(context.Context, *UploadConstImageRequest) (*S3Response, error)
	UploadIndexImage(context.Context, *UploadIndexImageRequest) (*S3Response, error)
	UpdateIndexImage(context.Context, *UpdateIndexImageRequest) (*UpdateIndexImageResponse, error)
	mustEmbedUnimplementedS3Server()
}

// UnimplementedS3Server must be embedded to have forward compatible implementations.
type UnimplementedS3Server struct {
}

func (UnimplementedS3Server) UploadConstImage(context.Context, *UploadConstImageRequest) (*S3Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadConstImage not implemented")
}
func (UnimplementedS3Server) UploadIndexImage(context.Context, *UploadIndexImageRequest) (*S3Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadIndexImage not implemented")
}
func (UnimplementedS3Server) UpdateIndexImage(context.Context, *UpdateIndexImageRequest) (*UpdateIndexImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIndexImage not implemented")
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

func _S3_UploadIndexImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadIndexImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3Server).UploadIndexImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3_UploadIndexImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3Server).UploadIndexImage(ctx, req.(*UploadIndexImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3_UpdateIndexImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIndexImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3Server).UpdateIndexImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: S3_UpdateIndexImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3Server).UpdateIndexImage(ctx, req.(*UpdateIndexImageRequest))
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
			MethodName: "UploadIndexImage",
			Handler:    _S3_UploadIndexImage_Handler,
		},
		{
			MethodName: "UpdateIndexImage",
			Handler:    _S3_UpdateIndexImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "s3.proto",
}
