package schemas

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID       uint
	Username string
	Password string
	Email    string
}

type Session struct {
	gorm.Model

	ID        uint
	UID       uint
	IP        string
	UserAgent string
	User      User `gorm:"foreignKey:UID"`
}
