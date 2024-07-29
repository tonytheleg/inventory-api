// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.0
// - protoc             v5.27.2
// source: kessel/inventory/v1beta1/hosts_service.proto

package v1beta1

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
	HostsService_CreateRHELHost_FullMethodName = "/api.kessel.inventory.v1beta1.HostsService/CreateRHELHost"
	HostsService_UpdateRHELHost_FullMethodName = "/api.kessel.inventory.v1beta1.HostsService/UpdateRHELHost"
	HostsService_DeleteRHELHost_FullMethodName = "/api.kessel.inventory.v1beta1.HostsService/DeleteRHELHost"
)

// HostsServiceClient is the client API for HostsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HostsServiceClient interface {
	CreateRHELHost(ctx context.Context, in *CreateRHELHostRequest, opts ...grpc.CallOption) (*CreateRHELHostResponse, error)
	UpdateRHELHost(ctx context.Context, in *UpdateRHELHostRequest, opts ...grpc.CallOption) (*UpdateRHELHostResponse, error)
	DeleteRHELHost(ctx context.Context, in *DeleteRHELHostRequest, opts ...grpc.CallOption) (*DeleteRHELHostResponse, error)
}

type hostsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHostsServiceClient(cc grpc.ClientConnInterface) HostsServiceClient {
	return &hostsServiceClient{cc}
}

func (c *hostsServiceClient) CreateRHELHost(ctx context.Context, in *CreateRHELHostRequest, opts ...grpc.CallOption) (*CreateRHELHostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateRHELHostResponse)
	err := c.cc.Invoke(ctx, HostsService_CreateRHELHost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostsServiceClient) UpdateRHELHost(ctx context.Context, in *UpdateRHELHostRequest, opts ...grpc.CallOption) (*UpdateRHELHostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateRHELHostResponse)
	err := c.cc.Invoke(ctx, HostsService_UpdateRHELHost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostsServiceClient) DeleteRHELHost(ctx context.Context, in *DeleteRHELHostRequest, opts ...grpc.CallOption) (*DeleteRHELHostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteRHELHostResponse)
	err := c.cc.Invoke(ctx, HostsService_DeleteRHELHost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostsServiceServer is the server API for HostsService service.
// All implementations must embed UnimplementedHostsServiceServer
// for forward compatibility.
type HostsServiceServer interface {
	CreateRHELHost(context.Context, *CreateRHELHostRequest) (*CreateRHELHostResponse, error)
	UpdateRHELHost(context.Context, *UpdateRHELHostRequest) (*UpdateRHELHostResponse, error)
	DeleteRHELHost(context.Context, *DeleteRHELHostRequest) (*DeleteRHELHostResponse, error)
	mustEmbedUnimplementedHostsServiceServer()
}

// UnimplementedHostsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHostsServiceServer struct{}

func (UnimplementedHostsServiceServer) CreateRHELHost(context.Context, *CreateRHELHostRequest) (*CreateRHELHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRHELHost not implemented")
}
func (UnimplementedHostsServiceServer) UpdateRHELHost(context.Context, *UpdateRHELHostRequest) (*UpdateRHELHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRHELHost not implemented")
}
func (UnimplementedHostsServiceServer) DeleteRHELHost(context.Context, *DeleteRHELHostRequest) (*DeleteRHELHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRHELHost not implemented")
}
func (UnimplementedHostsServiceServer) mustEmbedUnimplementedHostsServiceServer() {}
func (UnimplementedHostsServiceServer) testEmbeddedByValue()                      {}

// UnsafeHostsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HostsServiceServer will
// result in compilation errors.
type UnsafeHostsServiceServer interface {
	mustEmbedUnimplementedHostsServiceServer()
}

func RegisterHostsServiceServer(s grpc.ServiceRegistrar, srv HostsServiceServer) {
	// If the following call pancis, it indicates UnimplementedHostsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HostsService_ServiceDesc, srv)
}

func _HostsService_CreateRHELHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRHELHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostsServiceServer).CreateRHELHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HostsService_CreateRHELHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostsServiceServer).CreateRHELHost(ctx, req.(*CreateRHELHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostsService_UpdateRHELHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRHELHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostsServiceServer).UpdateRHELHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HostsService_UpdateRHELHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostsServiceServer).UpdateRHELHost(ctx, req.(*UpdateRHELHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostsService_DeleteRHELHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRHELHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostsServiceServer).DeleteRHELHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HostsService_DeleteRHELHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostsServiceServer).DeleteRHELHost(ctx, req.(*DeleteRHELHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HostsService_ServiceDesc is the grpc.ServiceDesc for HostsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HostsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.kessel.inventory.v1beta1.HostsService",
	HandlerType: (*HostsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRHELHost",
			Handler:    _HostsService_CreateRHELHost_Handler,
		},
		{
			MethodName: "UpdateRHELHost",
			Handler:    _HostsService_UpdateRHELHost_Handler,
		},
		{
			MethodName: "DeleteRHELHost",
			Handler:    _HostsService_DeleteRHELHost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kessel/inventory/v1beta1/hosts_service.proto",
}
