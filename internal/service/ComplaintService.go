package service

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/model"
	"CampusWorkGuardBackend/internal/repository"
	"errors"
	"log"
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

func GetComplaintListService(params dto.GetComplaintListParams, userID int, role string) ([]model.ComplaintRecordList, int, error) {
	var complaints []model.ComplaintRecord
	var total int
	var err error
	// 根据用户角色调用不同的repository层查询逻辑
	if role == "student" {
		complaints, total, err = repository.GetComplaintRecordsByStudentID(params, userID)
	} else if role == "company" {
		complaints, total, err = repository.GetComplaintRecordsByCompanyID(params, userID)
	} else if role == "admin" {
		complaints, total, err = repository.GetAllComplaintRecords(params)
	} else {
		return nil, 0, errors.New("无效的用户角色")
	}
	var complaintList []model.ComplaintRecordList
	for _, complaint := range complaints {
		company, err := repository.GetCompanyByID(complaint.CompanyID)
		if err != nil {
			return nil, 0, err
		}
		if company == nil {
			return nil, 0, errors.New("未找到对应的公司信息")
		}
		list := model.ComplaintRecordList{
			Id:            complaint.ID,
			Title:         complaint.Title,
			Company:       company.Company,
			ComplaintDate: complaint.ComplaintDate,
			ComplaintType: complaint.ComplaintType,
			Status:        complaint.Status,
		}
		complaintList = append(complaintList, list)
	}
	return complaintList, total, err
}

func GetComplaintReplyService(complaintID int) (*model.ComplaintReply, error) {
	complaint, err := repository.GetComplaintRecordByID(complaintID)
	if err != nil {
		return nil, err
	}
	reply := &model.ComplaintReply{
		Id:             complaint.ID,
		CompanyDefense: complaint.CompanyDefense,
		ResultInfo:     complaint.ResultInfo,
	}
	return reply, nil
}

func ProcessComplaintService(params dto.CompanyProcessComplaint, userID int) error {
	complaintID := params.Id
	complaint, err := repository.GetComplaintRecordByID(complaintID)
	if err != nil {
		return err
	}
	if complaint == nil {
		return errors.New("未找到对应的投诉记录")
	}
	if complaint.CompanyID != userID {
		return errors.New("无权限处理该投诉记录")
	}
	if complaint.Status != "submitted" {
		return errors.New("企业已经处理过或管理员已经解决该投诉")
	}
	return repository.UpdateComplaintRecordCompanyDefense(complaintID, params.CompanyDefense)
}

func ResolveComplaintService(params dto.AdminResolveComplaint) error {
	complaintID := params.Id
	complaint, err := repository.GetComplaintRecordByID(complaintID)
	if err != nil {
		return err
	}
	if complaint == nil {
		return errors.New("未找到对应的投诉记录")
	}
	if complaint.Status != "processed" {
		return errors.New("只能处理企业用户答辩后的投诉记录")
	}
	return repository.UpdateComplaintRecordResultInfo(complaintID, params.ResultInfo)
}

func GetComplaintStatisticService(userId int, role string) (*model.ComplaintStatistics, error) {
	days := time.Now().AddDate(0, 0, -30).Format("2006-01-02")
	// 根据用户角色调用不同的repository层统计逻辑
	totalNums, err := repository.CountComplaintRecords(userId, role)
	if err != nil {
		log.Println("统计投诉总数失败:", err)
		return nil, err
	}
	thirtyDaysNewNums, err := repository.CountNewComplaintRecordsInLast30Days(userId, role, days)
	if err != nil {
		log.Println("统计近30天新增投诉数失败:", err)
		return nil, err
	}
	submittedNums, err := repository.CountComplaintRecordsByStatus(role, userId, "submitted")
	if err != nil {
		log.Println("统计submitted状态投诉数失败:", err)
		return nil, err
	}
	processedNums, err := repository.CountComplaintRecordsByStatus(role, userId, "processed")
	if err != nil {
		log.Println("统计processed状态投诉数失败:", err)
		return nil, err
	}
	resolvedNums, err := repository.CountComplaintRecordsByStatus(role, userId, "resolved")
	if err != nil {
		log.Println("统计resolved状态投诉数失败:", err)
		return nil, err
	}
	stats := &model.ComplaintStatistics{
		SubmittedNums:     int(submittedNums),
		ProcessedNums:     int(processedNums),
		ResolvedNums:      int(resolvedNums),
		TotalNums:         int(totalNums),
		ThirtyDaysNewNums: int(thirtyDaysNewNums),
	}
	return stats, nil
}
