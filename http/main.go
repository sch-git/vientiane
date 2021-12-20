package main

import (
	"github.com/gin-gonic/gin"
	"vientiane/http/handle"
)

func main() {
	r := gin.Default()
	// TODO module api 各个模块的 API 写在这里
	handle.HandleHealth(r,"health")
	r.Run("localhost:8080")
}
