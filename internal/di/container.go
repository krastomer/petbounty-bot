package di

import "github.com/gin-gonic/gin"

type Container struct {
	srv *gin.Engine
}

func (c *Container) Run() {
	c.srv.Run()
}
