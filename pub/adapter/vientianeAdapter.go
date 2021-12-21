package adapter

import (
	"context"
	"fmt"
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
	client, conn := getClient()
	defer conn.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	res, err = client.HealthCheck(ctx, req)
	if nil != err {
		log.Fatalln(err)
		return
	}

	return
}
