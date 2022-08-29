// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: v1/observability/observability.proto

package observability

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

// ObservabilityClient is the client API for Observability service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ObservabilityClient interface {
	Summary(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetPodNames(ctx context.Context, in *Request, opts ...grpc.CallOption) (*PodNameResponse, error)
}

type observabilityClient struct {
	cc grpc.ClientConnInterface
}

func NewObservabilityClient(cc grpc.ClientConnInterface) ObservabilityClient {
	return &observabilityClient{cc}
}

func (c *observabilityClient) Summary(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v1.observability.Observability/Summary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *observabilityClient) GetPodNames(ctx context.Context, in *Request, opts ...grpc.CallOption) (*PodNameResponse, error) {
	out := new(PodNameResponse)
	err := c.cc.Invoke(ctx, "/v1.observability.Observability/GetPodNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObservabilityServer is the server API for Observability service.
// All implementations must embed UnimplementedObservabilityServer
// for forward compatibility
type ObservabilityServer interface {
	Summary(context.Context, *Request) (*Response, error)
	GetPodNames(context.Context, *Request) (*PodNameResponse, error)
	mustEmbedUnimplementedObservabilityServer()
}

// UnimplementedObservabilityServer must be embedded to have forward compatible implementations.
type UnimplementedObservabilityServer struct {
}

func (UnimplementedObservabilityServer) Summary(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Summary not implemented")
}
func (UnimplementedObservabilityServer) GetPodNames(context.Context, *Request) (*PodNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPodNames not implemented")
}
func (UnimplementedObservabilityServer) mustEmbedUnimplementedObservabilityServer() {}

// UnsafeObservabilityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ObservabilityServer will
// result in compilation errors.
type UnsafeObservabilityServer interface {
	mustEmbedUnimplementedObservabilityServer()
}

func RegisterObservabilityServer(s grpc.ServiceRegistrar, srv ObservabilityServer) {
	s.RegisterService(&Observability_ServiceDesc, srv)
}

func _Observability_Summary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObservabilityServer).Summary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.observability.Observability/Summary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObservabilityServer).Summary(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Observability_GetPodNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObservabilityServer).GetPodNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.observability.Observability/GetPodNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObservabilityServer).GetPodNames(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Observability_ServiceDesc is the grpc.ServiceDesc for Observability service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Observability_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.observability.Observability",
	HandlerType: (*ObservabilityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Summary",
			Handler:    _Observability_Summary_Handler,
		},
		{
			MethodName: "GetPodNames",
			Handler:    _Observability_GetPodNames_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/observability/observability.proto",
}
