package controller

import (
	"context"
	"go-micro.dev/v4/util/log"
	vientiane "vientiane/pub/idl/grpc"
	"vientiane/server/dao"
)

type indicesController struct {
	indexDAO dao.IndexDAO
}

func NewIndicesController() *indicesController {
	return &indicesController{
		indexDAO: dao.NewIndexDAO(),
	}
}

func (c *indicesController) CreateIdx(ctx context.Context, req *vientiane.CreateIdxReq) (resp *vientiane.CreateIdxRes, err error) {

	err = c.indexDAO.CreateIndex(ctx, req.IdxName, req.IdxConfig)
	if err != nil {
		log.Errorf("create index err: %+v", err)
		resp = &vientiane.CreateIdxRes{
			Code: -1,
			Msg:  "create index failed",
		}
		return
	}

	resp = &vientiane.CreateIdxRes{
		Msg: "create success",
	}
	return
}
