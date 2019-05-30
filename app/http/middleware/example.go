package middleware

import (
	"github.com/totoval/framework/helpers/log"
	"github.com/totoval/framework/logs"

	"github.com/gin-gonic/gin"

	"github.com/totoval/framework/helpers/zone"
)

func Example() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := zone.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := zone.Since(t)
		log.Info("latency", logs.Field{"latency": latency})

		// access the status we are sending
		status := c.Writer.Status()
		log.Info("status", logs.Field{"status": status})
	}
}
