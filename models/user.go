package models

import (
	"context"
	"errors"

	"github.com/jinzhu/gorm"
)

// Users 用户基础模型
type User struct {
	gorm.Model
	Username string `json:"username" validate:"required,min=3,max=15"`
	Password string `json:"password" validate:"required" `
	UserType string `json:"usertype" validate:"required"`
}

// Create 新建
func (u *User) Create(ctx context.Context) (errs []error) {
	user := User{}
	if errs = db.Where("username=?", u.Username).First(&user).GetErrors(); len(errs) > 0 {
		return
	}
	if user.ID > 0 {
		errs = make([]error, 1)
		errs[0] = errors.New("用户名早已存在!")
		return
	}
	if errs = db.Create(&u).GetErrors(); len(errs) > 0 {
		return
	}
	return
}

// Update 更新
func (u *User) Save(ctx context.Context) (errs []error) {
	return db.Save(&u).GetErrors()
}

// Get 根据id查询
func (u *User) Get() {
	db.First(&u, u)
}
