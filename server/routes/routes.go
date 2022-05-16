package routes

import (
	"myapp/core/boost"
	"myapp/server/constants"
	"myapp/server/controllers"
	"myapp/server/middlewares"

	"github.com/gin-gonic/gin"
)

// 带登录验证的路由
func guard(router gin.IRouter, path string) gin.IRouter {
	return router.Group(path, boost.JWTValidator(constants.JWT_SECRET), middlewares.Auth)
}

// 在这里注册所有的路由
func Setup(router gin.IRouter) {
	boost.RegisterController(router.Group("/"), new(controllers.Hello))
	boost.RegisterController(guard(router, "/user"), new(controllers.User))
}
