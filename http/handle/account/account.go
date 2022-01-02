package account

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"vientiane/http/handle"
	"vientiane/http/result"
	"vientiane/pub/adapter"
	. "vientiane/pub/idl/grpc"
)

type getAccount struct {
	GetAccountReq
}

func FactoryGetAccount() handle.Handler {
	return new(getAccount)
}

func (m *getAccount) Handle(ctx *gin.Context) {
	fun := "getAccount.Handle"

	id, _ := strconv.Atoi(ctx.Param("id"))
	m.Id = int64(id)

	res, err := adapter.GetAccountByGrpc(ctx, &m.GetAccountReq)
	if nil != err || res.Code == result.Failed {
		handle.Vlog.Error(fun, zap.Error(err))
		result.RespERR(ctx, res)
		return
	}

	result.RespOK(ctx, res, res.Data)
	return
}

type listAccount struct {
	ListAccountReq
}

func FactoryListAccount() handle.Handler {
	return new(listAccount)
}

func (m *listAccount) Handle(ctx *gin.Context) {
	fun := "listAccount.Handle-->"

	res, err := adapter.ListAccountByGrpc(ctx, &m.ListAccountReq)
	if nil != err || res.Code == result.Failed {
		handle.Vlog.Error(fun, zap.Error(err))
		result.RespERR(ctx, res)
		return
	}

	result.RespOK(ctx, res, res.Data)
	return
}
