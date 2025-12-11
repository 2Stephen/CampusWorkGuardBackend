package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`    // 状态码，0 表示成功，其他表示失败
	Message string      `json:"message"` // 描述信息
	Data    interface{} `json:"data"`    // 返回的数据内容
}

// Success 返回成功的响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "",
		Data:    data,
	})
}

// Fail 返回失败的响应
func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}
