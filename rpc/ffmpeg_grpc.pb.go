// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: rpc/ffmpeg.proto

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

// FfmpegClient is the client API for Ffmpeg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FfmpegClient interface {
	PersonDetection(ctx context.Context, in *CameraIp, opts ...grpc.CallOption) (*FfmpegResponse, error)
	VideoStream(ctx context.Context, in *VideoStreamStruct, opts ...grpc.CallOption) (*VideoStreamResponse, error)
	ChangeFfmpegFlag(ctx context.Context, in *FlagParam, opts ...grpc.CallOption) (*ChangeFlagFfmpegResponse, error)
}

type ffmpegClient struct {
	cc grpc.ClientConnInterface
}

func NewFfmpegClient(cc grpc.ClientConnInterface) FfmpegClient {
	return &ffmpegClient{cc}
}

func (c *ffmpegClient) PersonDetection(ctx context.Context, in *CameraIp, opts ...grpc.CallOption) (*FfmpegResponse, error) {
	out := new(FfmpegResponse)
	err := c.cc.Invoke(ctx, "/Ffmpeg/PersonDetection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ffmpegClient) VideoStream(ctx context.Context, in *VideoStreamStruct, opts ...grpc.CallOption) (*VideoStreamResponse, error) {
	out := new(VideoStreamResponse)
	err := c.cc.Invoke(ctx, "/Ffmpeg/VideoStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ffmpegClient) ChangeFfmpegFlag(ctx context.Context, in *FlagParam, opts ...grpc.CallOption) (*ChangeFlagFfmpegResponse, error) {
	out := new(ChangeFlagFfmpegResponse)
	err := c.cc.Invoke(ctx, "/Ffmpeg/ChangeFfmpegFlag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FfmpegServer is the server API for Ffmpeg service.
// All implementations must embed UnimplementedFfmpegServer
// for forward compatibility
type FfmpegServer interface {
	PersonDetection(context.Context, *CameraIp) (*FfmpegResponse, error)
	VideoStream(context.Context, *VideoStreamStruct) (*VideoStreamResponse, error)
	ChangeFfmpegFlag(context.Context, *FlagParam) (*ChangeFlagFfmpegResponse, error)
	mustEmbedUnimplementedFfmpegServer()
}

// UnimplementedFfmpegServer must be embedded to have forward compatible implementations.
type UnimplementedFfmpegServer struct {
}

func (UnimplementedFfmpegServer) PersonDetection(context.Context, *CameraIp) (*FfmpegResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PersonDetection not implemented")
}
func (UnimplementedFfmpegServer) VideoStream(context.Context, *VideoStreamStruct) (*VideoStreamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VideoStream not implemented")
}
func (UnimplementedFfmpegServer) ChangeFfmpegFlag(context.Context, *FlagParam) (*ChangeFlagFfmpegResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeFfmpegFlag not implemented")
}
func (UnimplementedFfmpegServer) mustEmbedUnimplementedFfmpegServer() {}

// UnsafeFfmpegServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FfmpegServer will
// result in compilation errors.
type UnsafeFfmpegServer interface {
	mustEmbedUnimplementedFfmpegServer()
}

func RegisterFfmpegServer(s grpc.ServiceRegistrar, srv FfmpegServer) {
	s.RegisterService(&Ffmpeg_ServiceDesc, srv)
}

func _Ffmpeg_PersonDetection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CameraIp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FfmpegServer).PersonDetection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ffmpeg/PersonDetection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FfmpegServer).PersonDetection(ctx, req.(*CameraIp))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ffmpeg_VideoStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoStreamStruct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FfmpegServer).VideoStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ffmpeg/VideoStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FfmpegServer).VideoStream(ctx, req.(*VideoStreamStruct))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ffmpeg_ChangeFfmpegFlag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlagParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FfmpegServer).ChangeFfmpegFlag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ffmpeg/ChangeFfmpegFlag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FfmpegServer).ChangeFfmpegFlag(ctx, req.(*FlagParam))
	}
	return interceptor(ctx, in, info, handler)
}

// Ffmpeg_ServiceDesc is the grpc.ServiceDesc for Ffmpeg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ffmpeg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Ffmpeg",
	HandlerType: (*FfmpegServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PersonDetection",
			Handler:    _Ffmpeg_PersonDetection_Handler,
		},
		{
			MethodName: "VideoStream",
			Handler:    _Ffmpeg_VideoStream_Handler,
		},
		{
			MethodName: "ChangeFfmpegFlag",
			Handler:    _Ffmpeg_ChangeFfmpegFlag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/ffmpeg.proto",
}
