package controller

import (
	"context"
	vientiane "vientiane/pub/idl/grpc"
)

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (c *HealthController) HealthCheck(ctx context.Context, req *vientiane.HealthCheckReq) (res *vientiane.HealthCheckRes) {
	res = &vientiane.HealthCheckRes{}
	return
}
