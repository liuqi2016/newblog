package models

import "github.com/jinzhu/gorm"

// Users 用户基础模型
type Users struct {
	gorm.Model
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

//AddTest 添加测试数据
func (u *Users) AddTest() (rs bool) {
	user := Users{Username: "Jinzhu", Password: "1234561"}
	db.Create(&user)
	if user.ID < 1 {
		return false
	}
	// rs = db.NewRecord(user)
	// db.Create(&Product{Code: "L1212", Price: 1000})
	return true
}

// Get 根据id查询
func (u *Users) Get() {
	db.First(&u, u)
}
