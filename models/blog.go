package models

import (
	"context"

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
func (b *Blog) GetByPage(ctx context.Context) (blogs []Blog, err error) {
	db := ctx.Value("where").(*gorm.DB)
	if errs := db.Find(&blogs).GetErrors(); len(errs) > 0 {
		err = errs[0]
	}
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
