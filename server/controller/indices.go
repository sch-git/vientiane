package controller

import (
	"context"
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

	c.indexDAO.CreateIndex(ctx, req.IdxName, req.IdxConfig)
	return
}
