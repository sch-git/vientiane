// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package vientiane

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

// VientianeServiceClient is the client API for VientianeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VientianeServiceClient interface {
	HealthCheck(ctx context.Context, in *HealthCheckReq, opts ...grpc.CallOption) (*HealthCheckRes, error)
	// account
	GetAccount(ctx context.Context, in *GetAccountReq, opts ...grpc.CallOption) (*GetAccountRes, error)
	ListAccount(ctx context.Context, in *ListAccountReq, opts ...grpc.CallOption) (*ListAccountRes, error)
	// category
	ListCategory(ctx context.Context, in *ListCategoryReq, opts ...grpc.CallOption) (*ListCategoryRes, error)
}

type vientianeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVientianeServiceClient(cc grpc.ClientConnInterface) VientianeServiceClient {
	return &vientianeServiceClient{cc}
}

func (c *vientianeServiceClient) HealthCheck(ctx context.Context, in *HealthCheckReq, opts ...grpc.CallOption) (*HealthCheckRes, error) {
	out := new(HealthCheckRes)
	err := c.cc.Invoke(ctx, "/vientiane.VientianeService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vientianeServiceClient) GetAccount(ctx context.Context, in *GetAccountReq, opts ...grpc.CallOption) (*GetAccountRes, error) {
	out := new(GetAccountRes)
	err := c.cc.Invoke(ctx, "/vientiane.VientianeService/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vientianeServiceClient) ListAccount(ctx context.Context, in *ListAccountReq, opts ...grpc.CallOption) (*ListAccountRes, error) {
	out := new(ListAccountRes)
	err := c.cc.Invoke(ctx, "/vientiane.VientianeService/ListAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vientianeServiceClient) ListCategory(ctx context.Context, in *ListCategoryReq, opts ...grpc.CallOption) (*ListCategoryRes, error) {
	out := new(ListCategoryRes)
	err := c.cc.Invoke(ctx, "/vientiane.VientianeService/ListCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VientianeServiceServer is the server API for VientianeService service.
// All implementations must embed UnimplementedVientianeServiceServer
// for forward compatibility
type VientianeServiceServer interface {
	HealthCheck(context.Context, *HealthCheckReq) (*HealthCheckRes, error)
	// account
	GetAccount(context.Context, *GetAccountReq) (*GetAccountRes, error)
	ListAccount(context.Context, *ListAccountReq) (*ListAccountRes, error)
	// category
	ListCategory(context.Context, *ListCategoryReq) (*ListCategoryRes, error)
	mustEmbedUnimplementedVientianeServiceServer()
}

// UnimplementedVientianeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVientianeServiceServer struct {
}

func (UnimplementedVientianeServiceServer) HealthCheck(context.Context, *HealthCheckReq) (*HealthCheckRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedVientianeServiceServer) GetAccount(context.Context, *GetAccountReq) (*GetAccountRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedVientianeServiceServer) ListAccount(context.Context, *ListAccountReq) (*ListAccountRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccount not implemented")
}
func (UnimplementedVientianeServiceServer) ListCategory(context.Context, *ListCategoryReq) (*ListCategoryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCategory not implemented")
}
func (UnimplementedVientianeServiceServer) mustEmbedUnimplementedVientianeServiceServer() {}

// UnsafeVientianeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VientianeServiceServer will
// result in compilation errors.
type UnsafeVientianeServiceServer interface {
	mustEmbedUnimplementedVientianeServiceServer()
}

func RegisterVientianeServiceServer(s grpc.ServiceRegistrar, srv VientianeServiceServer) {
	s.RegisterService(&VientianeService_ServiceDesc, srv)
}

func _VientianeService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VientianeServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vientiane.VientianeService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VientianeServiceServer).HealthCheck(ctx, req.(*HealthCheckReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VientianeService_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VientianeServiceServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vientiane.VientianeService/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VientianeServiceServer).GetAccount(ctx, req.(*GetAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VientianeService_ListAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VientianeServiceServer).ListAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vientiane.VientianeService/ListAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VientianeServiceServer).ListAccount(ctx, req.(*ListAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VientianeService_ListCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCategoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VientianeServiceServer).ListCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vientiane.VientianeService/ListCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VientianeServiceServer).ListCategory(ctx, req.(*ListCategoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

// VientianeService_ServiceDesc is the grpc.ServiceDesc for VientianeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VientianeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vientiane.VientianeService",
	HandlerType: (*VientianeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _VientianeService_HealthCheck_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _VientianeService_GetAccount_Handler,
		},
		{
			MethodName: "ListAccount",
			Handler:    _VientianeService_ListAccount_Handler,
		},
		{
			MethodName: "ListCategory",
			Handler:    _VientianeService_ListCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vientiane.proto",
}
