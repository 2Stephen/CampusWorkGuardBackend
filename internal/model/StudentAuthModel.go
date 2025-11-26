package model

type CHSIStudentInfo struct {
	Id           int    `gorm:"column:id;primaryKey"`
	Name         string `gorm:"column:name"`
	Gender       string `gorm:"column:gender"`
	Birthday     string `gorm:"column:birthday"`
	Nation       string `gorm:"column:nation"`
	School       string `gorm:"column:school"`
	Level        string `gorm:"column:level"`
	Major        string `gorm:"column:major"`
	Duration     string `gorm:"column:duration"`
	DegreeType   string `gorm:"column:degree_type"`
	StudyMode    string `gorm:"column:study_mode"`
	College      string `gorm:"column:college"`
	Department   string `gorm:"column:department"`
	EntranceDate string `gorm:"column:entrance_date"`
	Status       string `gorm:"column:status"`
	ExpectedGrad string `gorm:"column:expected_grad"`
	Vcode        string `gorm:"column:vcode"`
	StudentID    string `gorm:"column:student_id"`
	Email        string `gorm:"column:email"`
}
