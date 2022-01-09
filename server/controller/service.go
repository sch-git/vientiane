package controller

import (
	"context"
	. "vientiane/pub/idl/grpc"
)

func (s *VientianeServiceImpl) HealthCheck(ctx context.Context, req *HealthCheckReq) (res *HealthCheckRes, err error) {
	return HandleVientiane.health.HealthCheck(ctx, req), nil
}

// account

func (s *VientianeServiceImpl) GetAccount(ctx context.Context, req *GetAccountReq) (*GetAccountRes, error) {
	return HandleVientiane.account.GetAccount(ctx, req), nil
}

func (s *VientianeServiceImpl) ListAccount(ctx context.Context, req *ListAccountReq) (*ListAccountRes, error) {
	return HandleVientiane.account.ListAccount(ctx, req), nil
}

// category

func (s *VientianeServiceImpl) ListCategory(ctx context.Context, req *ListCategoryReq) (*ListCategoryRes, error) {
	panic("unimplemented func")
}
