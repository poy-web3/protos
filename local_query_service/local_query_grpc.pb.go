// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: local_query.proto

package local_query_service

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
	LocalQueryService_UserValidityCheck_FullMethodName = "/LocalQueryService/UserValidityCheck"
)

// LocalQueryServiceClient is the client API for LocalQueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocalQueryServiceClient interface {
	UserValidityCheck(ctx context.Context, in *UserValidityCheckRequest, opts ...grpc.CallOption) (*UserValidityResponse, error)
}

type localQueryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocalQueryServiceClient(cc grpc.ClientConnInterface) LocalQueryServiceClient {
	return &localQueryServiceClient{cc}
}

func (c *localQueryServiceClient) UserValidityCheck(ctx context.Context, in *UserValidityCheckRequest, opts ...grpc.CallOption) (*UserValidityResponse, error) {
	out := new(UserValidityResponse)
	err := c.cc.Invoke(ctx, LocalQueryService_UserValidityCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocalQueryServiceServer is the server API for LocalQueryService service.
// All implementations must embed UnimplementedLocalQueryServiceServer
// for forward compatibility
type LocalQueryServiceServer interface {
	UserValidityCheck(context.Context, *UserValidityCheckRequest) (*UserValidityResponse, error)
	mustEmbedUnimplementedLocalQueryServiceServer()
}

// UnimplementedLocalQueryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLocalQueryServiceServer struct {
}

func (UnimplementedLocalQueryServiceServer) UserValidityCheck(context.Context, *UserValidityCheckRequest) (*UserValidityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserValidityCheck not implemented")
}
func (UnimplementedLocalQueryServiceServer) mustEmbedUnimplementedLocalQueryServiceServer() {}

// UnsafeLocalQueryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocalQueryServiceServer will
// result in compilation errors.
type UnsafeLocalQueryServiceServer interface {
	mustEmbedUnimplementedLocalQueryServiceServer()
}

func RegisterLocalQueryServiceServer(s grpc.ServiceRegistrar, srv LocalQueryServiceServer) {
	s.RegisterService(&LocalQueryService_ServiceDesc, srv)
}

func _LocalQueryService_UserValidityCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserValidityCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalQueryServiceServer).UserValidityCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocalQueryService_UserValidityCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalQueryServiceServer).UserValidityCheck(ctx, req.(*UserValidityCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LocalQueryService_ServiceDesc is the grpc.ServiceDesc for LocalQueryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocalQueryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "LocalQueryService",
	HandlerType: (*LocalQueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserValidityCheck",
			Handler:    _LocalQueryService_UserValidityCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "local_query.proto",
}
