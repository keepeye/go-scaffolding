package controllers

import "github.com/gin-gonic/gin"

type User struct {
	base
}

func (c *User) Setup(router gin.IRouter) {
	router.GET("/profile", c.Profile)
}

func (c *User) Profile(ctx *gin.Context) {}
