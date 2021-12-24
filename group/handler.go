package group

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
	"tk-cache/pkg/httpx"
	pb "tk-cache/pkg/proto"

	"github.com/gin-gonic/gin"
)

var ctx = context.TODO()

// 注册 node 节点
func (g *group) register(c *gin.Context) {
	var req struct {
		IP string `json:"ip"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		httpx.RenderErr(c, err)
		return
	}

	ip := req.IP
	if ip == "" {
		httpx.RenderErr(c, errors.New("invalid ip"))
		return
	}

	g.manager.registerNode(ip)

	httpx.RenderOK(c)
}

func (g *group) getCache(c *gin.Context) {
	key := c.Query("key")

	client, err := g.manager.pickNode(key)
	if err != nil {
		httpx.RenderErr(c, err)
		return
	}
	// 使用 singleflight 机制防止缓存击穿
	view, err, _ := g.loader.Do(key, func() (interface{}, error) {
		return client.Get(ctx, &pb.GetCacheReq{Key: key})
	})
	if err != nil {
		httpx.RenderErr(c, err)
		return
	}

	res := view.(*pb.GetCacheRes)
	if !res.Ok {
		httpx.RenderErr(c, fmt.Errorf("key: \"%s\" is not existed", key))
		return
	}

	httpx.RenderData(c, string(res.Value))
}

func (g *group) setCache(c *gin.Context) {
	var req struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		httpx.RenderErr(c, err)
		return
	}

	client, err := g.manager.pickNode(req.Key)
	if err != nil {
		httpx.RenderErr(c, err)
		return
	}

	if _, err := client.Set(ctx, &pb.SetCacheReq{Key: req.Key, Value: []byte(req.Value)}); err != nil {
		httpx.RenderErr(c, err)
		return
	}

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
	client, err := g.manager.getNode(ip)
	if err != nil {
		g.manager.removeNode(ip)
		log.Printf("ping node: %s failed: %s\n", ip, err)
		return
	}

	res, err := client.Ping(ctx, &pb.PingCacheReq{})
	if err != nil || !res.Ok {
		g.manager.removeNode(ip)
		log.Printf("ping node: %s failed: %s\n", ip, err)
		return
	}
}
