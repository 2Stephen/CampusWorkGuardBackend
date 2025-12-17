package service

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/model"
	"CampusWorkGuardBackend/internal/repository"
)

func GetAdminCompanyListService(params *dto.CompanyListRequest) ([]model.CompanyList, int64, error) {
	companyUsers, total, err := repository.GetAdminCompanyList(params.Page, params.PageSize, params.Search, params.Status)
	var companyList []model.CompanyList
	if err != nil {
		return nil, 0, err
	}
	for _, cu := range companyUsers {
		cl := model.CompanyList{
			ID:           int(cu.ID),
			Name:         cu.Name,
			Company:      cu.Company,
			Email:        cu.Email,
			SocialCode:   cu.SocialCode,
			LicenseUrl:   cu.LicenseURL,
			VerifyStatus: cu.VerifyStatus,
		}
		companyList = append(companyList, cl)
	}
	return companyList, total, nil
}
