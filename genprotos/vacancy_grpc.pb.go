// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.1
// source: staffer-protos/vacancy.proto

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
	VacancyService_Create_FullMethodName          = "/staff.VacancyService/Create"
	VacancyService_GetByID_FullMethodName         = "/staff.VacancyService/GetByID"
	VacancyService_Update_FullMethodName          = "/staff.VacancyService/Update"
	VacancyService_Delete_FullMethodName          = "/staff.VacancyService/Delete"
	VacancyService_GetAll_FullMethodName          = "/staff.VacancyService/GetAll"
	VacancyService_GetApplications_FullMethodName = "/staff.VacancyService/GetApplications"
	VacancyService_GetOffers_FullMethodName       = "/staff.VacancyService/GetOffers"
)

// VacancyServiceClient is the client API for VacancyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VacancyServiceClient interface {
	Create(ctx context.Context, in *VacancyCreate, opts ...grpc.CallOption) (*Void, error)
	GetByID(ctx context.Context, in *Byid, opts ...grpc.CallOption) (*VacancyGetResUpdateReq, error)
	Update(ctx context.Context, in *VacancyGetResUpdateReq, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *Byid, opts ...grpc.CallOption) (*Void, error)
	GetAll(ctx context.Context, in *VacancyGetAllReq, opts ...grpc.CallOption) (*VacancyGetAllRes, error)
	GetApplications(ctx context.Context, in *Byid, opts ...grpc.CallOption) (*VacancyApplicationsRes, error)
	GetOffers(ctx context.Context, in *Byid, opts ...grpc.CallOption) (*VacancyOffersRes, error)
}

type vacancyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVacancyServiceClient(cc grpc.ClientConnInterface) VacancyServiceClient {
	return &vacancyServiceClient{cc}
}

func (c *vacancyServiceClient) Create(ctx context.Context, in *VacancyCreate, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, VacancyService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacancyServiceClient) GetByID(ctx context.Context, in *Byid, opts ...grpc.CallOption) (*VacancyGetResUpdateReq, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VacancyGetResUpdateReq)
	err := c.cc.Invoke(ctx, VacancyService_GetByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacancyServiceClient) Update(ctx context.Context, in *VacancyGetResUpdateReq, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, VacancyService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacancyServiceClient) Delete(ctx context.Context, in *Byid, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, VacancyService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacancyServiceClient) GetAll(ctx context.Context, in *VacancyGetAllReq, opts ...grpc.CallOption) (*VacancyGetAllRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VacancyGetAllRes)
	err := c.cc.Invoke(ctx, VacancyService_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacancyServiceClient) GetApplications(ctx context.Context, in *Byid, opts ...grpc.CallOption) (*VacancyApplicationsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VacancyApplicationsRes)
	err := c.cc.Invoke(ctx, VacancyService_GetApplications_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacancyServiceClient) GetOffers(ctx context.Context, in *Byid, opts ...grpc.CallOption) (*VacancyOffersRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VacancyOffersRes)
	err := c.cc.Invoke(ctx, VacancyService_GetOffers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VacancyServiceServer is the server API for VacancyService service.
// All implementations must embed UnimplementedVacancyServiceServer
// for forward compatibility
type VacancyServiceServer interface {
	Create(context.Context, *VacancyCreate) (*Void, error)
	GetByID(context.Context, *Byid) (*VacancyGetResUpdateReq, error)
	Update(context.Context, *VacancyGetResUpdateReq) (*Void, error)
	Delete(context.Context, *Byid) (*Void, error)
	GetAll(context.Context, *VacancyGetAllReq) (*VacancyGetAllRes, error)
	GetApplications(context.Context, *Byid) (*VacancyApplicationsRes, error)
	GetOffers(context.Context, *Byid) (*VacancyOffersRes, error)
	mustEmbedUnimplementedVacancyServiceServer()
}

// UnimplementedVacancyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVacancyServiceServer struct {
}

func (UnimplementedVacancyServiceServer) Create(context.Context, *VacancyCreate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedVacancyServiceServer) GetByID(context.Context, *Byid) (*VacancyGetResUpdateReq, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedVacancyServiceServer) Update(context.Context, *VacancyGetResUpdateReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedVacancyServiceServer) Delete(context.Context, *Byid) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedVacancyServiceServer) GetAll(context.Context, *VacancyGetAllReq) (*VacancyGetAllRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedVacancyServiceServer) GetApplications(context.Context, *Byid) (*VacancyApplicationsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApplications not implemented")
}
func (UnimplementedVacancyServiceServer) GetOffers(context.Context, *Byid) (*VacancyOffersRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOffers not implemented")
}
func (UnimplementedVacancyServiceServer) mustEmbedUnimplementedVacancyServiceServer() {}

// UnsafeVacancyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VacancyServiceServer will
// result in compilation errors.
type UnsafeVacancyServiceServer interface {
	mustEmbedUnimplementedVacancyServiceServer()
}

func RegisterVacancyServiceServer(s grpc.ServiceRegistrar, srv VacancyServiceServer) {
	s.RegisterService(&VacancyService_ServiceDesc, srv)
}

func _VacancyService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VacancyCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacancyServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VacancyService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacancyServiceServer).Create(ctx, req.(*VacancyCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacancyService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Byid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacancyServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VacancyService_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacancyServiceServer).GetByID(ctx, req.(*Byid))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacancyService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VacancyGetResUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacancyServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VacancyService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacancyServiceServer).Update(ctx, req.(*VacancyGetResUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacancyService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Byid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacancyServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VacancyService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacancyServiceServer).Delete(ctx, req.(*Byid))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacancyService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VacancyGetAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacancyServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VacancyService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacancyServiceServer).GetAll(ctx, req.(*VacancyGetAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacancyService_GetApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Byid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacancyServiceServer).GetApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VacancyService_GetApplications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacancyServiceServer).GetApplications(ctx, req.(*Byid))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacancyService_GetOffers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Byid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacancyServiceServer).GetOffers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VacancyService_GetOffers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacancyServiceServer).GetOffers(ctx, req.(*Byid))
	}
	return interceptor(ctx, in, info, handler)
}

// VacancyService_ServiceDesc is the grpc.ServiceDesc for VacancyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VacancyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "staff.VacancyService",
	HandlerType: (*VacancyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _VacancyService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _VacancyService_GetByID_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _VacancyService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _VacancyService_Delete_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _VacancyService_GetAll_Handler,
		},
		{
			MethodName: "GetApplications",
			Handler:    _VacancyService_GetApplications_Handler,
		},
		{
			MethodName: "GetOffers",
			Handler:    _VacancyService_GetOffers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staffer-protos/vacancy.proto",
}
