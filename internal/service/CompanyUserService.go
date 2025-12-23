package service

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/repository"
	"CampusWorkGuardBackend/internal/utils"
	"errors"
	"log"
)

type CompanyProfileInfo struct {
	AvatarURL string `json:"avatar_url"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Company   string `json:"company"`
}

func SetCompanyUserPassword(params dto.SetCompanyUserPasswordParams, userID string) error {
	// 校验密码长度、复杂度等
	if len(params.Password) < 8 {
		return errors.New("密码长度不足，至少需要8位")
	}
	if !containsNumber(params.Password) || !containsLetter(params.Password) {
		return errors.New("密码必须包含字母和数字")
	}
	if len(params.Password) > 64 {
		return errors.New("密码长度过长，不能超过64位")
	}
	hashedPassword, err := utils.HashPassword(params.Password)
	if err != nil {
		return errors.New("密码加密失败")
	}
	// 调用repository层保存密码逻辑
	return repository.SaveCompanyUserPassword(hashedPassword, userID)
}

func DeleteCompanyUserService(id int) error {
	return repository.DeleteCompanyUserByID(int64(id))
}

func GetCompanyUserProfileInfoService(userID int) (*CompanyProfileInfo, error) {
	user := repository.GetCompanyUserByID(int64(userID))
	if user == nil {
		log.Println("用户不存在，ID:", userID)
		return nil, errors.New("用户不存在")
	}
	profileInfo := &CompanyProfileInfo{
		AvatarURL: user.AvatarURL,
		Email:     user.Email,
		Name:      user.Name,
		Company:   user.Company,
	}
	return profileInfo, nil
}

func GetCompanyListService(search string) ([]dto.CompanyInfo, error) {
	companies, err := repository.GetAllCompanies(search)
	if err != nil {
		return nil, err
	}
	var ans []dto.CompanyInfo
	for _, company := range companies {
		ans = append(ans, dto.CompanyInfo{
			ID:   int(company.ID),
			Name: company.Company,
		})
	}
	return ans, nil
}
