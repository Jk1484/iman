package post

import (
	"context"
	"database/sql"
	"iman/internal/repositories/post"
	"iman/pkg/proto/post_service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// compile time check
var _ = (post_service.PostServiceServer)(&service{})

type Service interface {
	GetPosts(ctx context.Context, in *post_service.GetPostsRequest) (*post_service.GetPostsResponse, error)
	GetPostByID(ctx context.Context, in *post_service.GetPostByIDRequest) (*post_service.GetPostByIDResponse, error)
	DeletePostByID(ctx context.Context, in *post_service.DeletePostByIDRequest) (*emptypb.Empty, error)
	UpdatePostByID(ctx context.Context, in *post_service.UpdatePostByIDRequest) (*emptypb.Empty, error)
	post_service.UnsafePostServiceServer
}

type service struct {
	PostsRepository post.Repository
	post_service.UnimplementedPostServiceServer
}

type Params struct {
	DB *sql.DB
}

func New(p Params) Service {
	return &service{
		PostsRepository: post.New(post.Params{DB: p.DB}),
	}
}

func (s *service) GetPosts(ctx context.Context, in *post_service.GetPostsRequest) (*post_service.GetPostsResponse, error) {
	p, err := s.PostsRepository.GetPosts(ctx, int(in.Limit), int(in.Page-1)*int(in.Limit))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, err
	}

	return &post_service.GetPostsResponse{
		Posts: p,
	}, nil
}

func (s *service) GetPostByID(ctx context.Context, in *post_service.GetPostByIDRequest) (*post_service.GetPostByIDResponse, error) {
	p, err := s.PostsRepository.GetPostByID(ctx, int(in.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			if err == sql.ErrNoRows {
				return nil, status.Error(codes.NotFound, "not found")
			}
		}

		return nil, err
	}

	return &post_service.GetPostByIDResponse{
		Post: p,
	}, nil
}

func (s *service) DeletePostByID(ctx context.Context, in *post_service.DeletePostByIDRequest) (*emptypb.Empty, error) {
	err := s.PostsRepository.DeletePostByID(ctx, int(in.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *service) UpdatePostByID(ctx context.Context, in *post_service.UpdatePostByIDRequest) (*emptypb.Empty, error) {
	err := s.PostsRepository.UpdatePostByID(ctx, int(in.Post.Id), in.Post.Title, in.Post.Body)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, err
	}

	return &emptypb.Empty{}, nil
}
