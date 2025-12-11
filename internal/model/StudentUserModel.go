package model

type StudentUser struct {
	Id        int    `gorm:"column:id;primaryKey"`
	SchoolId  int    `gorm:"column:school_id"`
	StudentId string `gorm:"column:student_id"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
	AvatarURL string `gorm:"column:avatar_url"`
}
