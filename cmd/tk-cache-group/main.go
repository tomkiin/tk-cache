package main

import "tk-cache/group"

func main() {
	g := group.New(3)
	go g.PingNode()
	g.StartHTTP()
}
