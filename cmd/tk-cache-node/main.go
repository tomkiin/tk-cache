package main

import (
	"fmt"
	"tk-cache/node"
)

func main() {
	n := node.New(0)
	n.Set("k1", []byte("v1"))
	value, ok := n.Get("k1")

	fmt.Println(string(value), ok)
}
