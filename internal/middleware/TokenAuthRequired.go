package middleware

import (
	"CampusWorkGuardBackend/internal/model/response"
	"github.com/gin-gonic/gin"
)

func TokenAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "invalid-token" {
			//c.AbortWithStatusJSON(401, gin.H{"error": "未授权访问"}) 源码为abort + json
			c.Abort()
			response.Fail(c, 401, "无效的令牌")
			return
		}
		//const mockUserId = 1
		//nickName, err := model.GetNicknameById(mockUserId) // 模拟获取用户昵称
		//if err != nil {
		//	c.Abort()
		//	response.Fail(c, 500, "获取用户信息失败")
		//	return
		//}
		//c.Set("nickname", nickName) // 设置用户信息到上下文中
		//c.Set("userId", mockUserId) // 设置用户信息到上下文中
		c.Next()
	}
}
