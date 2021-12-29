package adapter

import (
	"context"
	"fmt"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"log"
	"time"
	vientiane "vientiane/pub/idl/grpc"
)

const (
	Port   = 8899
	domain = "localhost"
)

var (
	uri = fmt.Sprintf("%s:%d", domain, Port)
)

func getClient() (vientiane.VientianeServiceClient, *grpc.ClientConn) {
	conn, err := grpc.Dial(uri, grpc.WithInsecure())
	if nil != err {
		log.Fatalln(err)
		return nil, nil
	}

	client := vientiane.NewVientianeServiceClient(conn)
	return client, conn
}

func HealthCheckByGrpc(ctx context.Context, req *vientiane.HealthCheckReq, options ...grpc.CallOption) (res *vientiane.HealthCheckRes, err error) {
	fun := "HealthCheckByGrpc-->"
	client, conn := getClient()
	defer conn.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	glog.Infof("func: %s req: %v", fun, req)
	res, err = client.HealthCheck(ctx, req)
	if nil != err {
		glog.Fatalln(err)
		return
	}

	return
}

func GetAccountByGrpc(ctx context.Context, req *vientiane.GetAccountReq, options ...grpc.CallOption) (res *vientiane.GetAccountRes, err error) {
	fun := "GetAccountByGrpc-->"
	client, conn := getClient()
	defer conn.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	glog.Infof("func: %s req: %v", fun, req)
	res, err = client.GetAccount(ctx, req)
	return
}

func ListAccountByGrpc(ctx context.Context, req *vientiane.ListAccountReq, options ...grpc.CallOption) (res *vientiane.ListAccountRes, err error) {
	fun := "GetAccountByGrpc-->"
	client, conn := getClient()
	defer conn.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	glog.Infof("func: %s req: %v", fun, req)
	res, err = client.ListAccount(ctx, req)
	return
}
