package middleware

import "github.com/gin-gonic/gin"

func (mw *MiddlewareManager) Base() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
