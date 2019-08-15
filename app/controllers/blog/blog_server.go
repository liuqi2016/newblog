package blog

import (
	"context"
)

// BlogServer 博客接口
type BlogServer interface {
	GetOne(ctx context.Context) (r interface{}, err error)
	GetByPage(ctx context.Context) (r interface{}, err error)
	SaveOrUpdate(ctx context.Context) (r interface{}, err error)
}

// // SaveOrUpdateIn
// type SaveOrUpdateIn struct {
// 	ID          uint   `json:"id"`
// 	Title       string `json:"title" validate:"required"`
// 	Content     string `json:"content" validate:"required"`
// 	Author      string `json:"author" validate:"required"`
// 	DisplayTime string `json:"display_time" validate:"required"`
// 	Pageviews   uint   `json:"pageviews" validate:"required"`
// }

type GetByPageIn struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
	// DeletedAt *time.Time `sql:"index"`
}
