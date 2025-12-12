package model

type HomeStaticInfo struct {
	AvatarURL    string `json:"avatar_url"`
	Name         string `json:"name"`
	Role         string `json:"role"`
	VerifyStatus string `json:"verify_status"`
	FailInfo     string `json:"fail_info"`
}
