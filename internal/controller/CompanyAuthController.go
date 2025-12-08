package controller

import (
	"CampusWorkGuardBackend/internal/dto"
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

func AuthenticationCompanyController(ctx *gin.Context) {
	var req dto.CompanyRegisterRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.Fail(ctx, http.StatusBadRequest, "Invalid request data")
		return
	}
	token, err := service.RegisterCompanyService(&req)
	if err != nil {
		if err.Error() == "邮箱验证码已过期，请重新获取" || err.Error() == "邮箱验证码有误" {
			response.Fail(ctx, http.StatusBadRequest, err.Error())
			return
		} else {
			response.Fail(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}
	response.Success(ctx, gin.H{"token": token})
}
