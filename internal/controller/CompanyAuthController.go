package controller

import (
	"CampusWorkGuardBackend/internal/model/response"
	"CampusWorkGuardBackend/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func UploadLicenseController(ctx *gin.Context) {
	// 存储营业执照图片
	file, err := ctx.FormFile("file")
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, "No file is received")
		return
	}

	// 目录自动创建
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		err := os.Mkdir("uploads", os.ModePerm)
		if err != nil {
			response.Fail(ctx, http.StatusInternalServerError, "Could not create upload directory")
			return
		}
	}
	filename, err := service.SaveImage(file)
	if err != nil {
		if err.Error() == "不支持的文件类型" || err.Error() == "文件大小超过5MB限制" {
			response.Fail(ctx, http.StatusBadRequest, err.Error())
		} else {
			response.Fail(ctx, http.StatusInternalServerError, "Failed to save file")
		}
		return
	}
	// 返回可访问的 URL
	url := "/uploads/" + filename
	response.Success(ctx, gin.H{"url": url})
}
