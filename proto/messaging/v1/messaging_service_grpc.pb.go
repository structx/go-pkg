// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/messaging/messaging_service.proto

package v1

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

// MessagingServiceV1Client is the client API for MessagingServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagingServiceV1Client interface {
	Publish(ctx context.Context, in *Envelope, opts ...grpc.CallOption) (*Stub, error)
	Subscribe(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (MessagingServiceV1_SubscribeClient, error)
	RequestResponse(ctx context.Context, in *Envelope, opts ...grpc.CallOption) (*Envelope, error)
}

type messagingServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewMessagingServiceV1Client(cc grpc.ClientConnInterface) MessagingServiceV1Client {
	return &messagingServiceV1Client{cc}
}

func (c *messagingServiceV1Client) Publish(ctx context.Context, in *Envelope, opts ...grpc.CallOption) (*Stub, error) {
	out := new(Stub)
	err := c.cc.Invoke(ctx, "/messaging.v1.MessagingServiceV1/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagingServiceV1Client) Subscribe(ctx context.Context, in *Subscription, opts ...grpc.CallOption) (MessagingServiceV1_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &MessagingServiceV1_ServiceDesc.Streams[0], "/messaging.v1.MessagingServiceV1/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &messagingServiceV1SubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MessagingServiceV1_SubscribeClient interface {
	Recv() (*Envelope, error)
	grpc.ClientStream
}

type messagingServiceV1SubscribeClient struct {
	grpc.ClientStream
}

func (x *messagingServiceV1SubscribeClient) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messagingServiceV1Client) RequestResponse(ctx context.Context, in *Envelope, opts ...grpc.CallOption) (*Envelope, error) {
	out := new(Envelope)
	err := c.cc.Invoke(ctx, "/messaging.v1.MessagingServiceV1/RequestResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessagingServiceV1Server is the server API for MessagingServiceV1 service.
// All implementations must embed UnimplementedMessagingServiceV1Server
// for forward compatibility
type MessagingServiceV1Server interface {
	Publish(context.Context, *Envelope) (*Stub, error)
	Subscribe(*Subscription, MessagingServiceV1_SubscribeServer) error
	RequestResponse(context.Context, *Envelope) (*Envelope, error)
	mustEmbedUnimplementedMessagingServiceV1Server()
}

// UnimplementedMessagingServiceV1Server must be embedded to have forward compatible implementations.
type UnimplementedMessagingServiceV1Server struct {
}

func (UnimplementedMessagingServiceV1Server) Publish(context.Context, *Envelope) (*Stub, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedMessagingServiceV1Server) Subscribe(*Subscription, MessagingServiceV1_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedMessagingServiceV1Server) RequestResponse(context.Context, *Envelope) (*Envelope, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestResponse not implemented")
}
func (UnimplementedMessagingServiceV1Server) mustEmbedUnimplementedMessagingServiceV1Server() {}

// UnsafeMessagingServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagingServiceV1Server will
// result in compilation errors.
type UnsafeMessagingServiceV1Server interface {
	mustEmbedUnimplementedMessagingServiceV1Server()
}

func RegisterMessagingServiceV1Server(s grpc.ServiceRegistrar, srv MessagingServiceV1Server) {
	s.RegisterService(&MessagingServiceV1_ServiceDesc, srv)
}

func _MessagingServiceV1_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Envelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceV1Server).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messaging.v1.MessagingServiceV1/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceV1Server).Publish(ctx, req.(*Envelope))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagingServiceV1_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Subscription)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessagingServiceV1Server).Subscribe(m, &messagingServiceV1SubscribeServer{stream})
}

type MessagingServiceV1_SubscribeServer interface {
	Send(*Envelope) error
	grpc.ServerStream
}

type messagingServiceV1SubscribeServer struct {
	grpc.ServerStream
}

func (x *messagingServiceV1SubscribeServer) Send(m *Envelope) error {
	return x.ServerStream.SendMsg(m)
}

func _MessagingServiceV1_RequestResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Envelope)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServiceV1Server).RequestResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/messaging.v1.MessagingServiceV1/RequestResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServiceV1Server).RequestResponse(ctx, req.(*Envelope))
	}
	return interceptor(ctx, in, info, handler)
}

// MessagingServiceV1_ServiceDesc is the grpc.ServiceDesc for MessagingServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessagingServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "messaging.v1.MessagingServiceV1",
	HandlerType: (*MessagingServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _MessagingServiceV1_Publish_Handler,
		},
		{
			MethodName: "RequestResponse",
			Handler:    _MessagingServiceV1_RequestResponse_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _MessagingServiceV1_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/messaging/messaging_service.proto",
}
