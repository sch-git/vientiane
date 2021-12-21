package controller

import (
	"context"
	. "vientiane/pub/idl/grpc"
)

func (s *VientianeServiceImpl) HealthCheck(ctx context.Context, req *HealthCheckReq) (res *HealthCheckRes, err error) {
	return HandleVientiane.health.HealthCheck(ctx, req), nil
}
