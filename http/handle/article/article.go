package article

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"vientiane/http/handle"
	"vientiane/http/result"
	"vientiane/http/utils"
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


type addArticle struct {
	Article
}

func FactoryAddArticle() handle.Handler {
	return new(addArticle)
}

func (m *addArticle)Handle(ctx *gin.Context)  {
	msg := &ArticleMsg{
		WriteType: adapter.ArticleTypeInsert,
		Article: &m.Article,
	}
	go utils.WriteMsg(msg)
}