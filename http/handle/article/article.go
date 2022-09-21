package article

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"vientiane/http/handle"
	"vientiane/http/result"
	"vientiane/pub/adapter"
	. "vientiane/pub/idl/grpc"
)

type getArticle struct {
	GetArticleReq
}

func FactoryGetArticle() handle.Handler {
	return new(getArticle)
}

func (m *getArticle) Handle(ctx *gin.Context) {
	fun := "GetArticleReq.Handle"

	id, _ := strconv.Atoi(ctx.Param("id"))
	m.Id = int64(id)

	res, err := adapter.GetArticleByGrpc(ctx, &m.GetArticleReq)
	if nil != err || res.Code == result.Failed {
		handle.Vlog.Error(fun, zap.Error(err))
		result.RespERR(ctx, res)
		return
	}

	result.RespOK(ctx, res, res.Data)
	return
}

