package network

import "github.com/gin-gonic/gin"

type Network struct {
	engine *gin.Engine
}

func NewNetwork() *Network {
	r := &Network{
		engine: gin.New(),
	}

	newForumRouter(r)

	return r
}

func (n *Network) ServerStart(port string) error {
	return n.engine.Run(port)
}
