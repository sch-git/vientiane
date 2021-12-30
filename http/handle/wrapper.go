package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"vientiane/http/result"
)

type Handler interface {
	Handle(ctx *gin.Context)
}

func BindJsonWrapper(h func() Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		ins := h()
		err := c.Bind(&ins)
		if nil != err {
			glog.Errorf("%v", err)
			result.RespERR(c, &result.Result{
				Code:    http.StatusBadRequest,
				Message: "bind json err",
			})
			return
		}

		ins.Handle(c)
	}
}
