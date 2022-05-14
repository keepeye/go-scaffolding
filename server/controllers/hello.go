package controllers

import "github.com/gin-gonic/gin"

type Hello struct {
	base
}

func (c *Hello) Setup(router gin.IRouter) {
	router.GET("/", c.Index)
}

func (c *Hello) Index(ctx *gin.Context) {
	ctx.String(200, "Hello World!!")
}
