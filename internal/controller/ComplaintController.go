package controller

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/model/response"
	"CampusWorkGuardBackend/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
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

func DeleteComplaintController(c *gin.Context) {
	id := c.Query("id")
	log.Println("Received complaint ID to delete:", id)
	complaintID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Invalid complaint ID:", err)
		response.Fail(c, 400, "Invalid complaint ID")
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		response.Fail(c, 401, "用户未认证")
		return
	}
	err = service.DeleteComplaintService(complaintID, userID.(int))
	if err != nil {
		log.Println("删除投诉失败:", err)
		if err.Error() == "无权限删除该投诉记录" {
			response.Fail(c, 403, err.Error())
			return
		}
		response.Fail(c, 500, "Failed to delete complaint: "+err.Error())
		return
	}
	response.Success(c, nil)
}

func GetComplaintListController(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.Fail(c, 401, "用户未认证")
		return
	}
	role, exists := c.Get("role")
	if !exists {
		response.Fail(c, 401, "用户角色未识别")
		return
	}
	var params dto.GetComplaintListParams
	if err := c.ShouldBind(&params); err != nil {
		response.Fail(c, 400, "Invalid request parameters")
		return
	}
	complaints, total, err := service.GetComplaintListService(params, userID.(int), role.(string))
	if err != nil {
		log.Println("获取投诉列表失败:", err)
		response.Fail(c, 500, "Failed to get complaint list: "+err.Error())
		return
	}
	response.Success(c, gin.H{
		"complaints": complaints,
		"total":      total,
	})
}

func GetComplaintReplyController(c *gin.Context) {
	id := c.Param("id")
	complaintID, err := strconv.Atoi(id)
	if err != nil {
		response.Fail(c, 400, "Invalid complaint ID")
		return
	}
	complaint, err := service.GetComplaintReplyService(complaintID)
	if err != nil {
		log.Println("获取投诉回复失败:", err)
		response.Fail(c, 500, "Failed to get complaint reply: "+err.Error())
		return
	}
	response.Success(c, complaint)
}

func ProcessComplaintController(c *gin.Context) {
	var (
		params dto.CompanyProcessComplaint
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
	err := service.ProcessComplaintService(params, userID.(int))
	if err != nil {
		if err.Error() == "无权限处理该投诉" {
			response.Fail(c, 403, err.Error())
			return
		}
		if err.Error() == "无效的投诉ID" {
			response.Fail(c, 400, err.Error())
			return
		}
		if err.Error() == "未找到对应的投诉记录" {
			response.Fail(c, 404, err.Error())
			return
		}
		log.Println("处理投诉失败:", err)
		response.Fail(c, 500, "Failed to process complaint: "+err.Error())
		return
	}
	response.Success(c, nil)
}
