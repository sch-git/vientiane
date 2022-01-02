package controller

import (
	"go.uber.org/zap"
	vientiane "vientiane/pub/idl/grpc"
)

type VientianeServiceImpl struct {
	health                                 *healthController
	account                                *accountController
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
	}
	vlog = zap.NewExample()
	defer vlog.Sync()
}
