package middleware

import (
	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func ApiKeyMiddleware(expectedToken string) gin.HandlerFunc {
	return func(context *gin.Context) {
		apiToken := context.GetHeader("X-API-Key")
		if apiToken == "" {
			respondWithError(context, 401, "API token required")
			return
		}
		if apiToken != expectedToken {
			respondWithError(context, 401, "Invalid API token")
		}

		context.Next()
	}
}
