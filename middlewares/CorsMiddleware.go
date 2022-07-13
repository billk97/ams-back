package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "*, Authorization")
		c.Writer.Header().Set("Referrer-Policy", "no-referrer-when-downgrade")
		c.Writer.Header().Set("Strict-Transport-Security", "max-age=63072000")
		c.Writer.Header().Set("skata2", "skata2")
		fmt.Println("edooo")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
