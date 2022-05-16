package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

func Auth(ctx *gin.Context) {
	jwtClaims, exists := ctx.Get("jwt-claims")
	if !exists {
		log.Warnf("Authentication failed, pre JWT middleware is required")
		ctx.AbortWithStatus(401)
		return
	}
	claims := jwtClaims.(jwt.MapClaims)
	userID := claims["sub"]
	// 读取user信息，并存储到ctx中
	// user := models.GetUserByID(userId)
	// ctx.Set("user", user)
	ctx.Set("userId", userID)
}
