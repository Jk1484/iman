// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: post_service.proto

package post_service

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId int32  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title  string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Body   string `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{0}
}

func (x *Post) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Post) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type GetPostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page  int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetPostsRequest) Reset() {
	*x = GetPostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostsRequest) ProtoMessage() {}

func (x *GetPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostsRequest.ProtoReflect.Descriptor instead.
func (*GetPostsRequest) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetPostsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetPostsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type GetPostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *GetPostsResponse) Reset() {
	*x = GetPostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostsResponse) ProtoMessage() {}

func (x *GetPostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostsResponse.ProtoReflect.Descriptor instead.
func (*GetPostsResponse) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetPostsResponse) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type GetPostByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPostByIDRequest) Reset() {
	*x = GetPostByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostByIDRequest) ProtoMessage() {}

func (x *GetPostByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostByIDRequest.ProtoReflect.Descriptor instead.
func (*GetPostByIDRequest) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetPostByIDRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetPostByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *GetPostByIDResponse) Reset() {
	*x = GetPostByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostByIDResponse) ProtoMessage() {}

func (x *GetPostByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostByIDResponse.ProtoReflect.Descriptor instead.
func (*GetPostByIDResponse) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetPostByIDResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type DeletePostByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePostByIDRequest) Reset() {
	*x = DeletePostByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostByIDRequest) ProtoMessage() {}

func (x *DeletePostByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostByIDRequest.ProtoReflect.Descriptor instead.
func (*DeletePostByIDRequest) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeletePostByIDRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdatePostByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *UpdatePostByIDRequest) Reset() {
	*x = UpdatePostByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostByIDRequest) ProtoMessage() {}

func (x *UpdatePostByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostByIDRequest.ProtoReflect.Descriptor instead.
func (*UpdatePostByIDRequest) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePostByIDRequest) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type UpdatePostByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *UpdatePostByIDResponse) Reset() {
	*x = UpdatePostByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_post_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostByIDResponse) ProtoMessage() {}

func (x *UpdatePostByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_post_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostByIDResponse.ProtoReflect.Descriptor instead.
func (*UpdatePostByIDResponse) Descriptor() ([]byte, []int) {
	return file_post_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdatePostByIDResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

var File_post_service_proto protoreflect.FileDescriptor

var file_post_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x59, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x3b, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x3c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x70,
	0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05,
	0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x04,
	0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04,
	0x70, 0x6f, 0x73, 0x74, 0x22, 0x40, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x32, 0xd2, 0x02, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x12, 0x20, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x23, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x23, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e,
	0x2f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_post_service_proto_rawDescOnce sync.Once
	file_post_service_proto_rawDescData = file_post_service_proto_rawDesc
)

func file_post_service_proto_rawDescGZIP() []byte {
	file_post_service_proto_rawDescOnce.Do(func() {
		file_post_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_post_service_proto_rawDescData)
	})
	return file_post_service_proto_rawDescData
}

var file_post_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_post_service_proto_goTypes = []interface{}{
	(*Post)(nil),                   // 0: post_service.Post
	(*GetPostsRequest)(nil),        // 1: post_service.GetPostsRequest
	(*GetPostsResponse)(nil),       // 2: post_service.GetPostsResponse
	(*GetPostByIDRequest)(nil),     // 3: post_service.GetPostByIDRequest
	(*GetPostByIDResponse)(nil),    // 4: post_service.GetPostByIDResponse
	(*DeletePostByIDRequest)(nil),  // 5: post_service.DeletePostByIDRequest
	(*UpdatePostByIDRequest)(nil),  // 6: post_service.UpdatePostByIDRequest
	(*UpdatePostByIDResponse)(nil), // 7: post_service.UpdatePostByIDResponse
	(*empty.Empty)(nil),            // 8: google.protobuf.Empty
}
var file_post_service_proto_depIdxs = []int32{
	0, // 0: post_service.GetPostsResponse.posts:type_name -> post_service.Post
	0, // 1: post_service.GetPostByIDResponse.post:type_name -> post_service.Post
	0, // 2: post_service.UpdatePostByIDRequest.post:type_name -> post_service.Post
	0, // 3: post_service.UpdatePostByIDResponse.post:type_name -> post_service.Post
	1, // 4: post_service.PostService.GetPosts:input_type -> post_service.GetPostsRequest
	3, // 5: post_service.PostService.GetPostByID:input_type -> post_service.GetPostByIDRequest
	5, // 6: post_service.PostService.DeletePostByID:input_type -> post_service.DeletePostByIDRequest
	6, // 7: post_service.PostService.UpdatePostByID:input_type -> post_service.UpdatePostByIDRequest
	2, // 8: post_service.PostService.GetPosts:output_type -> post_service.GetPostsResponse
	4, // 9: post_service.PostService.GetPostByID:output_type -> post_service.GetPostByIDResponse
	8, // 10: post_service.PostService.DeletePostByID:output_type -> google.protobuf.Empty
	8, // 11: post_service.PostService.UpdatePostByID:output_type -> google.protobuf.Empty
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_post_service_proto_init() }
func file_post_service_proto_init() {
	if File_post_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_post_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostByIDRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostByIDResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostByIDRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostByIDRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_post_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostByIDResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_post_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_post_service_proto_goTypes,
		DependencyIndexes: file_post_service_proto_depIdxs,
		MessageInfos:      file_post_service_proto_msgTypes,
	}.Build()
	File_post_service_proto = out.File
	file_post_service_proto_rawDesc = nil
	file_post_service_proto_goTypes = nil
	file_post_service_proto_depIdxs = nil
}