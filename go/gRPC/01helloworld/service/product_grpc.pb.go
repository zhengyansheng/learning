// 指定当前proto语法的版本 有2和3

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: product.proto

// 指定文件生成的package

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ProdService_GetProductStock_FullMethodName = "/service.ProdService/GetProductStock"
)

// ProdServiceClient is the client API for ProdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 定义服务主体
type ProdServiceClient interface {
	// 定义方法
	GetProductStock(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
}

type prodServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProdServiceClient(cc grpc.ClientConnInterface) ProdServiceClient {
	return &prodServiceClient{cc}
}

func (c *prodServiceClient) GetProductStock(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, ProdService_GetProductStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProdServiceServer is the server API for ProdService service.
// All implementations must embed UnimplementedProdServiceServer
// for forward compatibility.
//
// 定义服务主体
type ProdServiceServer interface {
	// 定义方法
	GetProductStock(context.Context, *ProductRequest) (*ProductResponse, error)
	mustEmbedUnimplementedProdServiceServer()
}

// UnimplementedProdServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProdServiceServer struct{}

func (UnimplementedProdServiceServer) GetProductStock(context.Context, *ProductRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductStock not implemented")
}
func (UnimplementedProdServiceServer) mustEmbedUnimplementedProdServiceServer() {}
func (UnimplementedProdServiceServer) testEmbeddedByValue()                     {}

// UnsafeProdServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProdServiceServer will
// result in compilation errors.
type UnsafeProdServiceServer interface {
	mustEmbedUnimplementedProdServiceServer()
}

func RegisterProdServiceServer(s grpc.ServiceRegistrar, srv ProdServiceServer) {
	// If the following call pancis, it indicates UnimplementedProdServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProdService_ServiceDesc, srv)
}

func _ProdService_GetProductStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdServiceServer).GetProductStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProdService_GetProductStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdServiceServer).GetProductStock(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProdService_ServiceDesc is the grpc.ServiceDesc for ProdService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProdService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.ProdService",
	HandlerType: (*ProdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProductStock",
			Handler:    _ProdService_GetProductStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
