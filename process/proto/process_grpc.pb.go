// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: process/proto/process.proto

package proto

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

// ProcessServiceClient is the client API for ProcessService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProcessServiceClient interface {
	ExecuteCommand(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandResponse, error)
}

type processServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessServiceClient(cc grpc.ClientConnInterface) ProcessServiceClient {
	return &processServiceClient{cc}
}

func (c *processServiceClient) ExecuteCommand(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandResponse, error) {
	out := new(CommandResponse)
	err := c.cc.Invoke(ctx, "/process.ProcessService/ExecuteCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessServiceServer is the server API for ProcessService service.
// All implementations must embed UnimplementedProcessServiceServer
// for forward compatibility
type ProcessServiceServer interface {
	ExecuteCommand(context.Context, *CommandRequest) (*CommandResponse, error)
	mustEmbedUnimplementedProcessServiceServer()
}

// UnimplementedProcessServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProcessServiceServer struct {
}

func (UnimplementedProcessServiceServer) ExecuteCommand(context.Context, *CommandRequest) (*CommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteCommand not implemented")
}
func (UnimplementedProcessServiceServer) mustEmbedUnimplementedProcessServiceServer() {}

// UnsafeProcessServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProcessServiceServer will
// result in compilation errors.
type UnsafeProcessServiceServer interface {
	mustEmbedUnimplementedProcessServiceServer()
}

func RegisterProcessServiceServer(s grpc.ServiceRegistrar, srv ProcessServiceServer) {
	s.RegisterService(&ProcessService_ServiceDesc, srv)
}

func _ProcessService_ExecuteCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessServiceServer).ExecuteCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/process.ProcessService/ExecuteCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessServiceServer).ExecuteCommand(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProcessService_ServiceDesc is the grpc.ServiceDesc for ProcessService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProcessService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "process.ProcessService",
	HandlerType: (*ProcessServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteCommand",
			Handler:    _ProcessService_ExecuteCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "process/proto/process.proto",
}
