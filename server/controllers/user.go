package controllers

import "github.com/gin-gonic/gin"

type User struct {
	base
}

func (c *User) GetProfile(ctx *gin.Context) {}
