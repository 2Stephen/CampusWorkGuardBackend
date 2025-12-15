package controller

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/model/response"
	"CampusWorkGuardBackend/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AdminLoginController(c *gin.Context) {
	var (
		params dto.AdminLoginRequest
	)
	if err := c.ShouldBind(&params); err != nil {
		log.Println("AdminLoginController ShouldBind error:", err)
		response.Fail(c, http.StatusBadRequest, "参数绑定失败")
		return
	}
	token, err := service.AdminLoginService(&params)
	if err != nil {
		log.Println("AdminLoginController AdminLoginService error:", err)
		response.Fail(c, http.StatusUnauthorized, err.Error())
		return
	}
	response.Success(c, gin.H{"token": token})
}

func AdminEmailLoginController(c *gin.Context) {
	var (
		params dto.AdminEmailLoginRequest
	)
	if err := c.ShouldBind(&params); err != nil {
		log.Println("AdminEmailLoginController ShouldBind error:", err)
		response.Fail(c, http.StatusBadRequest, "参数绑定失败")
		return
	}
	token, err := service.AdminEmailLoginService(&params)
	if err != nil {
		if err.Error() == "邮箱验证码已过期，请重新获取" || err.Error() == "邮箱验证码有误" {
			response.Fail(c, http.StatusBadRequest, err.Error())
		} else {
			log.Println("AdminEmailLoginController AdminEmailLoginService error:", err)
			response.Fail(c, http.StatusUnauthorized, err.Error())
		}
		return
	}
	response.Success(c, gin.H{"token": token})
}

func SetAdminPasswordController(c *gin.Context) {
	var (
		params dto.SetAdminPasswordRequest
	)
	userID, exists := c.Get("userID")
	if !exists {
		response.Fail(c, 401, "用户未认证")
		return
	}
	if err := c.ShouldBind(&params); err != nil {
		log.Println("SetAdminPasswordController ShouldBind error:", err)
		response.Fail(c, http.StatusBadRequest, "参数绑定失败")
		return
	}
	err := service.SetAdminPasswordService(&params, userID.(int))
	if err != nil {
		log.Println("SetAdminPasswordController SetAdminPasswordService error:", err)
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}
