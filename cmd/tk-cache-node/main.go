package main

import (
	"flag"
	"log"
	"tk-cache/node"
)

var (
	maxSize    int
	registerIP string
)

func init() {
	flag.IntVar(&maxSize, "max_size", 10*1024, "max cache size, zero is nif")
	flag.StringVar(&registerIP, "register_ip", "", "remote group ip")
}

func main() {
	flag.Parse()

	n := node.New(maxSize)
	if err := n.Register(registerIP); err != nil {
		log.Fatalln(err)
	}
	n.StartRPC()
}
