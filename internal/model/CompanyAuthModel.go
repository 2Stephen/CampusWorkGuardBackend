package model

type CompanyUser struct {
	ID int64 `gorm:"column:id;primaryKey;autoIncrement;comment:主键ID"`

	Name    string `gorm:"column:name;type:varchar(50);comment:注册人姓名"`
	Email   string `gorm:"column:email;type:varchar(100);uniqueIndex:uk_email;comment:注册人邮箱（公司邮箱）"`
	Company string `gorm:"column:company;type:varchar(100);comment:公司名称（需与营业执照一致）"`

	LicenseURL string `gorm:"column:license_url;type:varchar(255);comment:营业执照相对URL"`
	SocialCode string `gorm:"column:social_code;type:varchar(32);uniqueIndex:uk_social_code;comment:统一社会信用代码"`

	Password  string `gorm:"column:password;type:varchar(255);default:null;comment:密码（初始为空，可后续修改）"`
	AvatarURL string `gorm:"column:avatar_url;type:varchar(255);default:null;comment:头像URL（默认头像）"`

	VerifyStatus string `gorm:"column:verify_status;type:varchar(20);not null;default:'验证中';comment:验证状态：验证中/验证成功/验证失败"`

	FailInfo string `gorm:"column:fail_info;type:varchar(255);default:null;comment:上次验证失败信息"`
}
