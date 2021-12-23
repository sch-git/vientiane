package controller

import (
	"context"
	"github.com/golang/glog"
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
		glog.Errorf("%s req is nil", fun)
		res = &vientiane.GetAccountRes{Code: models.InvalidReqIsNil, Msg: models.InvalidReqIsNilMsg}
		return
	}

	account, err := c.service.Get(ctx, req.Id)
	if nil != err {
		glog.Errorf("%s %v", fun, err)
		res = &vientiane.GetAccountRes{Code: models.ServerGetErr, Msg: err.Error()}
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
