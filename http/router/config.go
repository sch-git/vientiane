package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RouteConfig struct {
	Path   string
	Handle gin.HandlerFunc
}

type Result struct {
	Message string
	Code    int
}

func RespOK(c *gin.Context, result *Result) {
	c.JSON(http.StatusOK, &result)
}
