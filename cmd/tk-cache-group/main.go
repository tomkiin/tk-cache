package main

import "tk-cache/group"

func main() {
	g := group.New()
	g.StartHTTP()
}
