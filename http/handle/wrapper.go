package handle

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"vientiane/http/result"
)

type Handle interface {
	Handle(ctx *gin.Context)
}

func BindJsonWrapper(h func() Handle) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := c.Bind(h)
		if nil != err {
			log.Println("bind json err")
			result.RespERR(c, &result.Result{
				Code:    http.StatusBadRequest,
				Message: "bind json err",
			})
			return
		}

		h().Handle(c)
	}
}
