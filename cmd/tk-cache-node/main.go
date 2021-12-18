package main

import (
	"log"
	"tk-cache/node"
)

func main() {
	n := node.New(0)
	if err := n.Register(); err != nil {
		log.Fatalln(err)
	}
	n.StartRPC()
}
