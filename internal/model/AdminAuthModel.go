package model

type AdminUser struct {
	ID        int     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string  `gorm:"type:varchar(50);not null;unique" json:"name"`
	Email     string  `gorm:"type:varchar(100);not null;unique" json:"email"`
	Password  string  `gorm:"type:varchar(255);not null" json:"password"`
	AvatarURL *string `gorm:"type:varchar(255)" json:"avatar_url"`
}
