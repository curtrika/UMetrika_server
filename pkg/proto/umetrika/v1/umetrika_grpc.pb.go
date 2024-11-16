// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.23.2
// source: umetrika/v1/umetrika.proto

package v1

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
	UMetrika_Ping_FullMethodName        = "/umetrika.UMetrika/Ping"
	UMetrika_CreateOwner_FullMethodName = "/umetrika.UMetrika/CreateOwner"
)

// UMetrikaClient is the client API for UMetrika service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UMetrikaClient interface {
	// Ping RPC
	Ping(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*PingMessage, error)
	CreateOwner(ctx context.Context, in *OwnerPost, opts ...grpc.CallOption) (*OwnerResult, error)
}

type uMetrikaClient struct {
	cc grpc.ClientConnInterface
}

func NewUMetrikaClient(cc grpc.ClientConnInterface) UMetrikaClient {
	return &uMetrikaClient{cc}
}

func (c *uMetrikaClient) Ping(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*PingMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PingMessage)
	err := c.cc.Invoke(ctx, UMetrika_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uMetrikaClient) CreateOwner(ctx context.Context, in *OwnerPost, opts ...grpc.CallOption) (*OwnerResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OwnerResult)
	err := c.cc.Invoke(ctx, UMetrika_CreateOwner_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UMetrikaServer is the server API for UMetrika service.
// All implementations must embed UnimplementedUMetrikaServer
// for forward compatibility.
type UMetrikaServer interface {
	// Ping RPC
	Ping(context.Context, *EmptyMessage) (*PingMessage, error)
	CreateOwner(context.Context, *OwnerPost) (*OwnerResult, error)
	mustEmbedUnimplementedUMetrikaServer()
}

// UnimplementedUMetrikaServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUMetrikaServer struct{}

func (UnimplementedUMetrikaServer) Ping(context.Context, *EmptyMessage) (*PingMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedUMetrikaServer) CreateOwner(context.Context, *OwnerPost) (*OwnerResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOwner not implemented")
}
func (UnimplementedUMetrikaServer) mustEmbedUnimplementedUMetrikaServer() {}
func (UnimplementedUMetrikaServer) testEmbeddedByValue()                  {}

// UnsafeUMetrikaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UMetrikaServer will
// result in compilation errors.
type UnsafeUMetrikaServer interface {
	mustEmbedUnimplementedUMetrikaServer()
}

func RegisterUMetrikaServer(s grpc.ServiceRegistrar, srv UMetrikaServer) {
	// If the following call pancis, it indicates UnimplementedUMetrikaServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UMetrika_ServiceDesc, srv)
}

func _UMetrika_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UMetrikaServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UMetrika_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UMetrikaServer).Ping(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _UMetrika_CreateOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OwnerPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UMetrikaServer).CreateOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UMetrika_CreateOwner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UMetrikaServer).CreateOwner(ctx, req.(*OwnerPost))
	}
	return interceptor(ctx, in, info, handler)
}

// UMetrika_ServiceDesc is the grpc.ServiceDesc for UMetrika service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UMetrika_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "umetrika.UMetrika",
	HandlerType: (*UMetrikaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _UMetrika_Ping_Handler,
		},
		{
			MethodName: "CreateOwner",
			Handler:    _UMetrika_CreateOwner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "umetrika/v1/umetrika.proto",
}
