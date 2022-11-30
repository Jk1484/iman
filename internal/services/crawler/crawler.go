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

type Crawler struct {
	crawler_service.UnimplementedCrawlerServiceServer
	PostRepository post.Repository
}

func New(conn *sql.DB) *Crawler {
	return &Crawler{
		PostRepository: post.Repository{
			DB: conn,
		},
	}
}

// Crawl returns 10 posts from the page you specify in page field
func (c *Crawler) Crawl(ctx context.Context, in *crawler_service.CrawlRequest) (*crawler_service.CrawlResponse, error) {
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

func (c *Crawler) PopulateData(ctx context.Context) (err error) {
	cnt, err := c.PostRepository.GetPostsCount(ctx)
	if err != nil {
		return
	}

	if cnt >= 50 {
		return nil
	}

	want := 50 - cnt
	pages := want / 10
	if want%10 > 0 {
		pages++
	}

	var data []*crawler_service.Data

	for i := 1; i <= pages; i++ {
		d, err := c.Crawl(ctx, &crawler_service.CrawlRequest{Page: int32(i)})
		if err != nil {
			return err
		}

		data = append(data, d.Data...)
	}

	data = data[:want]

	for _, v := range data {
		err = c.PostRepository.CreatePost(ctx, &post_service.Post{Id: v.Id, UserId: v.UserId, Title: v.Title, Body: v.Body})
		if err != nil {
			return
		}
	}

	return
}