package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vientiane/http/router"
	"vientiane/pub/adapter"
	pub "vientiane/pub/idl/grpc"
)

type healthCheckReq struct {
	req *pub.HealthCheckReq
}

func (m *healthCheckReq) Handle(ctx *gin.Context) {
	_, err := adapter.HealthCheckByGrpc(ctx, m.req)
	if nil != err {
		router.RespOK(ctx, &router.Result{
			Message: "OK",
			Code:    200,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"data":    "name",
	})
}

func HealthCheckIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		m := &healthCheckReq{}
		err := c.Bind(m)
		if nil != err {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		m.Handle(c)
	}
}
