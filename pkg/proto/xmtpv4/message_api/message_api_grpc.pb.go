// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: xmtpv4/message_api/message_api.proto

package message_api

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
	ReplicationApi_SubscribeEnvelopes_FullMethodName = "/xmtp.xmtpv4.ReplicationApi/SubscribeEnvelopes"
	ReplicationApi_QueryEnvelopes_FullMethodName     = "/xmtp.xmtpv4.ReplicationApi/QueryEnvelopes"
)

// ReplicationApiClient is the client API for ReplicationApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReplicationApiClient interface {
	SubscribeEnvelopes(ctx context.Context, in *BatchSubscribeEnvelopesRequest, opts ...grpc.CallOption) (ReplicationApi_SubscribeEnvelopesClient, error)
	QueryEnvelopes(ctx context.Context, in *QueryEnvelopesRequest, opts ...grpc.CallOption) (*QueryEnvelopesResponse, error)
}

type replicationApiClient struct {
	cc grpc.ClientConnInterface
}

func NewReplicationApiClient(cc grpc.ClientConnInterface) ReplicationApiClient {
	return &replicationApiClient{cc}
}

func (c *replicationApiClient) SubscribeEnvelopes(ctx context.Context, in *BatchSubscribeEnvelopesRequest, opts ...grpc.CallOption) (ReplicationApi_SubscribeEnvelopesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ReplicationApi_ServiceDesc.Streams[0], ReplicationApi_SubscribeEnvelopes_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &replicationApiSubscribeEnvelopesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ReplicationApi_SubscribeEnvelopesClient interface {
	Recv() (*GatewayEnvelope, error)
	grpc.ClientStream
}

type replicationApiSubscribeEnvelopesClient struct {
	grpc.ClientStream
}

func (x *replicationApiSubscribeEnvelopesClient) Recv() (*GatewayEnvelope, error) {
	m := new(GatewayEnvelope)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *replicationApiClient) QueryEnvelopes(ctx context.Context, in *QueryEnvelopesRequest, opts ...grpc.CallOption) (*QueryEnvelopesResponse, error) {
	out := new(QueryEnvelopesResponse)
	err := c.cc.Invoke(ctx, ReplicationApi_QueryEnvelopes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReplicationApiServer is the server API for ReplicationApi service.
// All implementations must embed UnimplementedReplicationApiServer
// for forward compatibility
type ReplicationApiServer interface {
	SubscribeEnvelopes(*BatchSubscribeEnvelopesRequest, ReplicationApi_SubscribeEnvelopesServer) error
	QueryEnvelopes(context.Context, *QueryEnvelopesRequest) (*QueryEnvelopesResponse, error)
	mustEmbedUnimplementedReplicationApiServer()
}

// UnimplementedReplicationApiServer must be embedded to have forward compatible implementations.
type UnimplementedReplicationApiServer struct {
}

func (UnimplementedReplicationApiServer) SubscribeEnvelopes(*BatchSubscribeEnvelopesRequest, ReplicationApi_SubscribeEnvelopesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeEnvelopes not implemented")
}
func (UnimplementedReplicationApiServer) QueryEnvelopes(context.Context, *QueryEnvelopesRequest) (*QueryEnvelopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryEnvelopes not implemented")
}
func (UnimplementedReplicationApiServer) mustEmbedUnimplementedReplicationApiServer() {}

// UnsafeReplicationApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReplicationApiServer will
// result in compilation errors.
type UnsafeReplicationApiServer interface {
	mustEmbedUnimplementedReplicationApiServer()
}

func RegisterReplicationApiServer(s grpc.ServiceRegistrar, srv ReplicationApiServer) {
	s.RegisterService(&ReplicationApi_ServiceDesc, srv)
}

func _ReplicationApi_SubscribeEnvelopes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BatchSubscribeEnvelopesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReplicationApiServer).SubscribeEnvelopes(m, &replicationApiSubscribeEnvelopesServer{stream})
}

type ReplicationApi_SubscribeEnvelopesServer interface {
	Send(*GatewayEnvelope) error
	grpc.ServerStream
}

type replicationApiSubscribeEnvelopesServer struct {
	grpc.ServerStream
}

func (x *replicationApiSubscribeEnvelopesServer) Send(m *GatewayEnvelope) error {
	return x.ServerStream.SendMsg(m)
}

func _ReplicationApi_QueryEnvelopes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEnvelopesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationApiServer).QueryEnvelopes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReplicationApi_QueryEnvelopes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationApiServer).QueryEnvelopes(ctx, req.(*QueryEnvelopesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReplicationApi_ServiceDesc is the grpc.ServiceDesc for ReplicationApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReplicationApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xmtp.xmtpv4.ReplicationApi",
	HandlerType: (*ReplicationApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryEnvelopes",
			Handler:    _ReplicationApi_QueryEnvelopes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeEnvelopes",
			Handler:       _ReplicationApi_SubscribeEnvelopes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "xmtpv4/message_api/message_api.proto",
}
