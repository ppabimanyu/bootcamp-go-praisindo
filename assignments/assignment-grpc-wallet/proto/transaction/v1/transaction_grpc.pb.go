// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: proto/transaction/v1/transaction.proto

package transaction

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
	Service_CreditTransaction_FullMethodName   = "/proto.transaction.v1.Service/CreditTransaction"
	Service_TransferTransaction_FullMethodName = "/proto.transaction.v1.Service/TransferTransaction"
	Service_DetailTransaction_FullMethodName   = "/proto.transaction.v1.Service/DetailTransaction"
	Service_GetTransaction_FullMethodName      = "/proto.transaction.v1.Service/GetTransaction"
)

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	CreditTransaction(ctx context.Context, in *CreditTransactionRequest, opts ...grpc.CallOption) (*CreditTransactionResponse, error)
	TransferTransaction(ctx context.Context, in *TransferTransactionRequest, opts ...grpc.CallOption) (*TransferTransactionResponse, error)
	DetailTransaction(ctx context.Context, in *DetailTransactionRequest, opts ...grpc.CallOption) (*DetailTransactionResponse, error)
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) CreditTransaction(ctx context.Context, in *CreditTransactionRequest, opts ...grpc.CallOption) (*CreditTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreditTransactionResponse)
	err := c.cc.Invoke(ctx, Service_CreditTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) TransferTransaction(ctx context.Context, in *TransferTransactionRequest, opts ...grpc.CallOption) (*TransferTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TransferTransactionResponse)
	err := c.cc.Invoke(ctx, Service_TransferTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DetailTransaction(ctx context.Context, in *DetailTransactionRequest, opts ...grpc.CallOption) (*DetailTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DetailTransactionResponse)
	err := c.cc.Invoke(ctx, Service_DetailTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTransactionResponse)
	err := c.cc.Invoke(ctx, Service_GetTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	CreditTransaction(context.Context, *CreditTransactionRequest) (*CreditTransactionResponse, error)
	TransferTransaction(context.Context, *TransferTransactionRequest) (*TransferTransactionResponse, error)
	DetailTransaction(context.Context, *DetailTransactionRequest) (*DetailTransactionResponse, error)
	GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) CreditTransaction(context.Context, *CreditTransactionRequest) (*CreditTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreditTransaction not implemented")
}
func (UnimplementedServiceServer) TransferTransaction(context.Context, *TransferTransactionRequest) (*TransferTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferTransaction not implemented")
}
func (UnimplementedServiceServer) DetailTransaction(context.Context, *DetailTransactionRequest) (*DetailTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetailTransaction not implemented")
}
func (UnimplementedServiceServer) GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_CreditTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreditTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).CreditTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_CreditTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).CreditTransaction(ctx, req.(*CreditTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_TransferTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).TransferTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_TransferTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).TransferTransaction(ctx, req.(*TransferTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DetailTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DetailTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_DetailTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DetailTransaction(ctx, req.(*DetailTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Service_GetTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.transaction.v1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreditTransaction",
			Handler:    _Service_CreditTransaction_Handler,
		},
		{
			MethodName: "TransferTransaction",
			Handler:    _Service_TransferTransaction_Handler,
		},
		{
			MethodName: "DetailTransaction",
			Handler:    _Service_DetailTransaction_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _Service_GetTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/transaction/v1/transaction.proto",
}
