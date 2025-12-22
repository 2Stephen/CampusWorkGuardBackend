package model

type HomeStaticInfo struct {
	AvatarURL    string `json:"avatar_url"`
	Name         string `json:"name"`
	Role         string `json:"role"`
	VerifyStatus string `json:"verify_status"`
	FailInfo     string `json:"fail_info"`
}

type TopMajorJob struct {
	// const majorJobTop5Data = [
	//  {major: "计算机科学", value: 160},
	//  {major: "软件工程", value: 140},
	//  {major: "信息管理", value: 110},
	//  {major: "电子信息", value: 95},
	//  {major: "人工智能", value: 80},
	//];
	Major string `json:"major"`
	Value int    `json:"value"`
}

type JobType struct {
	Type  string `json:"type"`
	Value int    `json:"value"`
}

type AverageSalaryByMajor struct {
	Major string  `json:"major"`
	Value float64 `json:"value"`
}
