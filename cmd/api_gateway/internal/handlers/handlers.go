package handlers

import (
	"encoding/json"
	"iman/pkg/api"
	"iman/pkg/proto/post_service"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handlers struct {
	PostServiceClient post_service.PostServiceClient
}

func NewHandler(conn *grpc.ClientConn) *Handlers {
	return &Handlers{
		PostServiceClient: post_service.NewPostServiceClient(conn),
	}
}

func (h *Handlers) GetPosts(c *gin.Context) {
	resp := &api.Response{}
	defer c.JSON(resp.Code, resp)

	limitString := c.Query("limit")
	pageString := c.Query("page")

	limit, err := strconv.Atoi(limitString)
	if err != nil {
		resp.BadRequest("no limit provided")
		return
	}

	page, err := strconv.Atoi(pageString)
	if err != nil {
		resp.BadRequest("no page provided")
		return
	}

	data, err := h.PostServiceClient.GetPosts(c.Request.Context(), &post_service.GetPostsRequest{Limit: int32(limit), Page: int32(page)})
	if err != nil {
		s, ok := status.FromError(err)
		if !ok {
			resp.InternalServerError(err.Error())
			return
		}

		if s.Code() == codes.NotFound {
			resp.NotFound()
			return
		}

		resp.InternalServerError(err.Error())
		return
	}

	resp.Ok(data)
}

func (h *Handlers) GetPostByID(c *gin.Context) {
	resp := &api.Response{}
	defer c.JSON(resp.Code, resp)

	idString := c.Query("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		resp.BadRequest("no id provided")
		return
	}

	post, err := h.PostServiceClient.GetPostByID(c, &post_service.GetPostByIDRequest{Id: int32(id)})
	if err != nil {
		s, ok := status.FromError(err)
		if !ok {
			resp.InternalServerError(err.Error())
			return
		}

		if s.Code() == codes.NotFound {
			resp.NotFound()
			return
		}

		resp.InternalServerError(err.Error())
		return
	}

	resp.Ok(post)
}

func (h *Handlers) DeletePostByID(c *gin.Context) {
	resp := &api.Response{}
	defer c.JSON(resp.Code, resp)

	idString := c.Query("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		resp.BadRequest("no id provided")
		return
	}

	_, err = h.PostServiceClient.DeletePostByID(c, &post_service.DeletePostByIDRequest{Id: int32(id)})
	if err != nil {
		s, ok := status.FromError(err)
		if !ok {
			resp.InternalServerError(err.Error())
			return
		}

		if s.Code() == codes.NotFound {
			resp.NotFound()
			return
		}

		resp.InternalServerError(err.Error())
		return
	}

	resp.Ok("post deleted")
}

func (h *Handlers) UpdatePostByID(c *gin.Context) {
	resp := &api.Response{}
	defer c.JSON(resp.Code, resp)

	var request post_service.Post

	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		resp.BadRequest(err.Error())
		return
	}

	if request.Id == 0 {
		resp.BadRequest("no id provided")
		return
	}

	if request.Body == "" {
		resp.BadRequest("no body provided")
		return
	}

	if request.Title == "" {
		resp.BadRequest("no title provided")
		return
	}

	_, err = h.PostServiceClient.UpdatePostByID(c, &post_service.UpdatePostByIDRequest{Post: &request})
	if err != nil {
		s, ok := status.FromError(err)
		if !ok {
			resp.InternalServerError(err.Error())
			return
		}

		if s.Code() == codes.NotFound {
			resp.NotFound()
			return
		}

		resp.InternalServerError(err.Error())
		return
	}

	resp.Ok("post updated")
}
