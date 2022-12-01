package api

import (
	"errors"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Payload any    `json:"payload"`
}

var ErrNotFound = errors.New("not found")

func (r *Response) Ok(payload any) {
	r.Code = http.StatusOK
	r.Message = http.StatusText(http.StatusOK)
	r.Payload = payload
}

func (r *Response) NotFound() {
	r.Code = http.StatusNotFound
	r.Message = http.StatusText(http.StatusNotFound)
}

func (r *Response) BadRequest(payload any) {
	r.Code = http.StatusBadRequest
	r.Message = http.StatusText(http.StatusBadRequest)
	r.Payload = payload
}

func (r *Response) InternalServerError(payload any) {
	r.Code = http.StatusInternalServerError
	r.Message = http.StatusText(http.StatusInternalServerError)
	r.Payload = payload
}

func (r *Response) Set(code int, message string, payload any) {
	r.Code = code
	r.Message = message
	r.Payload = payload
}
