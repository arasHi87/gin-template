package middleware

import (
	"bytes"
	"time"

	"github.com/arashi87/gin-template/pkg/common"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// use to instead origin response writer
// so need to have Write and WriteString to implement origin response writer interface
type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// write bytes into buffer
func (w responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// write string into buffer
func (w responseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		writer := &responseWriter{ResponseWriter: ctx.Writer, body: bytes.NewBufferString("")}

		ctx.Writer = writer
		ctx.Next()

		common.Logger.WithFields(logrus.Fields{
			"latency_time":   time.Since(startTime),
			"request_method": ctx.Request.Method,
			"request_uri":    ctx.Request.RequestURI,
			"status_code":    ctx.Writer.Status(),
			"client_ip":      ctx.ClientIP(),
			"response":       writer.body.String(),
		}).Info("request body")
	}
}
