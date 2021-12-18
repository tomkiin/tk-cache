package hash

import (
	"strconv"
	"testing"
)

var myHash = func(data []byte) uint32 {
	i, _ := strconv.Atoi(string(data))
	return uint32(i)
}

func TestConsistent(t *testing.T) {
	c := NewConsistent(3, myHash)
	c.Add("2")
	c.Add("4")
	c.Add("6")
	// 当前 hash 环: 02 -> 04 -> 06 -> 12 -> 14 -> 16 -> 22 -> 24 -> 26 -> 02
	testCase := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}
	for k, v := range testCase {
		if c.Get(k) != v {
			t.Errorf("asking for %s, expect %s but get %s\n", k, v, c.Get(k))
		}
	}

	c.Add("8")
	// 当前 hash 环: 02 -> 04 -> 06 -> 08 -> 12 -> 14 -> 16 -> 18 -> 22 -> 24 -> 26 -> 28 -> 02
	testCase = map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "8",
	}
	for k, v := range testCase {
		if c.Get(k) != v {
			t.Errorf("asking for %s, expect %s but get %s\n", k, v, c.Get(k))
		}
	}

	c.Del("4")
	// 当前 hash 环: 02 -> 06 -> 08 -> 12 -> 16 -> 18 -> 22 -> 26 -> 28 -> 02
	testCase = map[string]string{
		"2":  "2",
		"11": "2",
		"23": "6",
		"27": "8",
	}
	for k, v := range testCase {
		if c.Get(k) != v {
			t.Errorf("asking for %s, expect %s but get %s\n", k, v, c.Get(k))
		}
	}
}
