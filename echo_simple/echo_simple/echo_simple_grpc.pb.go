// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: echo_simple/echo_simple.proto

package echo_simple

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	EchoSimpleService_Echo_FullMethodName = "/EchoSimpleService/Echo"
)

// EchoSimpleServiceClient is the client API for EchoSimpleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EchoSimpleServiceClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type echoSimpleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoSimpleServiceClient(cc grpc.ClientConnInterface) EchoSimpleServiceClient {
	return &echoSimpleServiceClient{cc}
}

func (c *echoSimpleServiceClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, EchoSimpleService_Echo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoSimpleServiceServer is the server API for EchoSimpleService service.
// All implementations must embed UnimplementedEchoSimpleServiceServer
// for forward compatibility.
type EchoSimpleServiceServer interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	mustEmbedUnimplementedEchoSimpleServiceServer()
}

// UnimplementedEchoSimpleServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEchoSimpleServiceServer struct{}

func (UnimplementedEchoSimpleServiceServer) Echo(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedEchoSimpleServiceServer) mustEmbedUnimplementedEchoSimpleServiceServer() {}
func (UnimplementedEchoSimpleServiceServer) testEmbeddedByValue()                           {}

// UnsafeEchoSimpleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EchoSimpleServiceServer will
// result in compilation errors.
type UnsafeEchoSimpleServiceServer interface {
	mustEmbedUnimplementedEchoSimpleServiceServer()
}

func RegisterEchoSimpleServiceServer(s grpc.ServiceRegistrar, srv EchoSimpleServiceServer) {
	// If the following call pancis, it indicates UnimplementedEchoSimpleServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EchoSimpleService_ServiceDesc, srv)
}

func _EchoSimpleService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoSimpleServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EchoSimpleService_Echo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoSimpleServiceServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EchoSimpleService_ServiceDesc is the grpc.ServiceDesc for EchoSimpleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EchoSimpleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "EchoSimpleService",
	HandlerType: (*EchoSimpleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoSimpleService_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echo_simple/echo_simple.proto",
}
