package post

import (
	"context"
	"database/sql"
	"iman/internal/repositories/post"
	"iman/pkg/proto/post_service"

	"google.golang.org/protobuf/types/known/emptypb"
)

// compile time check
var _ = (post_service.PostServiceServer)(&Service{})

type Service struct {
	post_service.UnimplementedPostServiceServer
	PostsRepository post.Repository
}

func New(db *sql.DB) *Service {
	return &Service{
		PostsRepository: post.Repository{
			DB: db,
		},
	}
}

func (s *Service) GetPosts(ctx context.Context, in *post_service.GetPostsRequest) (*post_service.GetPostsResponse, error) {
	p, err := s.PostsRepository.GetPosts(ctx, int(in.Limit), int(in.Page-1)*int(in.Limit))
	if err != nil {
		return nil, err
	}

	return &post_service.GetPostsResponse{
		Posts: p,
	}, nil
}

func (s *Service) GetPostByID(ctx context.Context, in *post_service.GetPostByIDRequest) (*post_service.GetPostByIDResponse, error) {
	p, err := s.PostsRepository.GetPostByID(ctx, int(in.Id))
	if err != nil {
		return nil, err
	}

	return &post_service.GetPostByIDResponse{
		Post: p,
	}, nil
}

func (s *Service) DeletePostByID(ctx context.Context, in *post_service.DeletePostByIDRequest) (*emptypb.Empty, error) {
	err := s.PostsRepository.DeletePostByID(ctx, int(in.Id))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *Service) UpdatePostByID(ctx context.Context, in *post_service.UpdatePostByIDRequest) (*emptypb.Empty, error) {
	err := s.PostsRepository.UpdatePostByID(ctx, int(in.Post.Id), in.Post.Title, in.Post.Body)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}