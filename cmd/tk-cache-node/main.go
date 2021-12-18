package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"tk-cache/node"
	pb "tk-cache/pkg/proto"

	"google.golang.org/grpc"
)

func main() {
	n := node.New(0)
	go n.StartRPC()

	time.Sleep(time.Second)
	rpcClientTest()
}

func rpcClientTest() {
	cc, err := grpc.Dial("127.0.0.1:8090", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("grpc dial err:", err)
	}

	client := pb.NewCacheClient(cc)

	ctx := context.TODO()
	client.Set(ctx, &pb.SetCacheReq{Key: "k1", Value: []byte("v1")})
	res, _ := client.Get(ctx, &pb.GetCacheReq{Key: "k1"})

	fmt.Println(string(res.Value))
}
