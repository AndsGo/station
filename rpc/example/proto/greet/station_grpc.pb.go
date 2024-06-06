// 声明 proto 语法版本，固定值

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: station.proto

// proto 包名

package greet

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
	Greet_SayHello_FullMethodName = "/station.Greet/SayHello"
)

// GreetClient is the client API for Greet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetClient interface {
	// 定义一个 SayHello 一元 rpc 方法，请求体和响应体必填。
	SayHello(ctx context.Context, in *SayHelloReq, opts ...grpc.CallOption) (*SayHelloResp, error)
}

type greetClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetClient(cc grpc.ClientConnInterface) GreetClient {
	return &greetClient{cc}
}

func (c *greetClient) SayHello(ctx context.Context, in *SayHelloReq, opts ...grpc.CallOption) (*SayHelloResp, error) {
	out := new(SayHelloResp)
	err := c.cc.Invoke(ctx, Greet_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetServer is the server API for Greet service.
// All implementations must embed UnimplementedGreetServer
// for forward compatibility
type GreetServer interface {
	// 定义一个 SayHello 一元 rpc 方法，请求体和响应体必填。
	SayHello(context.Context, *SayHelloReq) (*SayHelloResp, error)
	mustEmbedUnimplementedGreetServer()
}

// UnimplementedGreetServer must be embedded to have forward compatible implementations.
type UnimplementedGreetServer struct {
}

func (UnimplementedGreetServer) SayHello(context.Context, *SayHelloReq) (*SayHelloResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreetServer) mustEmbedUnimplementedGreetServer() {}

// UnsafeGreetServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServer will
// result in compilation errors.
type UnsafeGreetServer interface {
	mustEmbedUnimplementedGreetServer()
}

func RegisterGreetServer(s grpc.ServiceRegistrar, srv GreetServer) {
	s.RegisterService(&Greet_ServiceDesc, srv)
}

func _Greet_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greet_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServer).SayHello(ctx, req.(*SayHelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Greet_ServiceDesc is the grpc.ServiceDesc for Greet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "station.Greet",
	HandlerType: (*GreetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greet_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "station.proto",
}
