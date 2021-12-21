package controller

import vientiane "vientiane/pub/idl/grpc"

type VientianeServiceImpl struct {
	health                                 *HealthController
	vientiane.UnsafeVientianeServiceServer // 通过这个类，继承了 mustEmbedUnimplementedVientianeServiceServer 方法
}

var HandleVientiane *VientianeServiceImpl

func init() {
	HandleVientiane = &VientianeServiceImpl{
		health: NewHealthController(),
	}
}
