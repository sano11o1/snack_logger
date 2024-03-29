// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: logger.proto

// packageの宣言

package grpc

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
	LogService_Log_FullMethodName = "/logger.LogService/Log"
)

// LogServiceClient is the client API for LogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogServiceClient interface {
	// サービスが持つメソッドの定義
	Log(ctx context.Context, opts ...grpc.CallOption) (LogService_LogClient, error)
}

type logServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogServiceClient(cc grpc.ClientConnInterface) LogServiceClient {
	return &logServiceClient{cc}
}

func (c *logServiceClient) Log(ctx context.Context, opts ...grpc.CallOption) (LogService_LogClient, error) {
	stream, err := c.cc.NewStream(ctx, &LogService_ServiceDesc.Streams[0], LogService_Log_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &logServiceLogClient{stream}
	return x, nil
}

type LogService_LogClient interface {
	Send(*LogRequest) error
	CloseAndRecv() (*LogResponse, error)
	grpc.ClientStream
}

type logServiceLogClient struct {
	grpc.ClientStream
}

func (x *logServiceLogClient) Send(m *LogRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *logServiceLogClient) CloseAndRecv() (*LogResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(LogResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogServiceServer is the server API for LogService service.
// All implementations must embed UnimplementedLogServiceServer
// for forward compatibility
type LogServiceServer interface {
	// サービスが持つメソッドの定義
	Log(LogService_LogServer) error
	mustEmbedUnimplementedLogServiceServer()
}

// UnimplementedLogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLogServiceServer struct {
}

func (UnimplementedLogServiceServer) Log(LogService_LogServer) error {
	return status.Errorf(codes.Unimplemented, "method Log not implemented")
}
func (UnimplementedLogServiceServer) mustEmbedUnimplementedLogServiceServer() {}

// UnsafeLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogServiceServer will
// result in compilation errors.
type UnsafeLogServiceServer interface {
	mustEmbedUnimplementedLogServiceServer()
}

func RegisterLogServiceServer(s grpc.ServiceRegistrar, srv LogServiceServer) {
	s.RegisterService(&LogService_ServiceDesc, srv)
}

func _LogService_Log_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LogServiceServer).Log(&logServiceLogServer{stream})
}

type LogService_LogServer interface {
	SendAndClose(*LogResponse) error
	Recv() (*LogRequest, error)
	grpc.ServerStream
}

type logServiceLogServer struct {
	grpc.ServerStream
}

func (x *logServiceLogServer) SendAndClose(m *LogResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *logServiceLogServer) Recv() (*LogRequest, error) {
	m := new(LogRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogService_ServiceDesc is the grpc.ServiceDesc for LogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logger.LogService",
	HandlerType: (*LogServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Log",
			Handler:       _LogService_Log_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "logger.proto",
}
