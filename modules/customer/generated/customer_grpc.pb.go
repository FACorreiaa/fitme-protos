// Note: This is a proto3 file, proto2 is deprecated and should not be used.
// In proto3, all fields are optional by default, required and optional keywords
// are no longer required. Everything is optional by default.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: customer.proto

// Declaring a package name prevents collisions with other packages that use
// methods with the same name.

package generated

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
	Customer_GetCustomer_FullMethodName    = "/fitSphere.customer.Customer/GetCustomer"
	Customer_CreateCustomer_FullMethodName = "/fitSphere.customer.Customer/CreateCustomer"
	Customer_UpdateCustomer_FullMethodName = "/fitSphere.customer.Customer/UpdateCustomer"
	Customer_DeleteCustomer_FullMethodName = "/fitSphere.customer.Customer/DeleteCustomer"
)

// CustomerClient is the client API for Customer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// While Get and Create implement the same response, it's good practice to
// always return a unique response type for each method to save yourself
// headaches in the future.
type CustomerClient interface {
	GetCustomer(ctx context.Context, in *GetCustomerReq, opts ...grpc.CallOption) (*GetCustomerRes, error)
	CreateCustomer(ctx context.Context, in *CreateCustomerReq, opts ...grpc.CallOption) (*CreateCustomerRes, error)
	UpdateCustomer(ctx context.Context, in *UpdateCustomerReq, opts ...grpc.CallOption) (*UpdateCustomerRes, error)
	DeleteCustomer(ctx context.Context, in *DeleteCustomerReq, opts ...grpc.CallOption) (*NilRes, error)
}

type customerClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerClient(cc grpc.ClientConnInterface) CustomerClient {
	return &customerClient{cc}
}

func (c *customerClient) GetCustomer(ctx context.Context, in *GetCustomerReq, opts ...grpc.CallOption) (*GetCustomerRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCustomerRes)
	err := c.cc.Invoke(ctx, Customer_GetCustomer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) CreateCustomer(ctx context.Context, in *CreateCustomerReq, opts ...grpc.CallOption) (*CreateCustomerRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCustomerRes)
	err := c.cc.Invoke(ctx, Customer_CreateCustomer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) UpdateCustomer(ctx context.Context, in *UpdateCustomerReq, opts ...grpc.CallOption) (*UpdateCustomerRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCustomerRes)
	err := c.cc.Invoke(ctx, Customer_UpdateCustomer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerClient) DeleteCustomer(ctx context.Context, in *DeleteCustomerReq, opts ...grpc.CallOption) (*NilRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NilRes)
	err := c.cc.Invoke(ctx, Customer_DeleteCustomer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServer is the server API for Customer service.
// All implementations must embed UnimplementedCustomerServer
// for forward compatibility.
//
// While Get and Create implement the same response, it's good practice to
// always return a unique response type for each method to save yourself
// headaches in the future.
type CustomerServer interface {
	GetCustomer(context.Context, *GetCustomerReq) (*GetCustomerRes, error)
	CreateCustomer(context.Context, *CreateCustomerReq) (*CreateCustomerRes, error)
	UpdateCustomer(context.Context, *UpdateCustomerReq) (*UpdateCustomerRes, error)
	DeleteCustomer(context.Context, *DeleteCustomerReq) (*NilRes, error)
	mustEmbedUnimplementedCustomerServer()
}

// UnimplementedCustomerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCustomerServer struct{}

func (UnimplementedCustomerServer) GetCustomer(context.Context, *GetCustomerReq) (*GetCustomerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomer not implemented")
}
func (UnimplementedCustomerServer) CreateCustomer(context.Context, *CreateCustomerReq) (*CreateCustomerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomer not implemented")
}
func (UnimplementedCustomerServer) UpdateCustomer(context.Context, *UpdateCustomerReq) (*UpdateCustomerRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomer not implemented")
}
func (UnimplementedCustomerServer) DeleteCustomer(context.Context, *DeleteCustomerReq) (*NilRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCustomer not implemented")
}
func (UnimplementedCustomerServer) mustEmbedUnimplementedCustomerServer() {}
func (UnimplementedCustomerServer) testEmbeddedByValue()                  {}

// UnsafeCustomerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerServer will
// result in compilation errors.
type UnsafeCustomerServer interface {
	mustEmbedUnimplementedCustomerServer()
}

func RegisterCustomerServer(s grpc.ServiceRegistrar, srv CustomerServer) {
	// If the following call pancis, it indicates UnimplementedCustomerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Customer_ServiceDesc, srv)
}

func _Customer_GetCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).GetCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Customer_GetCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).GetCustomer(ctx, req.(*GetCustomerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustomerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Customer_CreateCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).CreateCustomer(ctx, req.(*CreateCustomerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_UpdateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCustomerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).UpdateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Customer_UpdateCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).UpdateCustomer(ctx, req.(*UpdateCustomerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Customer_DeleteCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCustomerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServer).DeleteCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Customer_DeleteCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServer).DeleteCustomer(ctx, req.(*DeleteCustomerReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Customer_ServiceDesc is the grpc.ServiceDesc for Customer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Customer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fitSphere.customer.Customer",
	HandlerType: (*CustomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomer",
			Handler:    _Customer_GetCustomer_Handler,
		},
		{
			MethodName: "CreateCustomer",
			Handler:    _Customer_CreateCustomer_Handler,
		},
		{
			MethodName: "UpdateCustomer",
			Handler:    _Customer_UpdateCustomer_Handler,
		},
		{
			MethodName: "DeleteCustomer",
			Handler:    _Customer_DeleteCustomer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customer.proto",
}
