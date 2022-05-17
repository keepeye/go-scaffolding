package boost

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

// CustomLogger 自定义logger，输出重定向到logrus.Writer
func CustomLogger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${host}${uri} - ${status} - ${remote_ip} \n", //末尾不能少了换行，否则无法输出
		Output: logrus.StandardLogger().WriterLevel(logrus.DebugLevel),
	})
}
