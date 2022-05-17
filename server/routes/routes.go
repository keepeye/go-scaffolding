package routes

import (
	"myapp/core/boost"
	"myapp/server/controllers"
	"myapp/server/middlewares"

	"github.com/labstack/echo/v4"
)

// 带登录验证的路由
func guard(router *echo.Echo, path string) *echo.Group {
	return router.Group(path, middlewares.Auth())
}

// 在这里注册所有的路由
func Setup(router *echo.Echo) {
	boost.RegisterController(router.Group("/"), new(controllers.Hello))
	boost.RegisterController(router.Group("/auth"), new(controllers.Auth))
	boost.RegisterController(guard(router, "/user"), new(controllers.User))
}
