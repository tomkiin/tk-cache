package node

import (
	"sync"
	"tk-cache/pkg/cache"
)

// 缓存节点
type node struct {
	mu  *sync.Mutex // 确保并发安全
	lru *cache.LRU
}

func New(maxSize int) *node {
	return &node{
		mu:  new(sync.Mutex),
		lru: cache.NewLRU(maxSize),
	}
}

func (n *node) Get(key string) ([]byte, bool) {
	n.mu.Lock()
	defer n.mu.Unlock()

	return n.lru.Get(key)
}

func (n *node) Set(key string, value []byte) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.lru.Set(key, value)
}
