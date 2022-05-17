package controllers

import "github.com/labstack/echo/v4"

type User struct {
	base
}

func (c *User) GetProfile(ctx echo.Context) error {
	return ctx.String(200, ctx.Get("userId").(string))
}
