package middleware

import (
	"time"

	"github.com/arashi87/gin-template/pkg/common"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		common.Logger.WithFields(logrus.Fields{
			"latency_time":   time.Since(startTime),
			"request_method": ctx.Request.Method,
			"request_uri":    ctx.Request.RequestURI,
			"status_code":    ctx.Writer.Status(),
			"client_ip":      ctx.ClientIP(),
		}).Info("request body")
	}
}
