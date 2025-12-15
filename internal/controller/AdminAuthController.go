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
