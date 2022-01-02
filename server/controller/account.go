package controller

import (
	"context"
	"go.uber.org/zap"
	vientiane "vientiane/pub/idl/grpc"
	"vientiane/server/models"
	"vientiane/server/service"
)

type accountController struct {
	service models.AccountService
}

func NewAccountController() *accountController {
	return &accountController{
		service: service.NewAccountService(),
	}
}

func (c *accountController) GetAccount(ctx context.Context, req *vientiane.GetAccountReq) (res *vientiane.GetAccountRes) {
	fun := "accountController.GetAccount-->"

	if nil == req {
		vlog.Error(fun, zap.String("req", "req is nil"))
		res = &vientiane.GetAccountRes{Code: models.InvalidReqIsNil, Msg: models.InvalidReqIsNilMsg}
		return
	}

	account, err := c.service.Get(ctx, req.Id)
	if nil != err {
		vlog.Error(fun, zap.Error(err))
		res = &vientiane.GetAccountRes{Code: models.ServerErr, Msg: err.Error()}
		return
	}

	res = &vientiane.GetAccountRes{
		Code: models.StatusOK,
		Data: &vientiane.GetAccountData{
			Account: account.ToGrpc(),
		},
	}
	return
}

func (c *accountController) ListAccount(ctx context.Context, req *vientiane.ListAccountReq) (res *vientiane.ListAccountRes) {
	fun := "accountController.ListAccount-->"
	res = &vientiane.ListAccountRes{}

	if nil == req {
		vlog.Error(fun, zap.String("req", "req is nil"))
		res = &vientiane.ListAccountRes{Code: models.InvalidReqIsNil, Msg: models.InvalidReqIsNilMsg}
		return
	}

	account := &models.Account{
		Name:   req.Name,
		Email:  req.Email,
		Limit:  req.Limit,
		Offset: req.Offset,
	}
	accounts, err := c.service.List(ctx, account)
	if nil != err {
		vlog.Error(fun, zap.Error(err))
		res = &vientiane.ListAccountRes{Code: models.ServerErr, Msg: err.Error()}
		return
	}

	res.Data = &vientiane.ListAccountData{
		Accounts: func() []*vientiane.Account {
			grpcAccounts := make([]*vientiane.Account, 0)
			if len(accounts) < 1 {
				return grpcAccounts
			}

			for _, account := range accounts {
				if nil == account {
					continue
				}

				grpcAccounts = append(grpcAccounts, account.ToGrpc())
			}
			return grpcAccounts
		}(),
		Offset: req.Offset + int64(len(accounts)),
		//Count: TODO
	}

	return
}
