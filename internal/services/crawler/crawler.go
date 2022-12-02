package crawler

import (
	context "context"
	"database/sql"
	"encoding/json"
	"fmt"
	"iman/internal/repositories/post"
	"iman/pkg/proto/crawler_service"
	"iman/pkg/proto/post_service"
	"net/http"
	"net/url"
	"strconv"
)

type Service interface {
	Crawl(ctx context.Context, in *crawler_service.CrawlRequest) (*crawler_service.CrawlResponse, error)
	PopulateData(ctx context.Context) (err error)
	crawler_service.UnsafeCrawlerServiceServer
}

type service struct {
	PostRepository post.Repository
	crawler_service.UnimplementedCrawlerServiceServer
}

type Params struct {
	DB *sql.DB
}

func New(p Params) Service {
	return &service{
		PostRepository: post.New(post.Params{DB: p.DB}),
	}
}

// Crawl returns 10 posts from the page you specify in page field
func (s *service) Crawl(ctx context.Context, in *crawler_service.CrawlRequest) (*crawler_service.CrawlResponse, error) {
	u, err := url.Parse("https://gorest.co.in/public/v1/posts")
	if err != nil {
		return nil, err
	}

	v := url.Values{}
	v.Add("page", strconv.Itoa(int(in.Page)))

	u.RawQuery = v.Encode()

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	var resp struct {
		Data []*crawler_service.Data `json:"data"`
	}

	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return &crawler_service.CrawlResponse{
		Data: resp.Data,
	}, nil
}

func (s *service) PopulateData(ctx context.Context) (err error) {
	cnt, err := s.PostRepository.GetPostsCount(ctx)
	if err != nil {
		return
	}

	if cnt >= 50 {
		return nil
	}

	want := 50 - cnt

	inserted := 0
	page := 1
	for inserted < want {
		d, err := s.Crawl(ctx, &crawler_service.CrawlRequest{Page: int32(page)})
		if err != nil {
			return err
		}

		for _, v := range d.Data {
			err := s.PostRepository.CreatePost(ctx, &post_service.Post{Id: v.Id, UserId: v.UserId, Title: v.Title, Body: v.Body})
			if err != nil {
				continue
			}
			inserted++

			if inserted == want {
				break
			}
		}

		page++
	}

	return
}
