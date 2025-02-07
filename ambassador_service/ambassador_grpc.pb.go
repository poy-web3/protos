// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: ambassador.proto

package ambassador_service

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
	Ambassador_AddConnection_FullMethodName = "/Ambassador/AddConnection"
)

// AmbassadorClient is the client API for Ambassador service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AmbassadorClient interface {
	AddConnection(ctx context.Context, in *AmbassadorRequest, opts ...grpc.CallOption) (*AmbassadorResponse, error)
}

type ambassadorClient struct {
	cc grpc.ClientConnInterface
}

func NewAmbassadorClient(cc grpc.ClientConnInterface) AmbassadorClient {
	return &ambassadorClient{cc}
}

func (c *ambassadorClient) AddConnection(ctx context.Context, in *AmbassadorRequest, opts ...grpc.CallOption) (*AmbassadorResponse, error) {
	out := new(AmbassadorResponse)
	err := c.cc.Invoke(ctx, Ambassador_AddConnection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AmbassadorServer is the server API for Ambassador service.
// All implementations must embed UnimplementedAmbassadorServer
// for forward compatibility
type AmbassadorServer interface {
	AddConnection(context.Context, *AmbassadorRequest) (*AmbassadorResponse, error)
	mustEmbedUnimplementedAmbassadorServer()
}

// UnimplementedAmbassadorServer must be embedded to have forward compatible implementations.
type UnimplementedAmbassadorServer struct {
}

func (UnimplementedAmbassadorServer) AddConnection(context.Context, *AmbassadorRequest) (*AmbassadorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddConnection not implemented")
}
func (UnimplementedAmbassadorServer) mustEmbedUnimplementedAmbassadorServer() {}

// UnsafeAmbassadorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AmbassadorServer will
// result in compilation errors.
type UnsafeAmbassadorServer interface {
	mustEmbedUnimplementedAmbassadorServer()
}

func RegisterAmbassadorServer(s grpc.ServiceRegistrar, srv AmbassadorServer) {
	s.RegisterService(&Ambassador_ServiceDesc, srv)
}

func _Ambassador_AddConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AmbassadorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmbassadorServer).AddConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ambassador_AddConnection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmbassadorServer).AddConnection(ctx, req.(*AmbassadorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Ambassador_ServiceDesc is the grpc.ServiceDesc for Ambassador service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ambassador_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Ambassador",
	HandlerType: (*AmbassadorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddConnection",
			Handler:    _Ambassador_AddConnection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ambassador.proto",
}
