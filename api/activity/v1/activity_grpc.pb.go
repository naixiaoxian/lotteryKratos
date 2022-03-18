// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0--rc1
// source: api/activity/v1/activity.proto

package v1

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

// ActivityClient is the client API for Activity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActivityClient interface {
	CreateActivity(ctx context.Context, in *CreateActivityRequest, opts ...grpc.CallOption) (*CreateActivityReply, error)
	UpdateActivity(ctx context.Context, in *UpdateActivityRequest, opts ...grpc.CallOption) (*UpdateActivityReply, error)
	DeleteActivity(ctx context.Context, in *DeleteActivityRequest, opts ...grpc.CallOption) (*DeleteActivityReply, error)
	GetActivity(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*GetActivityReply, error)
	ListActivity(ctx context.Context, in *ListActivityRequest, opts ...grpc.CallOption) (*ListActivityReply, error)
}

type activityClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityClient(cc grpc.ClientConnInterface) ActivityClient {
	return &activityClient{cc}
}

func (c *activityClient) CreateActivity(ctx context.Context, in *CreateActivityRequest, opts ...grpc.CallOption) (*CreateActivityReply, error) {
	out := new(CreateActivityReply)
	err := c.cc.Invoke(ctx, "/api.activity.v1.activity/CreateActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) UpdateActivity(ctx context.Context, in *UpdateActivityRequest, opts ...grpc.CallOption) (*UpdateActivityReply, error) {
	out := new(UpdateActivityReply)
	err := c.cc.Invoke(ctx, "/api.activity.v1.activity/UpdateActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) DeleteActivity(ctx context.Context, in *DeleteActivityRequest, opts ...grpc.CallOption) (*DeleteActivityReply, error) {
	out := new(DeleteActivityReply)
	err := c.cc.Invoke(ctx, "/api.activity.v1.activity/DeleteActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) GetActivity(ctx context.Context, in *GetActivityRequest, opts ...grpc.CallOption) (*GetActivityReply, error) {
	out := new(GetActivityReply)
	err := c.cc.Invoke(ctx, "/api.activity.v1.activity/GetActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityClient) ListActivity(ctx context.Context, in *ListActivityRequest, opts ...grpc.CallOption) (*ListActivityReply, error) {
	out := new(ListActivityReply)
	err := c.cc.Invoke(ctx, "/api.activity.v1.activity/ListActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityServer is the server API for Activity service.
// All implementations must embed UnimplementedActivityServer
// for forward compatibility
type ActivityServer interface {
	CreateActivity(context.Context, *CreateActivityRequest) (*CreateActivityReply, error)
	UpdateActivity(context.Context, *UpdateActivityRequest) (*UpdateActivityReply, error)
	DeleteActivity(context.Context, *DeleteActivityRequest) (*DeleteActivityReply, error)
	GetActivity(context.Context, *GetActivityRequest) (*GetActivityReply, error)
	ListActivity(context.Context, *ListActivityRequest) (*ListActivityReply, error)
	mustEmbedUnimplementedActivityServer()
}

// UnimplementedActivityServer must be embedded to have forward compatible implementations.
type UnimplementedActivityServer struct {
}

func (UnimplementedActivityServer) CreateActivity(context.Context, *CreateActivityRequest) (*CreateActivityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateActivity not implemented")
}
func (UnimplementedActivityServer) UpdateActivity(context.Context, *UpdateActivityRequest) (*UpdateActivityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateActivity not implemented")
}
func (UnimplementedActivityServer) DeleteActivity(context.Context, *DeleteActivityRequest) (*DeleteActivityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteActivity not implemented")
}
func (UnimplementedActivityServer) GetActivity(context.Context, *GetActivityRequest) (*GetActivityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivity not implemented")
}
func (UnimplementedActivityServer) ListActivity(context.Context, *ListActivityRequest) (*ListActivityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListActivity not implemented")
}
func (UnimplementedActivityServer) mustEmbedUnimplementedActivityServer() {}

// UnsafeActivityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActivityServer will
// result in compilation errors.
type UnsafeActivityServer interface {
	mustEmbedUnimplementedActivityServer()
}

func RegisterActivityServer(s grpc.ServiceRegistrar, srv ActivityServer) {
	s.RegisterService(&Activity_ServiceDesc, srv)
}

func _Activity_CreateActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).CreateActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.activity.v1.activity/CreateActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).CreateActivity(ctx, req.(*CreateActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_UpdateActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).UpdateActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.activity.v1.activity/UpdateActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).UpdateActivity(ctx, req.(*UpdateActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_DeleteActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).DeleteActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.activity.v1.activity/DeleteActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).DeleteActivity(ctx, req.(*DeleteActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_GetActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).GetActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.activity.v1.activity/GetActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).GetActivity(ctx, req.(*GetActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Activity_ListActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServer).ListActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.activity.v1.activity/ListActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServer).ListActivity(ctx, req.(*ListActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Activity_ServiceDesc is the grpc.ServiceDesc for Activity service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Activity_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.activity.v1.activity",
	HandlerType: (*ActivityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateActivity",
			Handler:    _Activity_CreateActivity_Handler,
		},
		{
			MethodName: "UpdateActivity",
			Handler:    _Activity_UpdateActivity_Handler,
		},
		{
			MethodName: "DeleteActivity",
			Handler:    _Activity_DeleteActivity_Handler,
		},
		{
			MethodName: "GetActivity",
			Handler:    _Activity_GetActivity_Handler,
		},
		{
			MethodName: "ListActivity",
			Handler:    _Activity_ListActivity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/activity/v1/activity.proto",
}
