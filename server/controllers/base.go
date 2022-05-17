package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

type base struct {
}

func (c *base) Fail(ctx echo.Context, code int, message string) error {
	return ctx.JSON(200, gin.H{
		"code":    code,
		"message": message,
	})
}

func (c *base) Succ(ctx echo.Context, data interface{}) error {
	return ctx.JSON(200, gin.H{
		"code": 0,
		"data": data,
	})
}
