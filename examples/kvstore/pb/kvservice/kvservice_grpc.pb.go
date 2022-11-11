// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: pb/proto/kvservice.proto

package kvservice

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

// KVServiceClient is the client API for KVService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KVServiceClient interface {
	Set(ctx context.Context, in *SetReq, opts ...grpc.CallOption) (*SetResp, error)
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error)
}

type kVServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKVServiceClient(cc grpc.ClientConnInterface) KVServiceClient {
	return &kVServiceClient{cc}
}

func (c *kVServiceClient) Set(ctx context.Context, in *SetReq, opts ...grpc.CallOption) (*SetResp, error) {
	out := new(SetResp)
	err := c.cc.Invoke(ctx, "/kvservice.KVService/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVServiceClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error) {
	out := new(GetResp)
	err := c.cc.Invoke(ctx, "/kvservice.KVService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KVServiceServer is the server API for KVService service.
// All implementations must embed UnimplementedKVServiceServer
// for forward compatibility
type KVServiceServer interface {
	Set(context.Context, *SetReq) (*SetResp, error)
	Get(context.Context, *GetReq) (*GetResp, error)
	mustEmbedUnimplementedKVServiceServer()
}

// UnimplementedKVServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKVServiceServer struct {
}

func (UnimplementedKVServiceServer) Set(context.Context, *SetReq) (*SetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedKVServiceServer) Get(context.Context, *GetReq) (*GetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedKVServiceServer) mustEmbedUnimplementedKVServiceServer() {}

// UnsafeKVServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KVServiceServer will
// result in compilation errors.
type UnsafeKVServiceServer interface {
	mustEmbedUnimplementedKVServiceServer()
}

func RegisterKVServiceServer(s grpc.ServiceRegistrar, srv KVServiceServer) {
	s.RegisterService(&KVService_ServiceDesc, srv)
}

func _KVService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvservice.KVService/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServiceServer).Set(ctx, req.(*SetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kvservice.KVService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServiceServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

// KVService_ServiceDesc is the grpc.ServiceDesc for KVService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KVService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kvservice.KVService",
	HandlerType: (*KVServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _KVService_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _KVService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/proto/kvservice.proto",
}
