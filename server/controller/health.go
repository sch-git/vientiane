package controller

import (
	"context"
	. "vientiane/pub/idl/grpc"
)

type healthController struct {
}

func NewHealthController() *healthController {
	return &healthController{}
}

func (c *healthController) HealthCheck(ctx context.Context, req *HealthCheckReq) (res *HealthCheckRes) {
	res = &HealthCheckRes{
		Data: &HealthCheckData{},
	}
	return
}
