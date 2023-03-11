// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: clients.proto

package api

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

// ClientsServicesClient is the client API for ClientsServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientsServicesClient interface {
	// Countries
	GetCountries(ctx context.Context, in *RequestCountry, opts ...grpc.CallOption) (*ResponseCountries, error)
	CreateCountry(ctx context.Context, in *RequestCountry, opts ...grpc.CallOption) (*ResponseCountries, error)
	UpdateCountry(ctx context.Context, in *RequestCountry, opts ...grpc.CallOption) (*ResponseCountries, error)
	UpdateCountriesDeletionFlags(ctx context.Context, in *RequestCountriesDeletionFlags, opts ...grpc.CallOption) (*ResponseCountries, error)
	// Address types
	GetAddressesTypes(ctx context.Context, in *RequestAddressType, opts ...grpc.CallOption) (*ResponseAddressesTypes, error)
	CreateAddressType(ctx context.Context, in *RequestAddressType, opts ...grpc.CallOption) (*ResponseAddressesTypes, error)
	UpdateAddressType(ctx context.Context, in *RequestAddressType, opts ...grpc.CallOption) (*ResponseAddressesTypes, error)
	UpdateAddressTypeDeletionFlags(ctx context.Context, in *RequestAddressesTypesDeletionFlags, opts ...grpc.CallOption) (*ResponseAddressesTypes, error)
	// Deliveries
	GetDeliveries(ctx context.Context, in *RequestDelivery, opts ...grpc.CallOption) (*ResponseDeliveries, error)
	CreateDelivery(ctx context.Context, in *RequestDelivery, opts ...grpc.CallOption) (*ResponseDeliveries, error)
	UpdateDelivery(ctx context.Context, in *RequestDelivery, opts ...grpc.CallOption) (*ResponseDeliveries, error)
	UpdateDeliveriesDeletionFlags(ctx context.Context, in *RequestDeliveriesDeletionFlags, opts ...grpc.CallOption) (*ResponseDeliveries, error)
}

type clientsServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewClientsServicesClient(cc grpc.ClientConnInterface) ClientsServicesClient {
	return &clientsServicesClient{cc}
}

func (c *clientsServicesClient) GetCountries(ctx context.Context, in *RequestCountry, opts ...grpc.CallOption) (*ResponseCountries, error) {
	out := new(ResponseCountries)
	err := c.cc.Invoke(ctx, "/clients.ClientsServices/GetCountries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientsServicesClient) CreateCountry(ctx context.Context, in *RequestCountry, opts ...grpc.CallOption) (*ResponseCountries, error) {
	out := new(ResponseCountries)
	err := c.cc.Invoke(ctx, "/clients.ClientsServices/CreateCountry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientsServicesClient) UpdateCountry(ctx context.Context, in *RequestCountry, opts ...grpc.CallOption) (*ResponseCountries, error) {
	out := new(ResponseCountries)
	err := c.cc.Invoke(ctx, "/clients.ClientsServices/UpdateCountry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientsServicesClient) UpdateCountriesDeletionFlags(ctx context.Context, in *RequestCountriesDeletionFlags, opts ...grpc.CallOption) (*ResponseCountries, error) {
	out := new(ResponseCountries)
	err := c.cc.Invoke(ctx, "/clients.ClientsServices/UpdateCountriesDeletionFlags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientsServicesClient) GetAddressesTypes(ctx context.Context, in *RequestAddressType, opts ...grpc.CallOption) (*ResponseAddressesTypes, error) {
	out := new(ResponseAddressesTypes)
	err := c.cc.Invoke(ctx, "/clients.ClientsServices/GetAddressesTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientsServicesClient) CreateAddressType(ctx context.Context, in *RequestAddressType, opts ...grpc.CallOption) (*ResponseAddressesTypes, error) {
	out := new(ResponseAddressesTypes)
	err := c.cc.Invoke(ctx, "/clients.ClientsServices/CreateAddressType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientsServicesClient) UpdateAddressType(ctx context.Context, in *RequestAddressType, opts ...grpc.CallOption) (*ResponseAddressesTypes, error) {
	out := new(ResponseAddressesTypes)
	err := c.cc.Invoke(ctx, "/clients.ClientsServices/UpdateAddressType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientsServicesClient) UpdateAddressTypeDeletionFlags(ctx context.Context, in *RequestAddressesTypesDeletionFlags, opts ...grpc.CallOption) (*ResponseAddressesTypes, error) {
	out := new(ResponseAddressesTypes)
	err := c.cc.Invoke(ctx, "/clients.ClientsServices/UpdateAddressTypeDeletionFlags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientsServicesClient) GetDeliveries(ctx context.Context, in *RequestDelivery, opts ...grpc.CallOption) (*ResponseDeliveries, error) {
	out := new(ResponseDeliveries)
	err := c.cc.Invoke(ctx, "/clients.ClientsServices/GetDeliveries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientsServicesClient) CreateDelivery(ctx context.Context, in *RequestDelivery, opts ...grpc.CallOption) (*ResponseDeliveries, error) {
	out := new(ResponseDeliveries)
	err := c.cc.Invoke(ctx, "/clients.ClientsServices/CreateDelivery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientsServicesClient) UpdateDelivery(ctx context.Context, in *RequestDelivery, opts ...grpc.CallOption) (*ResponseDeliveries, error) {
	out := new(ResponseDeliveries)
	err := c.cc.Invoke(ctx, "/clients.ClientsServices/UpdateDelivery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientsServicesClient) UpdateDeliveriesDeletionFlags(ctx context.Context, in *RequestDeliveriesDeletionFlags, opts ...grpc.CallOption) (*ResponseDeliveries, error) {
	out := new(ResponseDeliveries)
	err := c.cc.Invoke(ctx, "/clients.ClientsServices/UpdateDeliveriesDeletionFlags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientsServicesServer is the server API for ClientsServices service.
// All implementations must embed UnimplementedClientsServicesServer
// for forward compatibility
type ClientsServicesServer interface {
	// Countries
	GetCountries(context.Context, *RequestCountry) (*ResponseCountries, error)
	CreateCountry(context.Context, *RequestCountry) (*ResponseCountries, error)
	UpdateCountry(context.Context, *RequestCountry) (*ResponseCountries, error)
	UpdateCountriesDeletionFlags(context.Context, *RequestCountriesDeletionFlags) (*ResponseCountries, error)
	// Address types
	GetAddressesTypes(context.Context, *RequestAddressType) (*ResponseAddressesTypes, error)
	CreateAddressType(context.Context, *RequestAddressType) (*ResponseAddressesTypes, error)
	UpdateAddressType(context.Context, *RequestAddressType) (*ResponseAddressesTypes, error)
	UpdateAddressTypeDeletionFlags(context.Context, *RequestAddressesTypesDeletionFlags) (*ResponseAddressesTypes, error)
	// Deliveries
	GetDeliveries(context.Context, *RequestDelivery) (*ResponseDeliveries, error)
	CreateDelivery(context.Context, *RequestDelivery) (*ResponseDeliveries, error)
	UpdateDelivery(context.Context, *RequestDelivery) (*ResponseDeliveries, error)
	UpdateDeliveriesDeletionFlags(context.Context, *RequestDeliveriesDeletionFlags) (*ResponseDeliveries, error)
	mustEmbedUnimplementedClientsServicesServer()
}

// UnimplementedClientsServicesServer must be embedded to have forward compatible implementations.
type UnimplementedClientsServicesServer struct {
}

func (UnimplementedClientsServicesServer) GetCountries(context.Context, *RequestCountry) (*ResponseCountries, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCountries not implemented")
}
func (UnimplementedClientsServicesServer) CreateCountry(context.Context, *RequestCountry) (*ResponseCountries, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCountry not implemented")
}
func (UnimplementedClientsServicesServer) UpdateCountry(context.Context, *RequestCountry) (*ResponseCountries, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCountry not implemented")
}
func (UnimplementedClientsServicesServer) UpdateCountriesDeletionFlags(context.Context, *RequestCountriesDeletionFlags) (*ResponseCountries, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCountriesDeletionFlags not implemented")
}
func (UnimplementedClientsServicesServer) GetAddressesTypes(context.Context, *RequestAddressType) (*ResponseAddressesTypes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddressesTypes not implemented")
}
func (UnimplementedClientsServicesServer) CreateAddressType(context.Context, *RequestAddressType) (*ResponseAddressesTypes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAddressType not implemented")
}
func (UnimplementedClientsServicesServer) UpdateAddressType(context.Context, *RequestAddressType) (*ResponseAddressesTypes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAddressType not implemented")
}
func (UnimplementedClientsServicesServer) UpdateAddressTypeDeletionFlags(context.Context, *RequestAddressesTypesDeletionFlags) (*ResponseAddressesTypes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAddressTypeDeletionFlags not implemented")
}
func (UnimplementedClientsServicesServer) GetDeliveries(context.Context, *RequestDelivery) (*ResponseDeliveries, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeliveries not implemented")
}
func (UnimplementedClientsServicesServer) CreateDelivery(context.Context, *RequestDelivery) (*ResponseDeliveries, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDelivery not implemented")
}
func (UnimplementedClientsServicesServer) UpdateDelivery(context.Context, *RequestDelivery) (*ResponseDeliveries, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDelivery not implemented")
}
func (UnimplementedClientsServicesServer) UpdateDeliveriesDeletionFlags(context.Context, *RequestDeliveriesDeletionFlags) (*ResponseDeliveries, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeliveriesDeletionFlags not implemented")
}
func (UnimplementedClientsServicesServer) mustEmbedUnimplementedClientsServicesServer() {}

// UnsafeClientsServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientsServicesServer will
// result in compilation errors.
type UnsafeClientsServicesServer interface {
	mustEmbedUnimplementedClientsServicesServer()
}

func RegisterClientsServicesServer(s grpc.ServiceRegistrar, srv ClientsServicesServer) {
	s.RegisterService(&ClientsServices_ServiceDesc, srv)
}

func _ClientsServices_GetCountries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCountry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientsServicesServer).GetCountries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clients.ClientsServices/GetCountries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientsServicesServer).GetCountries(ctx, req.(*RequestCountry))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientsServices_CreateCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCountry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientsServicesServer).CreateCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clients.ClientsServices/CreateCountry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientsServicesServer).CreateCountry(ctx, req.(*RequestCountry))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientsServices_UpdateCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCountry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientsServicesServer).UpdateCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clients.ClientsServices/UpdateCountry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientsServicesServer).UpdateCountry(ctx, req.(*RequestCountry))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientsServices_UpdateCountriesDeletionFlags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCountriesDeletionFlags)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientsServicesServer).UpdateCountriesDeletionFlags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clients.ClientsServices/UpdateCountriesDeletionFlags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientsServicesServer).UpdateCountriesDeletionFlags(ctx, req.(*RequestCountriesDeletionFlags))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientsServices_GetAddressesTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAddressType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientsServicesServer).GetAddressesTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clients.ClientsServices/GetAddressesTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientsServicesServer).GetAddressesTypes(ctx, req.(*RequestAddressType))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientsServices_CreateAddressType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAddressType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientsServicesServer).CreateAddressType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clients.ClientsServices/CreateAddressType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientsServicesServer).CreateAddressType(ctx, req.(*RequestAddressType))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientsServices_UpdateAddressType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAddressType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientsServicesServer).UpdateAddressType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clients.ClientsServices/UpdateAddressType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientsServicesServer).UpdateAddressType(ctx, req.(*RequestAddressType))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientsServices_UpdateAddressTypeDeletionFlags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAddressesTypesDeletionFlags)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientsServicesServer).UpdateAddressTypeDeletionFlags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clients.ClientsServices/UpdateAddressTypeDeletionFlags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientsServicesServer).UpdateAddressTypeDeletionFlags(ctx, req.(*RequestAddressesTypesDeletionFlags))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientsServices_GetDeliveries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDelivery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientsServicesServer).GetDeliveries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clients.ClientsServices/GetDeliveries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientsServicesServer).GetDeliveries(ctx, req.(*RequestDelivery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientsServices_CreateDelivery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDelivery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientsServicesServer).CreateDelivery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clients.ClientsServices/CreateDelivery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientsServicesServer).CreateDelivery(ctx, req.(*RequestDelivery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientsServices_UpdateDelivery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDelivery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientsServicesServer).UpdateDelivery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clients.ClientsServices/UpdateDelivery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientsServicesServer).UpdateDelivery(ctx, req.(*RequestDelivery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientsServices_UpdateDeliveriesDeletionFlags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDeliveriesDeletionFlags)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientsServicesServer).UpdateDeliveriesDeletionFlags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clients.ClientsServices/UpdateDeliveriesDeletionFlags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientsServicesServer).UpdateDeliveriesDeletionFlags(ctx, req.(*RequestDeliveriesDeletionFlags))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientsServices_ServiceDesc is the grpc.ServiceDesc for ClientsServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientsServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clients.ClientsServices",
	HandlerType: (*ClientsServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCountries",
			Handler:    _ClientsServices_GetCountries_Handler,
		},
		{
			MethodName: "CreateCountry",
			Handler:    _ClientsServices_CreateCountry_Handler,
		},
		{
			MethodName: "UpdateCountry",
			Handler:    _ClientsServices_UpdateCountry_Handler,
		},
		{
			MethodName: "UpdateCountriesDeletionFlags",
			Handler:    _ClientsServices_UpdateCountriesDeletionFlags_Handler,
		},
		{
			MethodName: "GetAddressesTypes",
			Handler:    _ClientsServices_GetAddressesTypes_Handler,
		},
		{
			MethodName: "CreateAddressType",
			Handler:    _ClientsServices_CreateAddressType_Handler,
		},
		{
			MethodName: "UpdateAddressType",
			Handler:    _ClientsServices_UpdateAddressType_Handler,
		},
		{
			MethodName: "UpdateAddressTypeDeletionFlags",
			Handler:    _ClientsServices_UpdateAddressTypeDeletionFlags_Handler,
		},
		{
			MethodName: "GetDeliveries",
			Handler:    _ClientsServices_GetDeliveries_Handler,
		},
		{
			MethodName: "CreateDelivery",
			Handler:    _ClientsServices_CreateDelivery_Handler,
		},
		{
			MethodName: "UpdateDelivery",
			Handler:    _ClientsServices_UpdateDelivery_Handler,
		},
		{
			MethodName: "UpdateDeliveriesDeletionFlags",
			Handler:    _ClientsServices_UpdateDeliveriesDeletionFlags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "clients.proto",
}