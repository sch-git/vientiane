package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"vientiane/http/router"
)

func main() {
	// 记录日志到文件。
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()
	// TODO module api 各个模块的 API 写在这里
	router.HandleHealth(r,"health")
	r.Run("localhost:8080")
}
