package indices

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"vientiane/http/handle"
	"vientiane/http/result"
	"vientiane/pub/adapter"
	. "vientiane/pub/idl/grpc"
)

type createIndices struct {
	CreateIdxReq
}

func FactoryCreateIndices() handle.Handler {
	return new(createIndices)
}

func (m *createIndices) Handle(ctx *gin.Context) {
	fun := "createIndices.Handle"

	res, err := adapter.CreateIdxByGrpc(ctx, &m.CreateIdxReq)
	if nil != err || res.Code == result.Failed {
		handle.Vlog.Error(fun, zap.Error(err))
		result.RespERR(ctx, res)
		return
	}

	result.RespOK(ctx, res, res.Data)
	return
}
