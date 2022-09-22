package article

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
	"log"
	"strconv"
	"time"
	"vientiane/http/handle"
	"vientiane/http/result"
	"vientiane/http/utils"
	"vientiane/pub/adapter"
	. "vientiane/pub/idl/grpc"
)

var (
	articleCache = cache.New(time.Minute,time.Minute)
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
	val,ok:= articleCache.Get(ctx.Param("id"))
	if !ok {
		res, err := adapter.GetArticleByGrpc(ctx, &m.GetArticleReq)
		if nil != err || res.Code == result.Failed {
			handle.Vlog.Error(fun, zap.Error(err))
			result.RespERR(ctx, res)
			return
		}
		articleCache.Set(ctx.Param("id"),res,time.Second*5)
		result.RespOK(ctx, res, res.Data)
	}else{
		res := val.(*GetArticleRes)
		result.RespOK(ctx, res, res.Data)
	}

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
	log.Println(msg)
	go utils.WriteMsg(msg)
}


type delArticle struct {
	Article
}

func FactoryDelArticle() handle.Handler {
	return new(delArticle)
}

func (m *delArticle)Handle(ctx *gin.Context)  {
	msg := &ArticleMsg{
		WriteType: adapter.ArticleTypeDelete,
		Article: &m.Article,
	}
	go utils.WriteMsg(msg)
}

type setArticle struct {
	Article
}

func FactorySetArticle() handle.Handler {
	return new(setArticle)
}

func (m *setArticle)Handle(ctx *gin.Context)  {
	msg := &ArticleMsg{
		WriteType: adapter.ArticleTypeUpdate,
		Article: &m.Article,
	}
	go utils.WriteMsg(msg)
}