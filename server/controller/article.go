package controller

import (
	"context"
	"go.uber.org/zap"
	"log"
	vientiane "vientiane/pub/idl/grpc"
	"vientiane/server/consts"
	"vientiane/server/models"
	"vientiane/server/service"
)

type articleController struct {
	service models.ArticleService
}

func NewArticleController() *articleController {
	return &articleController{
		service: service.NewArticleService(),
	}
}

func (c *articleController) GetArticle(ctx context.Context, req *vientiane.GetArticleReq) (res *vientiane.GetArticleRes) {
	fun := "articleController.GetArticle-->"

	if nil == req {
		vlog.Error(fun, zap.String("req", "req is nil"))
		res = &vientiane.GetArticleRes{Code: consts.InvalidReqIsNil, Msg: consts.InvalidReqIsNilMsg}
		return
	}

	article, err := c.service.Get(ctx, req.Id)
	if nil != err {
		vlog.Error(fun, zap.Error(err))
		res = &vientiane.GetArticleRes{Code: consts.ServerErr, Msg: err.Error()}
		return
	}
	log.Println(req.GetId())
	log.Println(article)

	res = &vientiane.GetArticleRes{
		Code: consts.StatusOK,
		Data: &vientiane.GetArticleData{
			Article: article.ToGrpc(),
		},
	}
	return
}
