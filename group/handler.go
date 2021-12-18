package group

import (
	"context"
	"log"
	"time"
	"tk-cache/pkg/httpx"
	pb "tk-cache/pkg/proto"
	"tk-cache/pkg/rpcx"

	"github.com/gin-gonic/gin"
)

// 注册 node 节点
func (g *group) register(c *gin.Context) {
	var req = map[string]string{}
	if err := c.ShouldBindJSON(&req); err != nil {
		httpx.RenderErr(c, err)
		return
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	ip := req["ip"]
	g.nodes[ip] = true
	g.consistent.Add(ip)

	httpx.RenderOK(c)
}

// 监听 node 节点健康状态，周期：10s
func (g *group) PingNode() {
	ticker := time.NewTicker(10 * time.Second)

	for {
		for ip := range g.nodes {
			go g.pingNode(ip)
		}
		<-ticker.C
	}
}

func (g *group) pingNode(ip string) {
	g.mu.Lock()
	defer g.mu.Unlock()

	client, err := rpcx.NewHealthClient(ip + ":8090")
	if err != nil {
		g.nodes[ip] = false
		return
	}

	res, err := client.Ping(context.TODO(), &pb.PingReq{})
	if err != nil || !res.Ok {
		g.nodes[ip] = false
		g.consistent.Del(ip)
		log.Printf("ping node: %s failed\n", ip)
		return
	} else {
		log.Printf("ping node: %s ok\n", ip)
	}
}
