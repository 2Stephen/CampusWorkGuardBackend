package service

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/model"
	"CampusWorkGuardBackend/internal/repository"
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
