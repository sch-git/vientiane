package health

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"vientiane/http/handle"
	"vientiane/http/result"
	"vientiane/pub/adapter"
	pub "vientiane/pub/idl/grpc"
)

type healthCheck struct {
	pub.HealthCheckReq
}

func FactoryHealthCheck() handle.Handle {
	return new(healthCheck)
}

func (m *healthCheck) Handle(ctx *gin.Context) {
	res, err := adapter.HealthCheckByGrpc(ctx, &m.HealthCheckReq)
	if nil != err || res.Code == result.Failed {
		log.Println("err")
		result.RespERR(ctx, &result.Result{
			Message: http.StatusText(http.StatusInternalServerError),
			Code:    http.StatusInternalServerError,
		})

		return
	}

	result.RespOK(ctx, &result.Result{
		Message: http.StatusText(http.StatusOK),
		Code:    http.StatusOK,
	}, res)
	return
}
