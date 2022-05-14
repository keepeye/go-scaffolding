package controllers

import "github.com/gin-gonic/gin"

type base struct {
}

func (c *base) Fail(ctx *gin.Context, code int, message string) {
	ctx.JSON(200, gin.H{
		"code":    code,
		"message": message,
	})
}

func (c *base) Succ(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, gin.H{
		"code": 0,
		"data": data,
	})
}
