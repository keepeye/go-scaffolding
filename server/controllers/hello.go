package controllers

import (
	"github.com/labstack/echo/v4"
)

type Hello struct {
	base
}

func (c *Hello) Get(ctx echo.Context) error {
	return ctx.Render(200, "hello_index", echo.Map{
		"name": "cheng",
	})
}
