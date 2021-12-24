package group

import (
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

// 分布式缓存调度器
type group struct {
	manager *manager
	loader  *singleflight.Group // 防止缓存击穿
	router  *gin.Engine
}

func New(replicas int) *group {
	g := &group{
		manager: newManager(replicas),
		loader:  new(singleflight.Group),
	}
	g.newRouter()

	return g
}

func (g *group) StartHTTP() {
	if err := g.router.Run(":8080"); err != nil {
		log.Fatalln("start http server err:", err)
	}
}

func (g *group) newRouter() {
	router := gin.Default()
	router.POST("/register", g.register)
	router.GET("/cache", g.getCache)
	router.POST("/cache", g.setCache)

	g.router = router
}
