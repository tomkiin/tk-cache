package group

import (
	"context"
	"log"
	"time"
	"tk-cache/pkg/httpx"
	pb "tk-cache/pkg/proto"

	"github.com/gin-gonic/gin"
)

// 注册 node 节点
func (g *group) register(c *gin.Context) {
	var req = map[string]string{}
	if err := c.ShouldBindJSON(&req); err != nil {
		httpx.RenderErr(c, err)
		return
	}

	ip := req["ip"]
	g.manager.registerNode(ip)

	httpx.RenderOK(c)
}

// 监听 node 节点健康状态，周期：10s
func (g *group) PingNode() {
	ticker := time.NewTicker(10 * time.Second)

	for {
		for ip := range g.manager.nodes {
			go g.pingNode(ip)
		}

		<-ticker.C
	}
}

func (g *group) pingNode(ip string) {
	client, err := g.manager.pickNode(ip)
	if err != nil {
		g.manager.removeNode(ip)
	}

	res, err := client.Ping(context.TODO(), &pb.PingCacheReq{})
	if err != nil || !res.Ok {
		g.manager.removeNode(ip)
		log.Printf("ping node: %s failed\n", ip)
	} else {
		log.Printf("ping node: %s ok\n", ip)
	}
}
