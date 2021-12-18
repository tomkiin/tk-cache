package main

import "tk-cache/group"

func main() {
	g := group.New()
	go g.PingNode()
	g.StartHTTP()
}
