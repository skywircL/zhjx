// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: rpc/superresolution.proto

package rpc

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

// SuperResolutionClient is the client API for SuperResolution service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SuperResolutionClient interface {
	SuperResolutionFunc(ctx context.Context, in *SuperResolutionRequest, opts ...grpc.CallOption) (*SuperResolutionResponse, error)
	PersonBank(ctx context.Context, in *PersonBankRequest, opts ...grpc.CallOption) (*PersonBankResponse, error)
}

type superResolutionClient struct {
	cc grpc.ClientConnInterface
}

func NewSuperResolutionClient(cc grpc.ClientConnInterface) SuperResolutionClient {
	return &superResolutionClient{cc}
}

func (c *superResolutionClient) SuperResolutionFunc(ctx context.Context, in *SuperResolutionRequest, opts ...grpc.CallOption) (*SuperResolutionResponse, error) {
	out := new(SuperResolutionResponse)
	err := c.cc.Invoke(ctx, "/SuperResolution/SuperResolutionFunc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *superResolutionClient) PersonBank(ctx context.Context, in *PersonBankRequest, opts ...grpc.CallOption) (*PersonBankResponse, error) {
	out := new(PersonBankResponse)
	err := c.cc.Invoke(ctx, "/SuperResolution/PersonBank", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SuperResolutionServer is the server API for SuperResolution service.
// All implementations must embed UnimplementedSuperResolutionServer
// for forward compatibility
type SuperResolutionServer interface {
	SuperResolutionFunc(context.Context, *SuperResolutionRequest) (*SuperResolutionResponse, error)
	PersonBank(context.Context, *PersonBankRequest) (*PersonBankResponse, error)
	mustEmbedUnimplementedSuperResolutionServer()
}

// UnimplementedSuperResolutionServer must be embedded to have forward compatible implementations.
type UnimplementedSuperResolutionServer struct {
}

func (UnimplementedSuperResolutionServer) SuperResolutionFunc(context.Context, *SuperResolutionRequest) (*SuperResolutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SuperResolutionFunc not implemented")
}
func (UnimplementedSuperResolutionServer) PersonBank(context.Context, *PersonBankRequest) (*PersonBankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PersonBank not implemented")
}
func (UnimplementedSuperResolutionServer) mustEmbedUnimplementedSuperResolutionServer() {}

// UnsafeSuperResolutionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SuperResolutionServer will
// result in compilation errors.
type UnsafeSuperResolutionServer interface {
	mustEmbedUnimplementedSuperResolutionServer()
}

func RegisterSuperResolutionServer(s grpc.ServiceRegistrar, srv SuperResolutionServer) {
	s.RegisterService(&SuperResolution_ServiceDesc, srv)
}

func _SuperResolution_SuperResolutionFunc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuperResolutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuperResolutionServer).SuperResolutionFunc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SuperResolution/SuperResolutionFunc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuperResolutionServer).SuperResolutionFunc(ctx, req.(*SuperResolutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SuperResolution_PersonBank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonBankRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuperResolutionServer).PersonBank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SuperResolution/PersonBank",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuperResolutionServer).PersonBank(ctx, req.(*PersonBankRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SuperResolution_ServiceDesc is the grpc.ServiceDesc for SuperResolution service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SuperResolution_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SuperResolution",
	HandlerType: (*SuperResolutionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SuperResolutionFunc",
			Handler:    _SuperResolution_SuperResolutionFunc_Handler,
		},
		{
			MethodName: "PersonBank",
			Handler:    _SuperResolution_PersonBank_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/superresolution.proto",
}
