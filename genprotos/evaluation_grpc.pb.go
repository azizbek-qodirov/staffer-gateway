// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: evaluation.proto

package genprotos

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
	EvaulationService_Create_FullMethodName = "/staff.EvaulationService/Create"
	EvaulationService_Get_FullMethodName    = "/staff.EvaulationService/Get"
	EvaulationService_Update_FullMethodName = "/staff.EvaulationService/Update"
	EvaulationService_Delete_FullMethodName = "/staff.EvaulationService/Delete"
	EvaulationService_GetAll_FullMethodName = "/staff.EvaulationService/GetAll"
)

// EvaulationServiceClient is the client API for EvaulationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EvaulationServiceClient interface {
	Create(ctx context.Context, in *EvaluationCreate, opts ...grpc.CallOption) (*Void, error)
	Get(ctx context.Context, in *Byid, opts ...grpc.CallOption) (*EvaluationUpdate, error)
	Update(ctx context.Context, in *EvaluationUpdate, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *Byid, opts ...grpc.CallOption) (*Void, error)
	GetAll(ctx context.Context, in *EvaluationGetAllReq, opts ...grpc.CallOption) (*EvaluationGetAllRes, error)
}

type evaulationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEvaulationServiceClient(cc grpc.ClientConnInterface) EvaulationServiceClient {
	return &evaulationServiceClient{cc}
}

func (c *evaulationServiceClient) Create(ctx context.Context, in *EvaluationCreate, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, EvaulationService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaulationServiceClient) Get(ctx context.Context, in *Byid, opts ...grpc.CallOption) (*EvaluationUpdate, error) {
	out := new(EvaluationUpdate)
	err := c.cc.Invoke(ctx, EvaulationService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaulationServiceClient) Update(ctx context.Context, in *EvaluationUpdate, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, EvaulationService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaulationServiceClient) Delete(ctx context.Context, in *Byid, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, EvaulationService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *evaulationServiceClient) GetAll(ctx context.Context, in *EvaluationGetAllReq, opts ...grpc.CallOption) (*EvaluationGetAllRes, error) {
	out := new(EvaluationGetAllRes)
	err := c.cc.Invoke(ctx, EvaulationService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EvaulationServiceServer is the server API for EvaulationService service.
// All implementations must embed UnimplementedEvaulationServiceServer
// for forward compatibility
type EvaulationServiceServer interface {
	Create(context.Context, *EvaluationCreate) (*Void, error)
	Get(context.Context, *Byid) (*EvaluationUpdate, error)
	Update(context.Context, *EvaluationUpdate) (*Void, error)
	Delete(context.Context, *Byid) (*Void, error)
	GetAll(context.Context, *EvaluationGetAllReq) (*EvaluationGetAllRes, error)
	mustEmbedUnimplementedEvaulationServiceServer()
}

// UnimplementedEvaulationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEvaulationServiceServer struct {
}

func (UnimplementedEvaulationServiceServer) Create(context.Context, *EvaluationCreate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedEvaulationServiceServer) Get(context.Context, *Byid) (*EvaluationUpdate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedEvaulationServiceServer) Update(context.Context, *EvaluationUpdate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedEvaulationServiceServer) Delete(context.Context, *Byid) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedEvaulationServiceServer) GetAll(context.Context, *EvaluationGetAllReq) (*EvaluationGetAllRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedEvaulationServiceServer) mustEmbedUnimplementedEvaulationServiceServer() {}

// UnsafeEvaulationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EvaulationServiceServer will
// result in compilation errors.
type UnsafeEvaulationServiceServer interface {
	mustEmbedUnimplementedEvaulationServiceServer()
}

func RegisterEvaulationServiceServer(s grpc.ServiceRegistrar, srv EvaulationServiceServer) {
	s.RegisterService(&EvaulationService_ServiceDesc, srv)
}

func _EvaulationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvaluationCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaulationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EvaulationService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaulationServiceServer).Create(ctx, req.(*EvaluationCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaulationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Byid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaulationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EvaulationService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaulationServiceServer).Get(ctx, req.(*Byid))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaulationService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvaluationUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaulationServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EvaulationService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaulationServiceServer).Update(ctx, req.(*EvaluationUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaulationService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Byid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaulationServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EvaulationService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaulationServiceServer).Delete(ctx, req.(*Byid))
	}
	return interceptor(ctx, in, info, handler)
}

func _EvaulationService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvaluationGetAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EvaulationServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EvaulationService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EvaulationServiceServer).GetAll(ctx, req.(*EvaluationGetAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

// EvaulationService_ServiceDesc is the grpc.ServiceDesc for EvaulationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EvaulationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "staff.EvaulationService",
	HandlerType: (*EvaulationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _EvaulationService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _EvaulationService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _EvaulationService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _EvaulationService_Delete_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _EvaulationService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "evaluation.proto",
}
