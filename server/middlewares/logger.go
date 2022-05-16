package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

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
