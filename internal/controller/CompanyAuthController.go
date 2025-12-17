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
		if err.Error() == "邮箱验证码已过期，请重新获取" || err.Error() == "邮箱验证码有误" || err.Error() == "重复注册" {
			response.Fail(ctx, http.StatusBadRequest, err.Error())
			return
		} else {
			response.Fail(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}
	response.Success(ctx, gin.H{
		"token": token,
		"role":  "company",
	})
}

func CompanyEmailLoginController(c *gin.Context) {
	var req dto.CompanyEmailLoginRequest
	if err := c.ShouldBind(&req); err != nil {
		response.Fail(c, http.StatusBadRequest, "Invalid request data")
		return
	}
	// 调用service进行登录逻辑处理
	token, err := service.CompanyEmailLogin(req)
	if err != nil {
		if err.Error() == "邮箱验证码有误" || err.Error() == "邮箱验证码已过期，请重新获取" {
			response.Fail(c, 403, err.Error())
			return
		}
		response.Fail(c, 500, "Failed to login: "+err.Error())
		return
	}
	response.Success(c, gin.H{
		"token": token,
		"role":  "company",
	})
}

func CompanyLoginController(ctx *gin.Context) {
	var req dto.CompanyLoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.Fail(ctx, http.StatusBadRequest, "Invalid request data")
		return
	}
	// 调用service进行登录逻辑处理
	token, err := service.CompanyLoginService(&req)
	if err != nil {
		if err.Error() == "用户登录失败，检查邮箱或密码是否正确" || err.Error() == "用户未设置密码，请使用邮箱验证登录后设置密码" {
			response.Fail(ctx, http.StatusForbidden, err.Error())
			return
		}
		response.Fail(ctx, http.StatusInternalServerError, "Failed to login: "+err.Error())
		return
	}
	response.Success(ctx, gin.H{
		"token": token,
		"role":  "company",
	})

}
