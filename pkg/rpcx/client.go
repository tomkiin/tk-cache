package rpcx

import (
	pb "tk-cache/pkg/proto"

	"google.golang.org/grpc"
)

func NewCacheClient(addr string) (pb.CacheClient, error) {
	cc, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return pb.NewCacheClient(cc), nil
}
