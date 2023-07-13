package controller

import (
	"go.uber.org/zap"
	vientiane "vientiane/pub/idl/grpc"
)

type VientianeServiceImpl struct {
	health                                 *healthController
	account                                *accountController
	article                                *articleController
	indices                                *indicesController
	vientiane.UnsafeVientianeServiceServer // 通过这个类，继承了 mustEmbedUnimplementedVientianeServiceServer 方法
}

var (
	HandleVientiane *VientianeServiceImpl
	vlog            *zap.Logger
)

func init() {
	HandleVientiane = &VientianeServiceImpl{
		health:  NewHealthController(),
		account: NewAccountController(),
		article: NewArticleController(),
		indices: NewIndicesController(),
	}
	vlog = zap.NewExample()
	defer vlog.Sync()
}
