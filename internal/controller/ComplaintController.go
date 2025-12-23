package controller

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/model/response"
	"CampusWorkGuardBackend/internal/service"
	"github.com/gin-gonic/gin"
	"log"
)

func SubmitComplaintController(c *gin.Context) {
	var (
		params dto.SubmitComplaintParams
	)
	if err := c.ShouldBind(&params); err != nil {
		response.Fail(c, 400, "Invalid request parameters")
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		response.Fail(c, 401, "用户未认证")
		return
	}
	err := service.SubmitComplaintService(params, userID.(int))
	if err != nil {
		log.Println("提交投诉失败:", err)
		response.Fail(c, 500, "Failed to submit complaint: "+err.Error())
		return
	}
	response.Success(c, nil)
}
