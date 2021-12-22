package router

import (
	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	Path   string
	Handle gin.HandlerFunc
}

