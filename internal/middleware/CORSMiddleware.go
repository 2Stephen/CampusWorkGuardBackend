package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*") // 或指定域名
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// 预检请求直接返回
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
