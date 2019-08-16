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
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	UpdatedAt string `json:"updated_at"`
	Row       int    `json:"row" validate:"required"`  //	每页行数
	Page      int    `json:"page" validate:"required"` //	当前页码
	// DeletedAt *time.Time `sql:"index"`
}
type GetByPageOut struct {
	Total      int         `json:"total"`       //总记录数
	Row        int         `json:"row"`         //	每页行数
	Page       int         `json:"page"`        //	当前页码
	TotalPages int         `json:"total_pages"` //总页数
	Lists      interface{} `json:"lists"`
}
