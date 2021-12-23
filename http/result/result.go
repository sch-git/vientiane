package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resulter interface {
	GetCode() int64
	GetMsg() string
}

type Result struct {
	Message string
	Code    int64
	Data    interface{}
}

func (r *Result) GetCode() int64 {
	return r.Code
}

func (r *Result) GetMsg() string {
	return r.Message
}

func RespOKWithoutData(c *gin.Context, r Resulter) {
	c.JSON(http.StatusOK, &Result{
		Code:    r.GetCode(),
		Message: r.GetMsg(),
	})
}

func RespOK(c *gin.Context, r Resulter, data interface{}) {
	c.JSON(http.StatusOK, &Result{
		Code:    r.GetCode(),
		Message: r.GetMsg(),
		Data:    data,
	})
}

func RespERR(c *gin.Context, r Resulter) {
	c.JSON(http.StatusInternalServerError, &Result{
		Code:    r.GetCode(),
		Message: r.GetMsg(),
	})
}
