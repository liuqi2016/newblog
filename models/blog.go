package models

import (
	"github.com/jinzhu/gorm"
)

// Blog 用户基础模型
type Blog struct {
	gorm.Model
	Title     string `json:"title" validate:"required"`
	Content   string `json:"content" validate:"required"`
	Author    string `json:"author" validate:"required"`
	Pageviews uint   `json:"pageviews"`
}

// Get 根据id查询
func (b *Blog) Get() {
	db.First(&b, b)
}

// GetByPage 根据分页查询
func (b *Blog) GetByPage() (blogs []Blog, err []error) {
	db.Where(&b).Find(&blogs).GetErrors()
	return
}

// Create 新建
func (b *Blog) Create() (errs []error) {
	return db.Create(&b).GetErrors()
}

// Update 更新
func (b *Blog) Save() (errs []error) {
	return db.Save(&b).GetErrors()
}
