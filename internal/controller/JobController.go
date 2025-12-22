package controller

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/model"
	"CampusWorkGuardBackend/internal/model/response"
	"CampusWorkGuardBackend/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func GetCompanyUserJobInfoController(c *gin.Context) {
	id := c.Query("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.Fail(c, 400, "Invalid job ID")
		return
	}
	jobInfo, err := service.GetCompanyUserJobInfoService(idInt)
	if err != nil {
		response.Fail(c, 500, "Failed to get job info: "+err.Error())
		return
	}
	response.Success(c, jobInfo)
}

func PostJobController(c *gin.Context) {
	var (
		params dto.PostJobParams
	)
	if err := c.ShouldBind(&params); err != nil {
		response.Fail(c, 400, "Invalid request parameters")
	}
	userID, exists := c.Get("userID")
	if !exists {
		response.Fail(c, 401, "用户未认证")
		return
	}
	email, exists := c.Get("email")
	if !exists {
		response.Fail(c, 401, "用户未认证")
		return
	}
	err := service.PostJobService(params, userID.(int), email.(string))
	if err != nil {
		if err.Error() == "企业用户不存在" || err.Error() == "企业用户未通过认证，无法发布职位" || err.Error() == "用户邮箱与认证邮箱不匹配" {
			response.Fail(c, 403, err.Error())
			return
		}
		response.Fail(c, 500, "Failed to post job: "+err.Error())
		return
	}
	response.Success(c, nil)
}

func GetCompanyUserJobListController(c *gin.Context) {
	var (
		params dto.GetCompanyUserJobListParams
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
	email, exists := c.Get("email")
	if !exists {
		response.Fail(c, 401, "用户未认证")
		return
	}
	jobList, total, err := service.GetCompanyUserJobListService(userID.(int), email.(string), params)
	if err != nil {
		response.Fail(c, 500, "Failed to get job list: "+err.Error())
		return
	}
	response.Success(c, gin.H{
		"total": total,
		"jobs":  jobList,
	})
}

func UpdateJobController(c *gin.Context) {
	var (
		params dto.UpdateJobParams
	)
	if err := c.ShouldBind(&params); err != nil {
		response.Fail(c, 400, "Invalid request parameters")
	}
	userID, exists := c.Get("userID")
	if !exists {
		response.Fail(c, 401, "用户未认证")
		return
	}
	email, exists := c.Get("email")
	if !exists {
		response.Fail(c, 401, "用户未认证")
		return
	}
	err := service.UpdateJobService(params, userID.(int), email.(string))
	if err != nil {
		if err.Error() == "企业用户不存在" || err.Error() == "企业用户未通过认证，无法发布职位" || err.Error() == "用户邮箱与认证邮箱不匹配" || err.Error() == "职位不存在" || err.Error() == "无权限修改该职位信息" {
			response.Fail(c, 403, err.Error())
			return
		}
		response.Fail(c, 500, "Failed to post job: "+err.Error())
		return
	}
	response.Success(c, nil)
}

func DeleteJobController(c *gin.Context) {
	// 鉴权+删除
	id := c.Query("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.Fail(c, 400, "Invalid job ID")
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		response.Fail(c, 401, "用户未认证")
		return
	}
	email, exists := c.Get("email")
	if !exists {
		response.Fail(c, 401, "用户未认证")
		return
	}
	err = service.DeleteJobService(idInt, userID.(int), email.(string))
	if err != nil {
		if err.Error() == "企业用户不存在" || err.Error() == "无权限删除该职位信息" || err.Error() == "用户邮箱与认证邮箱不匹配" || err.Error() == "职位不存在" {
			response.Fail(c, 403, err.Error())
			return
		}
		response.Fail(c, 500, "Failed to delete job: "+err.Error())
		return
	}
	response.Success(c, nil)
}

func GetAdminJobListController(c *gin.Context) {
	var params dto.GetAdminJobListParams
	if err := c.ShouldBind(&params); err != nil {
		response.Fail(c, 400, "参数绑定错误")
		return
	}
	jobs, total, err := service.GetAdminJobListService(params)
	if err != nil {
		response.Fail(c, 500, "获取管理员职位列表失败")
		return
	}
	response.Success(c, gin.H{
		"jobs":  jobs,
		"total": total,
	})
}

func ReviewJobController(c *gin.Context) {
	var params dto.ReviewJobParams
	if err := c.ShouldBind(&params); err != nil {
		response.Fail(c, 400, "参数绑定错误")
		return
	}
	err := service.ReviewJobService(params)
	if err != nil {
		response.Fail(c, 500, "审核职位失败: "+err.Error())
		return
	}
	response.Success(c, nil)
}

func StudentUserJobMatchListController(c *gin.Context) {
	var params dto.StudentUserJobMatchListParams
	if err := c.ShouldBind(&params); err != nil {
		log.Println("Error binding parameters:", err)
		response.Fail(c, 400, "参数绑定错误")
		return
	}
	jobList, total, err := service.StudentUserJobMatchListService(params)
	if err != nil {
		response.Fail(c, 500, "获取职位匹配列表失败: "+err.Error())
		return
	}
	response.Success(c, gin.H{
		"total": total,
		"jobs":  jobList,
	})
}

func StudentUserApplyJobController(c *gin.Context) {
	id := c.Query("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.Fail(c, 400, "Invalid job ID")
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		response.Fail(c, 401, "用户未认证")
		return
	}
	err = service.StudentUserApplyJobService(userID.(int), idInt)
	if err != nil {
		if err.Error() == "职位不存在" || err.Error() == "您已申请该职位，不能重复申请" {
			response.Fail(c, 403, err.Error())
			return
		}
		response.Fail(c, 500, "申请职位失败: "+err.Error())
		return
	}
	response.Success(c, nil)
}

func GetJobApplicationListController(c *gin.Context) {
	var params dto.GetJobApplicationListParams
	if err := c.ShouldBind(&params); err != nil {
		response.Fail(c, 400, "参数绑定错误")
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		response.Fail(c, 401, "用户未认证")
		return
	}
	list, total, err := service.GetJobApplicationListService(userID.(int), params)
	if err != nil {
		response.Fail(c, 500, "获取职位申请列表失败: "+err.Error())
		return
	}
	response.Success(c, gin.H{
		"total":        total,
		"applications": list,
	})
}

func PayDepositController(c *gin.Context) {
	var params dto.PayDepositParams
	if err := c.ShouldBind(&params); err != nil {
		response.Fail(c, 400, "参数绑定错误")
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		response.Fail(c, 401, "用户未认证")
		return
	}
	err := service.PayDepositService(userID.(int), params)
	if err != nil {
		if err.Error() == "企业用户不存在" || err.Error() == "职位不存在" || err.Error() == "无权限为该职位支付押金" || err.Error() == "押金已支付，无需重复支付" {
			response.Fail(c, 403, err.Error())
			return
		}
		response.Fail(c, 500, "支付押金失败: "+err.Error())
		return
	}
	response.Success(c, nil)
}

func GetAdminJobApplicationListController(c *gin.Context) {
	var params dto.GetAdminJobApplicationListParams
	if err := c.ShouldBind(&params); err != nil {
		response.Fail(c, 400, "参数绑定错误")
		return
	}
	applications, total, err := service.GetAdminJobApplicationListService(params)
	if err != nil {
		response.Fail(c, 500, "获取管理员职位申请列表失败")
		return
	}
	response.Success(c, gin.H{
		"applications": applications,
		"total":        total,
	})
}

func GetStudentUserApplicationListController(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.Fail(c, 401, "用户未认证")
		return
	}
	res, err := service.GetStudentUserApplicationListService(userID.(int))
	if err != nil {
		response.Fail(c, 500, "获取学生用户申请列表失败: "+err.Error())
		return
	}
	response.Success(c, gin.H{
		"applications": []model.StudentUserApplicationDetail{res},
		"total":        1,
	})
}

func StudentUserAttendanceController(c *gin.Context) {
	var params dto.StudentUserAttendanceParams
	if err := c.ShouldBind(&params); err != nil {
		response.Fail(c, 400, "参数绑定错误")
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		response.Fail(c, 401, "用户未认证")
		return
	}
	err := service.StudentUserAttendanceService(userID.(int), params)
	if err != nil {
		if err.Error() == "未找到对应的工作申请" || err.Error() == "该职位未开始或已结束，无法进行考勤" || err.Error() == "今日已考勤，不能重复考勤" {
			response.Fail(c, 403, err.Error())
			return
		}
		response.Fail(c, 500, "学生用户考勤失败: "+err.Error())
		return
	}
	response.Success(c, nil)
}

func GetStudentUserAttendanceListController(c *gin.Context) {
	applicationJobID := c.Query("jobApplicationId")
	applicationJobIDInt, err := strconv.Atoi(applicationJobID)
	if err != nil {
		response.Fail(c, 400, "Invalid application job ID")
		return
	}
	records, err := service.GetStudentUserAttendanceListService(applicationJobIDInt)
	if err != nil {
		response.Fail(c, 500, "获取学生用户考勤列表失败: "+err.Error())
		return
	}
	response.Success(c, records)
}

func FinishJobController(c *gin.Context) {
	jobApplicationID := c.Query("jobApplicationId")
	jobApplicationIDInt, err := strconv.Atoi(jobApplicationID)
	if err != nil {
		response.Fail(c, 400, "Invalid job application ID")
		return
	}
	userId, exists := c.Get("userID")
	if !exists {
		response.Fail(c, 401, "用户未认证")
		return
	}
	err = service.FinishJobService(jobApplicationIDInt, userId.(int))
	if err != nil {
		if err.Error() == "职位申请不存在" || err.Error() == "无权限结束该职位" || err.Error() == "该职位申请未被录用，无法结束" || err.Error() == "该职位申请已结束" {
			response.Fail(c, 403, err.Error())
			return
		}
		response.Fail(c, 500, "结束职位失败: "+err.Error())
		return
	}
	response.Success(c, nil)
}
