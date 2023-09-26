package helpers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Entry

// Init logger initialize a logger instance by given name
func InitLogger(name string) {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter((&logrus.JSONFormatter{}))

	logger = log.WithFields(
		logrus.Fields{
			"program": "message_sender",
		},
	)
}

// Returns logger instance
func GetLogger() *logrus.Entry {
	if logger == nil {
		InitLogger("message_sender")
	}
	return logger
}

// Prints request data in JSON
func GinFormatMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		startTime := time.Now()

		//Process request
		ctx.Next()

		endTime := time.Now()

		latencyTime := endTime.Sub(startTime)

		reqMethod := ctx.Request.Method

		reqURI := ctx.Request.RequestURI

		statusCode := ctx.Writer.Status()

		clienIP := ctx.ClientIP()

		logger.WithFields(logrus.Fields{
			"status_code":    statusCode,
			"latency_time":   latencyTime,
			"client_ip":      clienIP,
			"request_method": reqMethod,
			"request_uri":    reqURI,
		}).Info()
	}
}

func CustomRecoveryWithWriter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.WithFields(logrus.Fields{
					"clien_ip":       ctx.ClientIP(),
					"request_method": ctx.Request.Method,
					"headers":        ctx.Request.Header,
					"url_params":     ctx.Request.URL.RawQuery,
					"panic_message":  err,
				}).Error()
			}
		}()
		ctx.Next()
	}
}
