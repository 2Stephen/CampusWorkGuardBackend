package service

import (
	"CampusWorkGuardBackend/internal/dto"
	middlewares "CampusWorkGuardBackend/internal/middleware/AuthenticationModule"
	"CampusWorkGuardBackend/internal/model"
	"CampusWorkGuardBackend/internal/repository"
	"errors"
	"log"
)

func StudentAuth(params dto.StudentAuthParams) (*model.CHSIStudentInfo, error) {
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
		return cHSIStudentInfo, nil
	} else {
		log.Println("Student authentication failed for StudentID:", params.ID)
		return nil, errors.New("学信网验证内容与输入内容不符")
	}
}
