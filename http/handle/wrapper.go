package handle

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"vientiane/http/result"
)

var Vlog = zap.NewExample()

type Handler interface {
	Handle(ctx *gin.Context)
}

func BindJsonWrapper(h func() Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		ins := h()
		err := c.Bind(&ins)
		if nil != err {
			Vlog.Error("BindJsonWrapper", zap.Error(err))
			result.RespERR(c, &result.Result{
				Code:    http.StatusBadRequest,
				Message: "bind json err",
			})
			return
		}

		ins.Handle(c)
	}
}
