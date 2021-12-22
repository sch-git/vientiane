package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Message string
	Code    int
	Data    interface{}
}

func RespOK(c *gin.Context, result *Result) {
	c.JSON(http.StatusOK, &result)
}

func RespERR(c *gin.Context, result *Result) {
	c.JSON(http.StatusInternalServerError, &result)
}
