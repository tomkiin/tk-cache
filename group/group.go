package group

import (
	"log"

	"github.com/gin-gonic/gin"
)

// 分布式缓存调度器
type group struct {
	manager *manager
	router  *gin.Engine
}

func New(replicas int) *group {
	g := &group{
		manager: newManager(replicas),
	}

	router := gin.Default()
	router.POST("/register", g.register)
	g.router = router

	return g
}

func (g *group) StartHTTP() {
	if err := g.router.Run(":8080"); err != nil {
		log.Fatalln("start http server err:", err)
	}
}
