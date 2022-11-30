package handlers

import (
	"encoding/json"
	"iman/pkg/proto/post_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
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
	resp := &ApiResponse{}
	defer c.JSON(resp.Code, resp)

	limitString := c.Query("limit")
	pageString := c.Query("page")

	limit, err := strconv.Atoi(limitString)
	if err != nil {
		resp.Set(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "no limit provided")
		return
	}

	page, err := strconv.Atoi(pageString)
	if err != nil {
		resp.Set(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "no page provided")
		return
	}

	data, err := h.PostServiceClient.GetPosts(c.Request.Context(), &post_service.GetPostsRequest{Limit: int32(limit), Page: int32(page)})
	if err != nil {
		resp.Set(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err.Error())
		return
	}

	resp.Set(http.StatusOK, http.StatusText(http.StatusOK), data)
}

func (h *Handlers) GetPostByID(c *gin.Context) {
	resp := &ApiResponse{}
	defer c.JSON(resp.Code, resp)

	idString := c.Query("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		resp.Set(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "no id provided")
		return
	}

	post, err := h.PostServiceClient.GetPostByID(c, &post_service.GetPostByIDRequest{Id: int32(id)})
	if err != nil {
		resp.Set(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	resp.Set(http.StatusOK, http.StatusText(http.StatusOK), post)
}

func (h *Handlers) DeletePostByID(c *gin.Context) {
	resp := &ApiResponse{}
	defer c.JSON(resp.Code, resp)

	idString := c.Query("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		resp.Set(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "no id provided")
		return
	}

	_, err = h.PostServiceClient.DeletePostByID(c, &post_service.DeletePostByIDRequest{Id: int32(id)})
	if err != nil {
		resp.Set(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	resp.Set(http.StatusOK, http.StatusText(http.StatusOK), "post deleted")
}

func (h *Handlers) UpdatePostByID(c *gin.Context) {
	resp := &ApiResponse{}
	defer c.JSON(resp.Code, resp)

	var request post_service.Post

	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		resp.Set(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	if request.Id == 0 {
		resp.Set(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "no id provided")
		return
	}

	if request.Body == "" {
		resp.Set(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "no body provided")
		return
	}

	if request.Title == "" {
		resp.Set(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), "no title provided")
		return
	}

	_, err = h.PostServiceClient.UpdatePostByID(c, &post_service.UpdatePostByIDRequest{Post: &request})
	if err != nil {
		resp.Set(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	resp.Set(http.StatusOK, http.StatusText(http.StatusOK), "post updated")
}

type ApiResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Payload any    `json:"payload"`
}

func (r *ApiResponse) Set(code int, message string, payload any) {
	r.Code = code
	r.Message = message
	r.Payload = payload
}
