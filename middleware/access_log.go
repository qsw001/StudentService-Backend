package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		method := c.Request.Method
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		ip := c.ClientIP()

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		size := c.Writer.Size() 

		if query != "" {
			path = path + "?" + query
		}

		log.Printf("[ACCESS] ip=%s | %3d | %s | %-7s %s | %dB | %v",
			ip, status, c.Request.Proto, method, path, size, latency,
		)
	}
}