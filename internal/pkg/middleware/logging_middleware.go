package middleware

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now() // starting time request

		ctx.Next() // process request

		endTime := time.Now()                 // end time request
		latencyTime := endTime.Sub(startTime) // calculate esecution time
		requestMethod := ctx.Request.Method   // request method
		reqURI := ctx.Request.RequestURI      // request route
		statusCode := ctx.Writer.Status()     // status code
		clientIP := ctx.ClientIP()            // client ip

		log.WithFields(log.Fields{
			"latency_time":   latencyTime,
			"request_method": requestMethod,
			"req_uri":        reqURI,
			"status_code":    statusCode,
			"client_ip":      clientIP,
		}).Info("http request")

		ctx.Next()
	}
}
