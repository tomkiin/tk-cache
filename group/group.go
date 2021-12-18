package group

import (
	"log"
	"sync"
	"tk-cache/pkg/httpx"

	"github.com/gin-gonic/gin"
)

type group struct {
	mu     *sync.RWMutex
	nodes  map[string]bool
	router *gin.Engine
}

func New() *group {
	g := &group{
		mu:    new(sync.RWMutex),
		nodes: make(map[string]bool),
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

func (g *group) register(c *gin.Context) {
	ip := c.PostForm("ip")

	g.mu.Lock()
	defer g.mu.Unlock()

	g.nodes[ip] = true

	httpx.RenderOK(c)
}
