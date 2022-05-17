package middlewares

import (
	"myapp/server/constants"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Auth() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(constants.JWT_SECRET),
		Skipper:       middleware.DefaultSkipper,
		SigningMethod: middleware.AlgorithmHS256,
		ContextKey:    "jwt",
		TokenLookup:   "header:Authorization,query:token",
		AuthScheme:    "Bearer",
		Claims:        &jwt.StandardClaims{},
		SuccessHandler: func(ctx echo.Context) {
			jwtToken := ctx.Get("jwt").(*jwt.Token)
			claims := jwtToken.Claims.(*jwt.StandardClaims)
			userID := claims.Subject
			// 读取user信息，并存储到ctx中
			// user := models.GetUserByID(userId)
			// ctx.Set("user", user)
			ctx.Set("userId", userID)
		},
		ErrorHandlerWithContext: func(err error, c echo.Context) error {
			return echo.ErrUnauthorized
		},
	})
}
