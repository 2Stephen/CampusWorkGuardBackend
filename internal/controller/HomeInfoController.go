package controller

import (
	"CampusWorkGuardBackend/internal/model/response"
	"CampusWorkGuardBackend/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func GetHomeStaticInfoController(ctx *gin.Context) {
	userId, exists := ctx.Get("userID")
	if !exists {
		response.Fail(ctx, 401, "用户未登录")
		return
	}
	role, exists := ctx.Get("role")
	if !exists {
		response.Fail(ctx, 401, "用户未登录")
		return
	}
	email, exists := ctx.Get("email")
	if !exists {
		response.Fail(ctx, 401, "用户未登录")
		return
	}

	info, err := service.GetHomeStaticInfo(userId.(int), role.(string), email.(string))
	if err != nil {
		response.Fail(ctx, 500, "获取个人基本信息失败: "+err.Error())
		return
	}
	response.Success(ctx, info)
}

func UploadAvatarController(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Fail(c, http.StatusBadRequest, "No file is received")
		return
	}
	userId, exists := c.Get("userID")
	if !exists {
		response.Fail(c, 401, "用户未登录")
		return
	}
	role, exists := c.Get("role")
	if !exists {
		response.Fail(c, 401, "用户未登录")
		return
	}
	// 目录自动创建
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		err := os.Mkdir("uploads", os.ModePerm)
		if err != nil {
			log.Println("Could not create upload directory:", err)
			response.Fail(c, http.StatusInternalServerError, "Could not create upload directory")
			return
		}
	}
	filename, err := service.SaveImage(file)
	if err != nil {
		if err.Error() == "不支持的文件类型" || err.Error() == "文件大小超过5MB限制" {
			response.Fail(c, http.StatusBadRequest, err.Error())
		} else {
			log.Println("Failed to save file:", err)
			response.Fail(c, http.StatusInternalServerError, "Failed to save file")
		}
		return
	}
	// 返回可访问的 URL
	url := "/uploads/" + filename
	err = service.UploadAvatarService(url, userId.(int), role.(string))
	if err != nil {
		log.Println("上传头像失败:", err)
		response.Fail(c, 500, "上传头像失败: "+err.Error())
		return
	}
	response.Success(c, gin.H{"url": url})
}
