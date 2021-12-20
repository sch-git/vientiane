package adapter

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	. "pub/idl/grpc"
)

const (
	port   = 8899
	domain = "localhost"
)

var (
	uri = fmt.Sprintf("%s:%d", domain, port)
)

func getClient() VientianeServiceClient {
	conn, err := grpc.Dial(uri, grpc.WithInsecure())
	if nil != err {
		log.Fatalln(err)
		return nil
	}
	defer conn.Close()

	client := NewVientianeServiceClient(conn)
	return client
}

func HealthCheck(ctx context.Context, req *HealthCheckReq, options ...grpc.CallOption) (res *HealthCheckRes, err error) {
	client := getClient()

	res, err = client.HealthCheck(context.Background(), req)
	if nil != err {
		log.Fatalln(err)
		return
	}

	return
}
