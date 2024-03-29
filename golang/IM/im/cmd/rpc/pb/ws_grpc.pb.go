// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: ws.proto

package pb

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

// WsClient is the client API for Ws service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WsClient interface {
	//返送单条消息
	PushSingleMsg(ctx context.Context, in *PushRoomMsgRequest, opts ...grpc.CallOption) (*SuccessReply, error)
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectReply, error)
	DisConnect(ctx context.Context, in *DisConnectRequest, opts ...grpc.CallOption) (*DisConnectReply, error)
}

type wsClient struct {
	cc grpc.ClientConnInterface
}

func NewWsClient(cc grpc.ClientConnInterface) WsClient {
	return &wsClient{cc}
}

func (c *wsClient) PushSingleMsg(ctx context.Context, in *PushRoomMsgRequest, opts ...grpc.CallOption) (*SuccessReply, error) {
	out := new(SuccessReply)
	err := c.cc.Invoke(ctx, "/pb.ws/PushSingleMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wsClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectReply, error) {
	out := new(ConnectReply)
	err := c.cc.Invoke(ctx, "/pb.ws/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wsClient) DisConnect(ctx context.Context, in *DisConnectRequest, opts ...grpc.CallOption) (*DisConnectReply, error) {
	out := new(DisConnectReply)
	err := c.cc.Invoke(ctx, "/pb.ws/DisConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WsServer is the server API for Ws service.
// All implementations must embed UnimplementedWsServer
// for forward compatibility
type WsServer interface {
	//返送单条消息
	PushSingleMsg(context.Context, *PushRoomMsgRequest) (*SuccessReply, error)
	Connect(context.Context, *ConnectRequest) (*ConnectReply, error)
	DisConnect(context.Context, *DisConnectRequest) (*DisConnectReply, error)
	mustEmbedUnimplementedWsServer()
}

// UnimplementedWsServer must be embedded to have forward compatible implementations.
type UnimplementedWsServer struct {
}

func (UnimplementedWsServer) PushSingleMsg(context.Context, *PushRoomMsgRequest) (*SuccessReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushSingleMsg not implemented")
}
func (UnimplementedWsServer) Connect(context.Context, *ConnectRequest) (*ConnectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedWsServer) DisConnect(context.Context, *DisConnectRequest) (*DisConnectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisConnect not implemented")
}
func (UnimplementedWsServer) mustEmbedUnimplementedWsServer() {}

// UnsafeWsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WsServer will
// result in compilation errors.
type UnsafeWsServer interface {
	mustEmbedUnimplementedWsServer()
}

func RegisterWsServer(s grpc.ServiceRegistrar, srv WsServer) {
	s.RegisterService(&Ws_ServiceDesc, srv)
}

func _Ws_PushSingleMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRoomMsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WsServer).PushSingleMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ws/PushSingleMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WsServer).PushSingleMsg(ctx, req.(*PushRoomMsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ws_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WsServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ws/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WsServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ws_DisConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WsServer).DisConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ws/DisConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WsServer).DisConnect(ctx, req.(*DisConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Ws_ServiceDesc is the grpc.ServiceDesc for Ws service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ws_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ws",
	HandlerType: (*WsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushSingleMsg",
			Handler:    _Ws_PushSingleMsg_Handler,
		},
		{
			MethodName: "Connect",
			Handler:    _Ws_Connect_Handler,
		},
		{
			MethodName: "DisConnect",
			Handler:    _Ws_DisConnect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ws.proto",
}
