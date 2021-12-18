package node

import (
	"log"
	"net"
	"sync"
	"tk-cache/pkg/cache"
	pb "tk-cache/pkg/proto"

	"google.golang.org/grpc"
)

// 缓存节点
type node struct {
	pb.UnimplementedCacheServer // 升级成 grpc 服务器

	mu  *sync.Mutex // 确保并发安全
	lru *cache.LRU
}

func New(maxSize int) *node {
	return &node{
		mu:  new(sync.Mutex),
		lru: cache.NewLRU(maxSize),
	}
}

func (n *node) StartRPC() {
	l, err := net.Listen("tcp", "0.0.0.0:8090")
	if err != nil {
		log.Fatalln("net listen err:", err)
	}
	defer l.Close()

	s := grpc.NewServer()
	// 注册 grpc 服务
	pb.RegisterCacheServer(s, n)

	if err := s.Serve(l); err != nil {
		log.Fatalln("start rpc server err:", err)
	}
}
