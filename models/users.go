package models

import "github.com/jinzhu/gorm"

// Users 用户基础模型
type Users struct {
	gorm.Model
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (u *Users) AddTest() (rs bool) {
	u.Username = "admin"
	u.Password = "password"
	db.Create(u)
	// rs = db.NewRecord(user)
	// db.Create(&Product{Code: "L1212", Price: 1000})
	return true
}
