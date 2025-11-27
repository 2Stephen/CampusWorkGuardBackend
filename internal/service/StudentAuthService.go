package service

import (
	"CampusWorkGuardBackend/internal/dto"
	middlewares "CampusWorkGuardBackend/internal/middleware"
	"CampusWorkGuardBackend/internal/model"
	"CampusWorkGuardBackend/internal/repository"
	"CampusWorkGuardBackend/internal/utils"
	"errors"
	"github.com/redis/go-redis/v9"
	"log"
)

func GetSchoolList(search string) ([]model.School, error) {
	// 调用repository层的方法从数据库获取学校列表
	filteredSchools, err := repository.GetSchoolList(search)
	return filteredSchools, err
}

func StudentAuth(params dto.StudentAuthParams) (*model.CHSIStudentInfo, error) {
	// 验证邮箱验证码
	realCode, err := utils.RedisGet("register_code:" + params.Email)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, errors.New("邮箱验证码已过期，请重新获取")
		}
		return nil, errors.New("获取邮箱验证码失败")
	}
	if realCode != params.Code {
		return nil, errors.New("邮箱验证码有误")
	} else {
		// 清除redis缓存验证码
		err := utils.RedisDel("register_code:" + params.Email)
		if err != nil {
			log.Println("删除邮箱验证码缓存失败", err)
		}
	}
	// http服务调用chrome，chrome打开学信网接口进行认证
	log.Println("Starting student authentication for ID:", params.ID)

	cHSIStudentInfo, err := middlewares.CHSIAuth(params.Vcode)
	if err != nil {
		return nil, err
	}
	// 对比信息
	// 输出获取到的学生信息
	log.Printf("Retrieved Student Info: %+v\n", cHSIStudentInfo)
	if cHSIStudentInfo.School == params.School && cHSIStudentInfo.Vcode == params.Vcode {
		log.Println("Student authentication successful for StudentID:", params.ID)
		cHSIStudentInfo.StudentID = params.ID
		cHSIStudentInfo.Email = params.Email
		// 入库
		if err := repository.CreateCHSIStudentInfo(cHSIStudentInfo); err != nil {
			return nil, errors.New("数据库保存学生信息失败")
		}
		// 查询school ID
		schoolId, err := repository.GetSchoolId(params.School)
		if err != nil {
			return nil, errors.New("学校信息比对失败")
		}
		studentUser := model.StudentUser{
			SchoolId:  schoolId,
			StudentId: params.ID,
			Email:     params.Email,
			Password:  "",
		}
		if err := repository.CreateStudentUser(studentUser); err != nil {
			return nil, errors.New("数据库保存学生信息失败")
		}
		return cHSIStudentInfo, nil
	} else {
		log.Println("Student authentication failed for StudentID:", params.ID)
		return nil, errors.New("学信网验证内容与输入内容不符")
	}
}

func StudentLogin(params dto.StudentLoginParams) (string, error) {
	user := repository.GetSchoolUser(params.SchoolId, params.StudentId)
	if user == nil {
		return "", errors.New("学生登录失败，检查学号或密码是否正确")
	}
	// 哈希验证
	ok, err := utils.VerifyPassword(params.Password, user.Password)
	if err != nil {
		return "", errors.New("学生登录失败" + err.Error())
	}
	if !ok {
		return "", errors.New("学生登录失败，检查学号或密码是否正确")
	}
	// 生成JWT token
	token, err := utils.GenerateJWTToken(user.Id, user.Email)
	if err != nil {
		return "", errors.New("生成登录令牌失败")
	}
	return token, nil
}
