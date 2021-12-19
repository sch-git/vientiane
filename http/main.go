package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// TODO module api 各个模块的 API 写在这里
	// handleModuleName("module name")
	r.Run("localhost:8080")
}
