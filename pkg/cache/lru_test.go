package cache

import "testing"

func TestLruGet(t *testing.T) {
	l := NewLRU(0)
	l.Set("k1", []byte("v1"))
	if v, ok := l.Get("k1"); !ok || string(v) != "v1" {
		t.Fatal("cache hit key = k1 failed")
	}
}

func TestLruRemoveOldest(t *testing.T) {
	l := NewLRU(8)
	l.Set("k1", []byte("v1"))
	l.Set("k2", []byte("v2"))
	l.Set("k3", []byte("v3"))
	// k1 会被淘汰
	if _, ok := l.Get("k1"); ok || len(l.mapping) != 2 {
		t.Fatal("cache remove oldest k1 failed")
	}
}
