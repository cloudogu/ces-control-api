// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: logging.proto

package logging

import (
	context "context"
	types "github.com/cloudogu/ces-control-api/generated/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DoguLogMessagesClient is the client API for DoguLogMessages service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DoguLogMessagesClient interface {
	GetForDogu(ctx context.Context, in *DoguLogMessageRequest, opts ...grpc.CallOption) (DoguLogMessages_GetForDoguClient, error)
	GetForDoguWithDate(ctx context.Context, in *DoguLogMessageWithDateRequest, opts ...grpc.CallOption) (DoguLogMessages_GetForDoguWithDateClient, error)
}

type doguLogMessagesClient struct {
	cc grpc.ClientConnInterface
}

func NewDoguLogMessagesClient(cc grpc.ClientConnInterface) DoguLogMessagesClient {
	return &doguLogMessagesClient{cc}
}

func (c *doguLogMessagesClient) GetForDogu(ctx context.Context, in *DoguLogMessageRequest, opts ...grpc.CallOption) (DoguLogMessages_GetForDoguClient, error) {
	stream, err := c.cc.NewStream(ctx, &DoguLogMessages_ServiceDesc.Streams[0], "/logging.DoguLogMessages/GetForDogu", opts...)
	if err != nil {
		return nil, err
	}
	x := &doguLogMessagesGetForDoguClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DoguLogMessages_GetForDoguClient interface {
	Recv() (*types.ChunkedDataResponse, error)
	grpc.ClientStream
}

type doguLogMessagesGetForDoguClient struct {
	grpc.ClientStream
}

func (x *doguLogMessagesGetForDoguClient) Recv() (*types.ChunkedDataResponse, error) {
	m := new(types.ChunkedDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *doguLogMessagesClient) GetForDoguWithDate(ctx context.Context, in *DoguLogMessageWithDateRequest, opts ...grpc.CallOption) (DoguLogMessages_GetForDoguWithDateClient, error) {
	stream, err := c.cc.NewStream(ctx, &DoguLogMessages_ServiceDesc.Streams[1], "/logging.DoguLogMessages/GetForDoguWithDate", opts...)
	if err != nil {
		return nil, err
	}
	x := &doguLogMessagesGetForDoguWithDateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DoguLogMessages_GetForDoguWithDateClient interface {
	Recv() (*DoguLogMessage, error)
	grpc.ClientStream
}

type doguLogMessagesGetForDoguWithDateClient struct {
	grpc.ClientStream
}

func (x *doguLogMessagesGetForDoguWithDateClient) Recv() (*DoguLogMessage, error) {
	m := new(DoguLogMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DoguLogMessagesServer is the server API for DoguLogMessages service.
// All implementations must embed UnimplementedDoguLogMessagesServer
// for forward compatibility
type DoguLogMessagesServer interface {
	GetForDogu(*DoguLogMessageRequest, DoguLogMessages_GetForDoguServer) error
	GetForDoguWithDate(*DoguLogMessageWithDateRequest, DoguLogMessages_GetForDoguWithDateServer) error
	mustEmbedUnimplementedDoguLogMessagesServer()
}

// UnimplementedDoguLogMessagesServer must be embedded to have forward compatible implementations.
type UnimplementedDoguLogMessagesServer struct {
}

func (UnimplementedDoguLogMessagesServer) GetForDogu(*DoguLogMessageRequest, DoguLogMessages_GetForDoguServer) error {
	return status.Errorf(codes.Unimplemented, "method GetForDogu not implemented")
}
func (UnimplementedDoguLogMessagesServer) GetForDoguWithDate(*DoguLogMessageWithDateRequest, DoguLogMessages_GetForDoguWithDateServer) error {
	return status.Errorf(codes.Unimplemented, "method GetForDoguWithDate not implemented")
}
func (UnimplementedDoguLogMessagesServer) mustEmbedUnimplementedDoguLogMessagesServer() {}

// UnsafeDoguLogMessagesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DoguLogMessagesServer will
// result in compilation errors.
type UnsafeDoguLogMessagesServer interface {
	mustEmbedUnimplementedDoguLogMessagesServer()
}

func RegisterDoguLogMessagesServer(s grpc.ServiceRegistrar, srv DoguLogMessagesServer) {
	s.RegisterService(&DoguLogMessages_ServiceDesc, srv)
}

func _DoguLogMessages_GetForDogu_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DoguLogMessageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DoguLogMessagesServer).GetForDogu(m, &doguLogMessagesGetForDoguServer{stream})
}

type DoguLogMessages_GetForDoguServer interface {
	Send(*types.ChunkedDataResponse) error
	grpc.ServerStream
}

type doguLogMessagesGetForDoguServer struct {
	grpc.ServerStream
}

func (x *doguLogMessagesGetForDoguServer) Send(m *types.ChunkedDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _DoguLogMessages_GetForDoguWithDate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DoguLogMessageWithDateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DoguLogMessagesServer).GetForDoguWithDate(m, &doguLogMessagesGetForDoguWithDateServer{stream})
}

type DoguLogMessages_GetForDoguWithDateServer interface {
	Send(*DoguLogMessage) error
	grpc.ServerStream
}

type doguLogMessagesGetForDoguWithDateServer struct {
	grpc.ServerStream
}

func (x *doguLogMessagesGetForDoguWithDateServer) Send(m *DoguLogMessage) error {
	return x.ServerStream.SendMsg(m)
}

// DoguLogMessages_ServiceDesc is the grpc.ServiceDesc for DoguLogMessages service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DoguLogMessages_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logging.DoguLogMessages",
	HandlerType: (*DoguLogMessagesServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetForDogu",
			Handler:       _DoguLogMessages_GetForDogu_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetForDoguWithDate",
			Handler:       _DoguLogMessages_GetForDoguWithDate_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "logging.proto",
}
