package main

import (
	"flag"
	"tk-cache/group"
)

var replicas int

func init() {
	flag.IntVar(&replicas, "replicas", 3, "consistent hash virtual replicas")
}

func main() {
	flag.Parse()

	g := group.New(replicas)
	go g.PingNode()
	g.StartHTTP()
}
