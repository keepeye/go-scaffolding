package controllers

import (
	"github.com/gin-gonic/gin"
)

type Hello struct {
	base
}

func (c *Hello) Get(ctx *gin.Context) {
	ctx.String(200, "Hello World!!")
}
