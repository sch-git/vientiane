package controller

import (
	"context"
	vientiane "vientiane/pub/idl/grpc"
)

type accountController struct {
}

func NewAccountController() *accountController {
	return &accountController{}
}

func (c *accountController) GetAccount(ctx context.Context, req *vientiane.GetAccountReq) (res *vientiane.GetAccountRes) {
	res = &vientiane.GetAccountRes{}
	return
}
