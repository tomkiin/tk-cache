package hash

import (
	"fmt"
	"hash/crc32"
	"sort"
)

type Hash func([]byte) uint32

// 一致性哈希算法
type Consistent struct {
	hash     Hash
	replicas int
	circle   []int
	mapping  map[int]string
}

func NewConsistent(replicas int, hash Hash) *Consistent {
	c := &Consistent{
		hash:     hash,
		replicas: replicas,
		circle:   make([]int, 0),
		mapping:  make(map[int]string),
	}
	if hash == nil {
		c.hash = crc32.ChecksumIEEE
	}

	return c
}

func (c *Consistent) Add(key string) {
	for i := 0; i < c.replicas; i++ {
		vkey := fmt.Sprintf("%d%s", i, key) // 虚拟节点 key
		hash := int(c.hash([]byte(vkey)))
		c.circle = append(c.circle, hash)
		c.mapping[hash] = key
	}

	sort.Ints(c.circle)
}

func (c *Consistent) Get(key string) string {
	if len(c.circle) == 0 {
		return ""
	}

	hash := int(c.hash([]byte(key)))
	idx := sort.Search(len(c.circle), func(i int) bool { // 寻找顺时针最近节点
		return c.circle[i] >= hash
	})

	return c.mapping[c.circle[idx%len(c.circle)]]
}

func (c *Consistent) Del(key string) {
	for i := 0; i < c.replicas; i++ {
		vkey := fmt.Sprintf("%d%s", i, key)
		hash := int(c.hash([]byte(vkey)))
		c.delHash(hash)
	}

	sort.Ints(c.circle)
}

func (c *Consistent) delHash(hash int) {
	for i := range c.circle {
		if hash == c.circle[i] {
			c.circle = append(c.circle[:i], c.circle[i+1:]...)
			delete(c.mapping, hash)

			return
		}
	}
}
