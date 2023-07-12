package adapter

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
	vientiane "vientiane/pub/idl/grpc"
)

const (
	Port   = 8899
	domain = "localhost"
)

var (
	uri  = fmt.Sprintf("%s:%d", domain, Port)
	vlog = zap.NewExample()
)

func getClient() (vientiane.VientianeServiceClient, *grpc.ClientConn) {
	conn, err := grpc.Dial(uri, grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	res, err = client.HealthCheck(ctx, req)
	if nil != err {
		vlog.Fatal(fun, zap.Error(err))
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

	vlog.Info(fun, zap.Any("req", req))
	res, err = client.GetAccount(ctx, req)
	return
}

func ListAccountByGrpc(ctx context.Context, req *vientiane.ListAccountReq, options ...grpc.CallOption) (res *vientiane.ListAccountRes, err error) {
	fun := "ListAccountByGrpc-->"
	client, conn := getClient()
	defer conn.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	vlog.Info(fun, zap.Any("req", req))
	res, err = client.ListAccount(ctx, req)
	return
}

func ListCategoryByGrpc(ctx context.Context, req *vientiane.ListCategoryReq, options ...grpc.CallOption) (res *vientiane.ListCategoryRes, err error) {
	fun := "ListCategoryByGrpc-->"
	client, conn := getClient()
	defer conn.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	vlog.Info(fun, zap.Any("req", req))
	res, err = client.ListCategory(ctx, req)
	return
}

func GetArticleByGrpc(ctx context.Context, req *vientiane.GetArticleReq, options ...grpc.CallOption) (res *vientiane.GetArticleRes, err error) {
	fun := "GetArticleByGrpc-->"
	client, conn := getClient()
	defer conn.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	vlog.Info(fun, zap.Any("req", req))
	res, err = client.GetArticle(ctx, req)
	return
}

func CreateIdxByGrpc(ctx context.Context, req *vientiane.CreateIdxReq, options ...grpc.CallOption) (res *vientiane.CreateIdxRes, err error) {
	fun := "CreateIdxByGrpc-->"
	client, conn := getClient()
	defer conn.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	vlog.Info(fun, zap.Any("req", req))
	res, err = client.CreateIdx(ctx, req)
	return
}
