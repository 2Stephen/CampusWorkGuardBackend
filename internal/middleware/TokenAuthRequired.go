package middleware

import (
	"CampusWorkGuardBackend/internal/model/response"
	"CampusWorkGuardBackend/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func TokenAuthRequired(c *gin.Context) {
	// 1. 获取 Authorization 字段
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		response.Fail(c, 401, "Authorization不能为空")
		c.Abort()
		return
	}

	// 必须是 Bearer 开头
	if !strings.HasPrefix(authHeader, "Bearer ") {
		response.Fail(c, 401, "Authorization格式错误，必须是Bearer开头")
		c.Abort()
		return
	}

	// 提取 token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	tokenString = strings.TrimSpace(tokenString)

	if tokenString == "" {
		response.Fail(c, 401, "Token不能为空")
		c.Abort()
		return
	}

	// 2. 解析 Token（内部自动验证密钥 & 过期）
	claims, err := utils.ParseJWTToken(tokenString)
	if err != nil {
		// 判断是否过期
		if utils.IsTokenExpired(err) {
			response.Fail(c, 401, "Token已过期，请重新登录")
			c.Abort()
			return
		}

		response.Fail(c, http.StatusUnauthorized, "无效的Token:"+err.Error())
		c.Abort()
		return
	}

	// 3. 把解析后的用户信息存入上下文，业务代码可直接使用
	c.Set("userID", claims.UserID)
	c.Set("email", claims.Email)

	c.Next()
}
