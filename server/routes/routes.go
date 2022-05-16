package routes

import (
	"myapp/server/constants"
	"myapp/server/controllers"
	"myapp/server/middlewares"

	"github.com/gin-gonic/gin"
)

// 控制器接口
type Controller interface {
	Setup(gin.IRouter)
}

// 加载控制器并调用控制器的Setup注册子路由
func registerController(router gin.IRouter, controller Controller) {
	controller.Setup(router)
}

// 带登录验证的路由
func guard(router gin.IRouter, path string) gin.IRouter {
	return router.Group(path, middlewares.JWT(constants.JWT_SECRET), middlewares.Auth)
}

// 在这里注册所有的路由
func Setup(router gin.IRouter) {
	registerController(router, new(controllers.Hello))
	registerController(guard(router, "/user"), &controllers.User{})
}
