// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.1
// source: staffer-protos/notification.proto

package genprotos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ServiceNotification_Create_FullMethodName      = "/staff.ServiceNotification/Create"
	ServiceNotification_GetByUserId_FullMethodName = "/staff.ServiceNotification/GetByUserId"
	ServiceNotification_ReadeAll_FullMethodName    = "/staff.ServiceNotification/ReadeAll"
	ServiceNotification_SendAll_FullMethodName     = "/staff.ServiceNotification/SendAll"
)

// ServiceNotificationClient is the client API for ServiceNotification service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceNotificationClient interface {
	Create(ctx context.Context, in *CreateNotification, opts ...grpc.CallOption) (*Void, error)
	GetByUserId(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	ReadeAll(ctx context.Context, in *ReadeAllRequest, opts ...grpc.CallOption) (*Void, error)
	SendAll(ctx context.Context, in *SendByCompanyidToUsers, opts ...grpc.CallOption) (*Void, error)
}

type serviceNotificationClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceNotificationClient(cc grpc.ClientConnInterface) ServiceNotificationClient {
	return &serviceNotificationClient{cc}
}

func (c *serviceNotificationClient) Create(ctx context.Context, in *CreateNotification, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, ServiceNotification_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceNotificationClient) GetByUserId(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, ServiceNotification_GetByUserId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceNotificationClient) ReadeAll(ctx context.Context, in *ReadeAllRequest, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, ServiceNotification_ReadeAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceNotificationClient) SendAll(ctx context.Context, in *SendByCompanyidToUsers, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, ServiceNotification_SendAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceNotificationServer is the server API for ServiceNotification service.
// All implementations must embed UnimplementedServiceNotificationServer
// for forward compatibility
type ServiceNotificationServer interface {
	Create(context.Context, *CreateNotification) (*Void, error)
	GetByUserId(context.Context, *GetAllRequest) (*GetAllResponse, error)
	ReadeAll(context.Context, *ReadeAllRequest) (*Void, error)
	SendAll(context.Context, *SendByCompanyidToUsers) (*Void, error)
	mustEmbedUnimplementedServiceNotificationServer()
}

// UnimplementedServiceNotificationServer must be embedded to have forward compatible implementations.
type UnimplementedServiceNotificationServer struct {
}

func (UnimplementedServiceNotificationServer) Create(context.Context, *CreateNotification) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedServiceNotificationServer) GetByUserId(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUserId not implemented")
}
func (UnimplementedServiceNotificationServer) ReadeAll(context.Context, *ReadeAllRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadeAll not implemented")
}
func (UnimplementedServiceNotificationServer) SendAll(context.Context, *SendByCompanyidToUsers) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAll not implemented")
}
func (UnimplementedServiceNotificationServer) mustEmbedUnimplementedServiceNotificationServer() {}

// UnsafeServiceNotificationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceNotificationServer will
// result in compilation errors.
type UnsafeServiceNotificationServer interface {
	mustEmbedUnimplementedServiceNotificationServer()
}

func RegisterServiceNotificationServer(s grpc.ServiceRegistrar, srv ServiceNotificationServer) {
	s.RegisterService(&ServiceNotification_ServiceDesc, srv)
}

func _ServiceNotification_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceNotificationServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceNotification_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceNotificationServer).Create(ctx, req.(*CreateNotification))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceNotification_GetByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceNotificationServer).GetByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceNotification_GetByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceNotificationServer).GetByUserId(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceNotification_ReadeAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadeAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceNotificationServer).ReadeAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceNotification_ReadeAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceNotificationServer).ReadeAll(ctx, req.(*ReadeAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceNotification_SendAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendByCompanyidToUsers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceNotificationServer).SendAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceNotification_SendAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceNotificationServer).SendAll(ctx, req.(*SendByCompanyidToUsers))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceNotification_ServiceDesc is the grpc.ServiceDesc for ServiceNotification service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceNotification_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "staff.ServiceNotification",
	HandlerType: (*ServiceNotificationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ServiceNotification_Create_Handler,
		},
		{
			MethodName: "GetByUserId",
			Handler:    _ServiceNotification_GetByUserId_Handler,
		},
		{
			MethodName: "ReadeAll",
			Handler:    _ServiceNotification_ReadeAll_Handler,
		},
		{
			MethodName: "SendAll",
			Handler:    _ServiceNotification_SendAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staffer-protos/notification.proto",
}
