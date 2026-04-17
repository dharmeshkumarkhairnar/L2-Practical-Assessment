package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		log.Printf("Request: %s %s", ctx.Request.Method, ctx.Request.URL.Path)

		ctx.Next()

		duration := time.Since(start)
		log.Printf("Completed in %v", duration)
	}
}
