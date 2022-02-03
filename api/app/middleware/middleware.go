package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("\n%s %s %s", c.Request.Method, c.Request.RequestURI, c.Request.Host)
		c.Next()
	}
}
