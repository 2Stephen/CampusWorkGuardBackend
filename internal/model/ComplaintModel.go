package model

type ComplaintRecord struct {
	ID             int    `gorm:"column:id;type:int;primaryKey;autoIncrement;comment:主键id" json:"id"`
	StudentID      int    `gorm:"column:student_id;type:int;not null;comment:投诉学生ID" json:"student_id"`
	CompanyID      int    `gorm:"column:company_id;type:int;not null;comment:投诉企业ID" json:"company_id"`
	ComplaintDate  string `gorm:"column:complaint_date;type:varchar(64);not null;comment:投诉发起日期" json:"complaint_date"`
	Title          string `gorm:"column:title;type:varchar(255);not null;comment:投诉标题" json:"title"`
	ComplaintType  string `gorm:"column:complaint_type;type:varchar(32);not null;comment:投诉类型" json:"complaint_type"`
	CompanyDefense string `gorm:"column:company_defense;type:text;default:null;comment:企业答辩内容" json:"company_defense"`
	Status         string `gorm:"column:status;type:varchar(32);not null;comment:投诉状态（submitted/processed/resolved）" json:"status"`
	ResultInfo     string `gorm:"column:result_info;type:varchar(500);default:null;comment:申诉结果信息（管理员填写）" json:"result_info"`
}
