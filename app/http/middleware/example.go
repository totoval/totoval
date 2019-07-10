package middleware

import (
	"github.com/totoval/framework/helpers/log"
	"github.com/totoval/framework/helpers/zone"
)

func Example() request.HandlerFunc {
	return func(c *request.Context) {
		t := zone.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := zone.Since(t)
		log.Info("latency", toto.V{"latency": latency})

		// access the status we are sending
		status := c.Writer.Status()
		log.Info("status", toto.V{"status": status})
	}
}
