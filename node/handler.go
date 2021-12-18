package node

import (
	"context"
	"log"
	"tk-cache/pkg/httpx"
	"tk-cache/pkg/network"
	pb "tk-cache/pkg/proto"
)

// 向 group 注册本地 ip 信息
func (n *node) Register() error {
	ip, err := network.GetLocalIP()
	if err != nil {
		return err
	}

	registerURL := "http://127.0.0.1:8080/register"
	body := map[string]string{"ip": ip}

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

	return &pb.GetCacheRes{Value: value, Ok: ok}, nil
}

func (n *node) Set(ctx context.Context, req *pb.SetCacheReq) (*pb.SetCacheRes, error) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.lru.Set(req.Key, req.Value)

	return &pb.SetCacheRes{}, nil
}
