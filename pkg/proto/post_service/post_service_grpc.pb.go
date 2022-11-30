// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: post_service.proto

package post_service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostServiceClient interface {
	GetPosts(ctx context.Context, in *GetPostsRequest, opts ...grpc.CallOption) (*GetPostsResponse, error)
	GetPostByID(ctx context.Context, in *GetPostByIDRequest, opts ...grpc.CallOption) (*GetPostByIDResponse, error)
	DeletePostByID(ctx context.Context, in *DeletePostByIDRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdatePostByID(ctx context.Context, in *UpdatePostByIDRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type postServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostServiceClient(cc grpc.ClientConnInterface) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) GetPosts(ctx context.Context, in *GetPostsRequest, opts ...grpc.CallOption) (*GetPostsResponse, error) {
	out := new(GetPostsResponse)
	err := c.cc.Invoke(ctx, "/post_service.PostService/GetPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetPostByID(ctx context.Context, in *GetPostByIDRequest, opts ...grpc.CallOption) (*GetPostByIDResponse, error) {
	out := new(GetPostByIDResponse)
	err := c.cc.Invoke(ctx, "/post_service.PostService/GetPostByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeletePostByID(ctx context.Context, in *DeletePostByIDRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/post_service.PostService/DeletePostByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) UpdatePostByID(ctx context.Context, in *UpdatePostByIDRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/post_service.PostService/UpdatePostByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
// All implementations must embed UnimplementedPostServiceServer
// for forward compatibility
type PostServiceServer interface {
	GetPosts(context.Context, *GetPostsRequest) (*GetPostsResponse, error)
	GetPostByID(context.Context, *GetPostByIDRequest) (*GetPostByIDResponse, error)
	DeletePostByID(context.Context, *DeletePostByIDRequest) (*empty.Empty, error)
	UpdatePostByID(context.Context, *UpdatePostByIDRequest) (*empty.Empty, error)
	mustEmbedUnimplementedPostServiceServer()
}

// UnimplementedPostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostServiceServer struct {
}

func (UnimplementedPostServiceServer) GetPosts(context.Context, *GetPostsRequest) (*GetPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPosts not implemented")
}
func (UnimplementedPostServiceServer) GetPostByID(context.Context, *GetPostByIDRequest) (*GetPostByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPostByID not implemented")
}
func (UnimplementedPostServiceServer) DeletePostByID(context.Context, *DeletePostByIDRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePostByID not implemented")
}
func (UnimplementedPostServiceServer) UpdatePostByID(context.Context, *UpdatePostByIDRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePostByID not implemented")
}
func (UnimplementedPostServiceServer) mustEmbedUnimplementedPostServiceServer() {}

// UnsafePostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServiceServer will
// result in compilation errors.
type UnsafePostServiceServer interface {
	mustEmbedUnimplementedPostServiceServer()
}

func RegisterPostServiceServer(s grpc.ServiceRegistrar, srv PostServiceServer) {
	s.RegisterService(&PostService_ServiceDesc, srv)
}

func _PostService_GetPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post_service.PostService/GetPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPosts(ctx, req.(*GetPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetPostByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetPostByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post_service.PostService/GetPostByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetPostByID(ctx, req.(*GetPostByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeletePostByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeletePostByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post_service.PostService/DeletePostByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeletePostByID(ctx, req.(*DeletePostByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_UpdatePostByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).UpdatePostByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post_service.PostService/UpdatePostByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).UpdatePostByID(ctx, req.(*UpdatePostByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostService_ServiceDesc is the grpc.ServiceDesc for PostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "post_service.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPosts",
			Handler:    _PostService_GetPosts_Handler,
		},
		{
			MethodName: "GetPostByID",
			Handler:    _PostService_GetPostByID_Handler,
		},
		{
			MethodName: "DeletePostByID",
			Handler:    _PostService_DeletePostByID_Handler,
		},
		{
			MethodName: "UpdatePostByID",
			Handler:    _PostService_UpdatePostByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post_service.proto",
}
