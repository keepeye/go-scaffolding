package controllers

import "github.com/labstack/echo/v4"

type Hello struct {
	base
}

func (c *Hello) Get(ctx echo.Context) error {
	return ctx.String(200, "Hello World!!")
}
