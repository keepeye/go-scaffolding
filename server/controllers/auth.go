package controllers

import (
	"fmt"
	"myapp/server/constants"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Auth struct {
	base
}

func (c *Auth) PostLogin(ctx echo.Context) error {
	// 模拟用户登录生成jwt token
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "myapp",
		Subject:   fmt.Sprintf("%d", 1), // 这里存的是用户id
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString([]byte(constants.JWT_SECRET))
	return c.Succ(ctx, echo.Map{
		"token": tokenStr,
	})
}
