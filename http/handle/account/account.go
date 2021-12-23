package account

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"strconv"
	"vientiane/http/handle"
	"vientiane/http/result"
	"vientiane/pub/adapter"
	. "vientiane/pub/idl/grpc"
)

type getAccount struct {
	GetAccountReq
}

func FactoryGetAccount() handle.Handle {
	return new(getAccount)
}

func (m *getAccount) Handle(ctx *gin.Context) {
	fun := "getAccount.Handle"

	id, _ := strconv.Atoi(ctx.Param("id"))
	m.Id = int64(id)

	res, err := adapter.GetAccountByGrpc(ctx, &m.GetAccountReq)
	if nil != err || res.Code == result.Failed {
		glog.Errorf("%s %v", fun, err)
		result.RespERR(ctx, res)
		return
	}

	result.RespOK(ctx, res, res.Data)
	return
}
