// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.1
// source: staffer-protos/job_offer.proto

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
	JobOfferService_CreateJobOffer_FullMethodName  = "/staff.JobOfferService/CreateJobOffer"
	JobOfferService_GetByIdJobOffer_FullMethodName = "/staff.JobOfferService/GetByIdJobOffer"
	JobOfferService_GetAllJobOffers_FullMethodName = "/staff.JobOfferService/GetAllJobOffers"
	JobOfferService_UpdateJobOffer_FullMethodName  = "/staff.JobOfferService/UpdateJobOffer"
	JobOfferService_DeleteJobOffer_FullMethodName  = "/staff.JobOfferService/DeleteJobOffer"
)

// JobOfferServiceClient is the client API for JobOfferService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Xizmat interfeysi
type JobOfferServiceClient interface {
	CreateJobOffer(ctx context.Context, in *CreateJobOfferRequest, opts ...grpc.CallOption) (*JobOfferResponse, error)
	GetByIdJobOffer(ctx context.Context, in *GetByIdJobOfferRequest, opts ...grpc.CallOption) (*JobOfferResponse, error)
	GetAllJobOffers(ctx context.Context, in *GetAllJobOffersRequest, opts ...grpc.CallOption) (*JobOffersResponse, error)
	UpdateJobOffer(ctx context.Context, in *UpdateJobOfferRequest, opts ...grpc.CallOption) (*JobOfferResponse, error)
	DeleteJobOffer(ctx context.Context, in *DeleteJobOfferRequest, opts ...grpc.CallOption) (*JobOfferResponse, error)
}

type jobOfferServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobOfferServiceClient(cc grpc.ClientConnInterface) JobOfferServiceClient {
	return &jobOfferServiceClient{cc}
}

func (c *jobOfferServiceClient) CreateJobOffer(ctx context.Context, in *CreateJobOfferRequest, opts ...grpc.CallOption) (*JobOfferResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JobOfferResponse)
	err := c.cc.Invoke(ctx, JobOfferService_CreateJobOffer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobOfferServiceClient) GetByIdJobOffer(ctx context.Context, in *GetByIdJobOfferRequest, opts ...grpc.CallOption) (*JobOfferResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JobOfferResponse)
	err := c.cc.Invoke(ctx, JobOfferService_GetByIdJobOffer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobOfferServiceClient) GetAllJobOffers(ctx context.Context, in *GetAllJobOffersRequest, opts ...grpc.CallOption) (*JobOffersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JobOffersResponse)
	err := c.cc.Invoke(ctx, JobOfferService_GetAllJobOffers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobOfferServiceClient) UpdateJobOffer(ctx context.Context, in *UpdateJobOfferRequest, opts ...grpc.CallOption) (*JobOfferResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JobOfferResponse)
	err := c.cc.Invoke(ctx, JobOfferService_UpdateJobOffer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobOfferServiceClient) DeleteJobOffer(ctx context.Context, in *DeleteJobOfferRequest, opts ...grpc.CallOption) (*JobOfferResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JobOfferResponse)
	err := c.cc.Invoke(ctx, JobOfferService_DeleteJobOffer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobOfferServiceServer is the server API for JobOfferService service.
// All implementations must embed UnimplementedJobOfferServiceServer
// for forward compatibility
//
// Xizmat interfeysi
type JobOfferServiceServer interface {
	CreateJobOffer(context.Context, *CreateJobOfferRequest) (*JobOfferResponse, error)
	GetByIdJobOffer(context.Context, *GetByIdJobOfferRequest) (*JobOfferResponse, error)
	GetAllJobOffers(context.Context, *GetAllJobOffersRequest) (*JobOffersResponse, error)
	UpdateJobOffer(context.Context, *UpdateJobOfferRequest) (*JobOfferResponse, error)
	DeleteJobOffer(context.Context, *DeleteJobOfferRequest) (*JobOfferResponse, error)
	mustEmbedUnimplementedJobOfferServiceServer()
}

// UnimplementedJobOfferServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJobOfferServiceServer struct {
}

func (UnimplementedJobOfferServiceServer) CreateJobOffer(context.Context, *CreateJobOfferRequest) (*JobOfferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJobOffer not implemented")
}
func (UnimplementedJobOfferServiceServer) GetByIdJobOffer(context.Context, *GetByIdJobOfferRequest) (*JobOfferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdJobOffer not implemented")
}
func (UnimplementedJobOfferServiceServer) GetAllJobOffers(context.Context, *GetAllJobOffersRequest) (*JobOffersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllJobOffers not implemented")
}
func (UnimplementedJobOfferServiceServer) UpdateJobOffer(context.Context, *UpdateJobOfferRequest) (*JobOfferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateJobOffer not implemented")
}
func (UnimplementedJobOfferServiceServer) DeleteJobOffer(context.Context, *DeleteJobOfferRequest) (*JobOfferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteJobOffer not implemented")
}
func (UnimplementedJobOfferServiceServer) mustEmbedUnimplementedJobOfferServiceServer() {}

// UnsafeJobOfferServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobOfferServiceServer will
// result in compilation errors.
type UnsafeJobOfferServiceServer interface {
	mustEmbedUnimplementedJobOfferServiceServer()
}

func RegisterJobOfferServiceServer(s grpc.ServiceRegistrar, srv JobOfferServiceServer) {
	s.RegisterService(&JobOfferService_ServiceDesc, srv)
}

func _JobOfferService_CreateJobOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJobOfferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobOfferServiceServer).CreateJobOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobOfferService_CreateJobOffer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobOfferServiceServer).CreateJobOffer(ctx, req.(*CreateJobOfferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobOfferService_GetByIdJobOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdJobOfferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobOfferServiceServer).GetByIdJobOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobOfferService_GetByIdJobOffer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobOfferServiceServer).GetByIdJobOffer(ctx, req.(*GetByIdJobOfferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobOfferService_GetAllJobOffers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllJobOffersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobOfferServiceServer).GetAllJobOffers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobOfferService_GetAllJobOffers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobOfferServiceServer).GetAllJobOffers(ctx, req.(*GetAllJobOffersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobOfferService_UpdateJobOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateJobOfferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobOfferServiceServer).UpdateJobOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobOfferService_UpdateJobOffer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobOfferServiceServer).UpdateJobOffer(ctx, req.(*UpdateJobOfferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobOfferService_DeleteJobOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteJobOfferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobOfferServiceServer).DeleteJobOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobOfferService_DeleteJobOffer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobOfferServiceServer).DeleteJobOffer(ctx, req.(*DeleteJobOfferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JobOfferService_ServiceDesc is the grpc.ServiceDesc for JobOfferService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JobOfferService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "staff.JobOfferService",
	HandlerType: (*JobOfferServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJobOffer",
			Handler:    _JobOfferService_CreateJobOffer_Handler,
		},
		{
			MethodName: "GetByIdJobOffer",
			Handler:    _JobOfferService_GetByIdJobOffer_Handler,
		},
		{
			MethodName: "GetAllJobOffers",
			Handler:    _JobOfferService_GetAllJobOffers_Handler,
		},
		{
			MethodName: "UpdateJobOffer",
			Handler:    _JobOfferService_UpdateJobOffer_Handler,
		},
		{
			MethodName: "DeleteJobOffer",
			Handler:    _JobOfferService_DeleteJobOffer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staffer-protos/job_offer.proto",
}
