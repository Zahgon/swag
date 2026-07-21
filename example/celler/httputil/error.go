package httputil

import "github.com/gin-gonic/gin"

func NewError(ctx *gin.Context, status int, err error) { _ = "STUB: not implemented"; return }

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
