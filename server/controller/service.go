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

func (s *VientianeServiceImpl) GetArticle(ctx context.Context, req *GetArticleReq) (*GetArticleRes, error) {
	return HandleVientiane.article.GetArticle(ctx, req), nil
}

// es 索引相关操作

func (s *VientianeServiceImpl) CreateIdx(ctx context.Context, req *CreateIdxReq) (*CreateIdxRes, error) {
	//TODO implement me
	panic("implement me")
}
