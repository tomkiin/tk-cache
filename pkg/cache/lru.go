package cache

import "container/list"

// 最近最少使用缓存淘汰策略
type LRU struct {
	maxSize int
	curSize int
	ll      *list.List
	mapping map[string]*list.Element
}

// 缓存键值对
type entry struct {
	key   string
	value []byte
}

// 新建 lru 缓存，设置最大容量 0 则为无限容量
func NewLRU(maxSize int) *LRU {
	return &LRU{
		maxSize: maxSize,
		curSize: 0,
		ll:      list.New(),
		mapping: make(map[string]*list.Element),
	}
}

func (l *LRU) Get(key string) ([]byte, bool) {
	ele, ok := l.mapping[key]
	if !ok {
		return nil, false
	}

	l.ll.MoveToFront(ele)
	e := ele.Value.(*entry)

	return e.value, true
}

func (l *LRU) Set(key string, value []byte) {
	if ele, ok := l.mapping[key]; ok {
		l.ll.MoveToFront(ele)
		e := ele.Value.(*entry)
		l.curSize += len(value) - len(e.value)
		e.value = value
	} else {
		ele := l.ll.PushFront(&entry{key: key, value: value})
		l.mapping[key] = ele
		l.curSize += len(key) + len(value)
	}
	// 淘汰最久未使用缓存
	for l.maxSize > 0 && l.maxSize < l.curSize {
		l.removeOldest()
	}
}

func (l *LRU) removeOldest() {
	ele := l.ll.Back()
	if ele == nil {
		return
	}

	l.ll.Remove(ele)
	e := ele.Value.(*entry)
	delete(l.mapping, e.key)
	l.curSize -= len(e.key) + len(e.value)
}
