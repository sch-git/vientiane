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
	port   = 8899
	domain = "localhost"
)

var (
	uri = fmt.Sprintf("%s:%d", domain, port)
)

func getClient() vientiane.VientianeServiceClient {
	conn, err := grpc.Dial(uri, grpc.WithInsecure())
	if nil != err {
		log.Fatalln(err)
		return nil
	}
	defer conn.Close()

	client := vientiane.NewVientianeServiceClient(conn)
	return client
}

func HealthCheckByGrpc(ctx context.Context, req *vientiane.HealthCheckReq, options ...grpc.CallOption) (res *vientiane.HealthCheckRes, err error) {
	client := getClient()

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	res, err = client.HealthCheck(ctx, req)
	if nil != err {
		log.Fatalln(err)
		return
	}

	return
}
