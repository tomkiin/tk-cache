package node

import (
	"context"
	"errors"
	"fmt"
	"log"
	"tk-cache/pkg/httpx"
	"tk-cache/pkg/network"
	pb "tk-cache/pkg/proto"
)

// 向 group 注册本地 ip 信息
func (n *node) Register(ip string) error {
	if ip == "" {
		return errors.New("register ip is nil")
	}

	localIP, err := network.GetLocalIP()
	if err != nil {
		return err
	}

	registerURL := fmt.Sprintf("http://%s:8080/register", ip)
	body := map[string]string{"ip": localIP}

	if err := httpx.Post(registerURL, body); err != nil {
		return err
	}

	log.Printf("register ip: %s ok\n", ip)
	return nil
}

func (n *node) Ping(ctx context.Context, req *pb.PingCacheReq) (*pb.PingCacheRes, error) {
	return &pb.PingCacheRes{Ok: true}, nil
}

func (n *node) Get(ctx context.Context, req *pb.GetCacheReq) (*pb.GetCacheRes, error) {
	n.mu.Lock()
	defer n.mu.Unlock()

	value, ok := n.lru.Get(req.Key)
	if ok {
		log.Println("hit cache key:", req.Key)
	}

	return &pb.GetCacheRes{Value: value, Ok: ok}, nil
}

func (n *node) Set(ctx context.Context, req *pb.SetCacheReq) (*pb.SetCacheRes, error) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.lru.Set(req.Key, req.Value)
	log.Printf("set cache key: %s, value: %s\n", req.Key, req.Value)

	return &pb.SetCacheRes{}, nil
}
