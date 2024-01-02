package network

import "github.com/gin-gonic/gin"

func (n *Network) registerCreate(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.POST(path, handler...)
}

func (n *Network) registerGet(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.GET(path, handler...)
}

func (n *Network) registerUpdate(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.PUT(path, handler...)
}

func (n *Network) registerDelete(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.DELETE(path, handler...)
}

func (n *Network) okResponse(c *gin.Context, result interface{}) {
	c.JSON(200, result)
}

func (n *Network) failedResponse(c *gin.Context, result interface{}) {
	c.JSON(400, result)
}
