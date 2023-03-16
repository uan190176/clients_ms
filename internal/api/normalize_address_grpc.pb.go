// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: normalize_address.proto

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

// NormalizeAddressServicesServerClient is the client API for NormalizeAddressServicesServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NormalizeAddressServicesServerClient interface {
	NormalizeAddress(ctx context.Context, in *RequestNormalizeAddress, opts ...grpc.CallOption) (*ResponseNormalizeAddress, error)
}

type normalizeAddressServicesServerClient struct {
	cc grpc.ClientConnInterface
}

func NewNormalizeAddressServicesServerClient(cc grpc.ClientConnInterface) NormalizeAddressServicesServerClient {
	return &normalizeAddressServicesServerClient{cc}
}

func (c *normalizeAddressServicesServerClient) NormalizeAddress(ctx context.Context, in *RequestNormalizeAddress, opts ...grpc.CallOption) (*ResponseNormalizeAddress, error) {
	out := new(ResponseNormalizeAddress)
	err := c.cc.Invoke(ctx, "/normalize_address.NormalizeAddressServicesServer/NormalizeAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NormalizeAddressServicesServerServer is the server API for NormalizeAddressServicesServer service.
// All implementations must embed UnimplementedNormalizeAddressServicesServerServer
// for forward compatibility
type NormalizeAddressServicesServerServer interface {
	NormalizeAddress(context.Context, *RequestNormalizeAddress) (*ResponseNormalizeAddress, error)
	mustEmbedUnimplementedNormalizeAddressServicesServerServer()
}

// UnimplementedNormalizeAddressServicesServerServer must be embedded to have forward compatible implementations.
type UnimplementedNormalizeAddressServicesServerServer struct {
}

func (UnimplementedNormalizeAddressServicesServerServer) NormalizeAddress(context.Context, *RequestNormalizeAddress) (*ResponseNormalizeAddress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NormalizeAddress not implemented")
}
func (UnimplementedNormalizeAddressServicesServerServer) mustEmbedUnimplementedNormalizeAddressServicesServerServer() {
}

// UnsafeNormalizeAddressServicesServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NormalizeAddressServicesServerServer will
// result in compilation errors.
type UnsafeNormalizeAddressServicesServerServer interface {
	mustEmbedUnimplementedNormalizeAddressServicesServerServer()
}

func RegisterNormalizeAddressServicesServerServer(s grpc.ServiceRegistrar, srv NormalizeAddressServicesServerServer) {
	s.RegisterService(&NormalizeAddressServicesServer_ServiceDesc, srv)
}

func _NormalizeAddressServicesServer_NormalizeAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestNormalizeAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NormalizeAddressServicesServerServer).NormalizeAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/normalize_address.NormalizeAddressServicesServer/NormalizeAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NormalizeAddressServicesServerServer).NormalizeAddress(ctx, req.(*RequestNormalizeAddress))
	}
	return interceptor(ctx, in, info, handler)
}

// NormalizeAddressServicesServer_ServiceDesc is the grpc.ServiceDesc for NormalizeAddressServicesServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NormalizeAddressServicesServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "normalize_address.NormalizeAddressServicesServer",
	HandlerType: (*NormalizeAddressServicesServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NormalizeAddress",
			Handler:    _NormalizeAddressServicesServer_NormalizeAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "normalize_address.proto",
}
