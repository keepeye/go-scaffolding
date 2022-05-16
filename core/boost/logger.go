package boost

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// CustomLogger 自定义gin的logger，输出重定向到logrus.Writer
func CustomLogger() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("%s %s %d %s | %s | %s | %s\n",
				param.Method,
				param.Path,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ClientIP,
				param.ErrorMessage,
			)
		},
		Output: logrus.StandardLogger().WriterLevel(logrus.DebugLevel),
	})
}
