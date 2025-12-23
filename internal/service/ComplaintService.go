package service

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/model"
	"CampusWorkGuardBackend/internal/repository"
	"errors"
	"time"
)

func SubmitComplaintService(params dto.SubmitComplaintParams, userID int) error {
	// 调用repository层保存投诉逻辑
	complaint := model.ComplaintRecord{
		StudentID:      userID,
		CompanyID:      params.CompanyId,
		ComplaintDate:  time.Now().Format("2006-01-02"),
		Title:          params.Title,
		ComplaintType:  params.ComplaintType,
		CompanyDefense: "",
		Status:         "submitted",
		ResultInfo:     "",
	}
	return repository.SaveComplaintRecord(&complaint)
}

func DeleteComplaintService(complaintID int, userID int) error {
	// 判断用户是否有权限删除该投诉记录
	complaint, err := repository.GetComplaintRecordByID(complaintID)
	if err != nil {
		return err
	}
	if complaint.StudentID != userID {
		return errors.New("无权限删除该投诉记录")
	}
	return repository.DeleteComplaintRecord(complaintID)
}
