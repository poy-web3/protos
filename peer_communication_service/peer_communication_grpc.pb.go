// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: peer_communication.proto

package peer_communication_service

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

// PeerCommunicationServiceClient is the client API for PeerCommunicationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeerCommunicationServiceClient interface {
	BlockHeight(ctx context.Context, in *BlockHeightRequest, opts ...grpc.CallOption) (*BlockHeightReponse, error)
	GetRecentHeaders(ctx context.Context, in *GetRecentHeadersRequest, opts ...grpc.CallOption) (*GetRecentHeadersResponse, error)
	ValidatorSetByHeight(ctx context.Context, in *ValidatorSetByHeightRequest, opts ...grpc.CallOption) (*ValidatorSetByHeightResponse, error)
	GetCoinOwner(ctx context.Context, in *GetCoinOwnerRequest, opts ...grpc.CallOption) (*GetCoinOwnerResponse, error)
}

type peerCommunicationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPeerCommunicationServiceClient(cc grpc.ClientConnInterface) PeerCommunicationServiceClient {
	return &peerCommunicationServiceClient{cc}
}

func (c *peerCommunicationServiceClient) BlockHeight(ctx context.Context, in *BlockHeightRequest, opts ...grpc.CallOption) (*BlockHeightReponse, error) {
	out := new(BlockHeightReponse)
	err := c.cc.Invoke(ctx, "/PeerCommunicationService/BlockHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerCommunicationServiceClient) GetRecentHeaders(ctx context.Context, in *GetRecentHeadersRequest, opts ...grpc.CallOption) (*GetRecentHeadersResponse, error) {
	out := new(GetRecentHeadersResponse)
	err := c.cc.Invoke(ctx, "/PeerCommunicationService/GetRecentHeaders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerCommunicationServiceClient) ValidatorSetByHeight(ctx context.Context, in *ValidatorSetByHeightRequest, opts ...grpc.CallOption) (*ValidatorSetByHeightResponse, error) {
	out := new(ValidatorSetByHeightResponse)
	err := c.cc.Invoke(ctx, "/PeerCommunicationService/ValidatorSetByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerCommunicationServiceClient) GetCoinOwner(ctx context.Context, in *GetCoinOwnerRequest, opts ...grpc.CallOption) (*GetCoinOwnerResponse, error) {
	out := new(GetCoinOwnerResponse)
	err := c.cc.Invoke(ctx, "/PeerCommunicationService/GetCoinOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeerCommunicationServiceServer is the server API for PeerCommunicationService service.
// All implementations must embed UnimplementedPeerCommunicationServiceServer
// for forward compatibility
type PeerCommunicationServiceServer interface {
	BlockHeight(context.Context, *BlockHeightRequest) (*BlockHeightReponse, error)
	GetRecentHeaders(context.Context, *GetRecentHeadersRequest) (*GetRecentHeadersResponse, error)
	ValidatorSetByHeight(context.Context, *ValidatorSetByHeightRequest) (*ValidatorSetByHeightResponse, error)
	GetCoinOwner(context.Context, *GetCoinOwnerRequest) (*GetCoinOwnerResponse, error)
	mustEmbedUnimplementedPeerCommunicationServiceServer()
}

// UnimplementedPeerCommunicationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPeerCommunicationServiceServer struct {
}

func (UnimplementedPeerCommunicationServiceServer) BlockHeight(context.Context, *BlockHeightRequest) (*BlockHeightReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockHeight not implemented")
}
func (UnimplementedPeerCommunicationServiceServer) GetRecentHeaders(context.Context, *GetRecentHeadersRequest) (*GetRecentHeadersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecentHeaders not implemented")
}
func (UnimplementedPeerCommunicationServiceServer) ValidatorSetByHeight(context.Context, *ValidatorSetByHeightRequest) (*ValidatorSetByHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorSetByHeight not implemented")
}
func (UnimplementedPeerCommunicationServiceServer) GetCoinOwner(context.Context, *GetCoinOwnerRequest) (*GetCoinOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoinOwner not implemented")
}
func (UnimplementedPeerCommunicationServiceServer) mustEmbedUnimplementedPeerCommunicationServiceServer() {
}

// UnsafePeerCommunicationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PeerCommunicationServiceServer will
// result in compilation errors.
type UnsafePeerCommunicationServiceServer interface {
	mustEmbedUnimplementedPeerCommunicationServiceServer()
}

func RegisterPeerCommunicationServiceServer(s grpc.ServiceRegistrar, srv PeerCommunicationServiceServer) {
	s.RegisterService(&PeerCommunicationService_ServiceDesc, srv)
}

func _PeerCommunicationService_BlockHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerCommunicationServiceServer).BlockHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PeerCommunicationService/BlockHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerCommunicationServiceServer).BlockHeight(ctx, req.(*BlockHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerCommunicationService_GetRecentHeaders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecentHeadersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerCommunicationServiceServer).GetRecentHeaders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PeerCommunicationService/GetRecentHeaders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerCommunicationServiceServer).GetRecentHeaders(ctx, req.(*GetRecentHeadersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerCommunicationService_ValidatorSetByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidatorSetByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerCommunicationServiceServer).ValidatorSetByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PeerCommunicationService/ValidatorSetByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerCommunicationServiceServer).ValidatorSetByHeight(ctx, req.(*ValidatorSetByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerCommunicationService_GetCoinOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoinOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerCommunicationServiceServer).GetCoinOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PeerCommunicationService/GetCoinOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerCommunicationServiceServer).GetCoinOwner(ctx, req.(*GetCoinOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PeerCommunicationService_ServiceDesc is the grpc.ServiceDesc for PeerCommunicationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PeerCommunicationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PeerCommunicationService",
	HandlerType: (*PeerCommunicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BlockHeight",
			Handler:    _PeerCommunicationService_BlockHeight_Handler,
		},
		{
			MethodName: "GetRecentHeaders",
			Handler:    _PeerCommunicationService_GetRecentHeaders_Handler,
		},
		{
			MethodName: "ValidatorSetByHeight",
			Handler:    _PeerCommunicationService_ValidatorSetByHeight_Handler,
		},
		{
			MethodName: "GetCoinOwner",
			Handler:    _PeerCommunicationService_GetCoinOwner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peer_communication.proto",
}
