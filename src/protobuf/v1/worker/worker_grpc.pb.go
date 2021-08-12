// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package worker

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

// WorkerClient is the client API for Worker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkerClient interface {
	GetWorkerStatus(ctx context.Context, in *WorkerRequest, opts ...grpc.CallOption) (*WorkerResponse, error)
	Start(ctx context.Context, in *WorkerRequest, opts ...grpc.CallOption) (*WorkerResponse, error)
	Stop(ctx context.Context, in *WorkerRequest, opts ...grpc.CallOption) (*WorkerResponse, error)
	Convert(ctx context.Context, in *WorkerRequest, opts ...grpc.CallOption) (*WorkerResponse, error)
}

type workerClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerClient(cc grpc.ClientConnInterface) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) GetWorkerStatus(ctx context.Context, in *WorkerRequest, opts ...grpc.CallOption) (*WorkerResponse, error) {
	out := new(WorkerResponse)
	err := c.cc.Invoke(ctx, "/v1.worker.Worker/GetWorkerStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Start(ctx context.Context, in *WorkerRequest, opts ...grpc.CallOption) (*WorkerResponse, error) {
	out := new(WorkerResponse)
	err := c.cc.Invoke(ctx, "/v1.worker.Worker/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Stop(ctx context.Context, in *WorkerRequest, opts ...grpc.CallOption) (*WorkerResponse, error) {
	out := new(WorkerResponse)
	err := c.cc.Invoke(ctx, "/v1.worker.Worker/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerClient) Convert(ctx context.Context, in *WorkerRequest, opts ...grpc.CallOption) (*WorkerResponse, error) {
	out := new(WorkerResponse)
	err := c.cc.Invoke(ctx, "/v1.worker.Worker/Convert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServer is the server API for Worker service.
// All implementations must embed UnimplementedWorkerServer
// for forward compatibility
type WorkerServer interface {
	GetWorkerStatus(context.Context, *WorkerRequest) (*WorkerResponse, error)
	Start(context.Context, *WorkerRequest) (*WorkerResponse, error)
	Stop(context.Context, *WorkerRequest) (*WorkerResponse, error)
	Convert(context.Context, *WorkerRequest) (*WorkerResponse, error)
	mustEmbedUnimplementedWorkerServer()
}

// UnimplementedWorkerServer must be embedded to have forward compatible implementations.
type UnimplementedWorkerServer struct {
}

func (UnimplementedWorkerServer) GetWorkerStatus(context.Context, *WorkerRequest) (*WorkerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkerStatus not implemented")
}
func (UnimplementedWorkerServer) Start(context.Context, *WorkerRequest) (*WorkerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedWorkerServer) Stop(context.Context, *WorkerRequest) (*WorkerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedWorkerServer) Convert(context.Context, *WorkerRequest) (*WorkerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Convert not implemented")
}
func (UnimplementedWorkerServer) mustEmbedUnimplementedWorkerServer() {}

// UnsafeWorkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkerServer will
// result in compilation errors.
type UnsafeWorkerServer interface {
	mustEmbedUnimplementedWorkerServer()
}

func RegisterWorkerServer(s grpc.ServiceRegistrar, srv WorkerServer) {
	s.RegisterService(&Worker_ServiceDesc, srv)
}

func _Worker_GetWorkerStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).GetWorkerStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.worker.Worker/GetWorkerStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).GetWorkerStatus(ctx, req.(*WorkerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.worker.Worker/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Start(ctx, req.(*WorkerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.worker.Worker/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Stop(ctx, req.(*WorkerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Worker_Convert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServer).Convert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.worker.Worker/Convert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServer).Convert(ctx, req.(*WorkerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Worker_ServiceDesc is the grpc.ServiceDesc for Worker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Worker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.worker.Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWorkerStatus",
			Handler:    _Worker_GetWorkerStatus_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Worker_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Worker_Stop_Handler,
		},
		{
			MethodName: "Convert",
			Handler:    _Worker_Convert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/worker/worker.proto",
}
