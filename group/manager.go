package group

import (
	"fmt"
	"log"
	"sync"
	"tk-cache/pkg/hash"
	pb "tk-cache/pkg/proto"
	"tk-cache/pkg/rpcx"
)

// 缓存节点管理器
type manager struct {
	mu         *sync.Mutex
	nodes      map[string]bool  // 节点注册信息
	consistent *hash.Consistent // 一致性 hash 环
}

func newManager(replicas int) *manager {
	return &manager{
		mu:         new(sync.Mutex),
		nodes:      make(map[string]bool),
		consistent: hash.NewConsistent(replicas, nil),
	}
}

// 注册节点
func (m *manager) registerNode(ip string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.nodes[ip] = true
	m.consistent.Add(ip)
}

// 注销节点
func (m *manager) removeNode(ip string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.nodes[ip] = false
	m.consistent.Del(ip)
}

// 根据 ip 获取对应的节点客户端
func (m *manager) getNode(ip string) (pb.CacheClient, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if !m.nodes[ip] {
		return nil, fmt.Errorf("node: %s not registered", ip)
	}

	return rpcx.NewCacheClient(ip + ":8090")
}

// 根据一致性哈希获取对应的节点客户端
func (m *manager) pickNode(key string) (pb.CacheClient, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	ip := m.consistent.Get(key)
	log.Println("pick node:", ip)
	if !m.nodes[ip] {
		return nil, fmt.Errorf("node: %s not registered", ip)
	}

	return rpcx.NewCacheClient(ip + ":8090")
}
