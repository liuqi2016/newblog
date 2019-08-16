package blog

import (
	"blog/models"
	"context"
	"encoding/json"
	"math"
	"net/http"

	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
)

type BlogController struct {
	blog models.Blog
}

//GetOne 获取单个详情
func (b BlogController) GetOne(ctx context.Context) (rs interface{}, err error) {
	r := ctx.Value("request").(*http.Request)
	r.ParseForm()
	rs = r.Form.Get("id")
	return
}

//GetByPage 获取列表
func (b BlogController) GetByPage(ctx context.Context) (rs interface{}, err error) {
	r := ctx.Value("request").(*http.Request)
	in := GetByPageIn{}
	blog := b.blog
	if err = json.NewDecoder(r.Body).Decode(&in); err != nil {
		return
	}
	if err = validator.New().Struct(in); err != nil {
		return
	}
	db := ctx.Value("db").(*gorm.DB)
	if in.ID > 0 {
		db = db.Where("id = ?", in.ID)
	}
	if in.Title != "" {
		db = db.Where("title like ?", "%"+in.Title+"%")
	}
	if in.Author != "" {
		db = db.Where("author = ?", in.Author)
	}
	if in.StartTime != "" {
		db = db.Where("created_at >= ?", in.StartTime)
	}
	if in.EndTime != "" {
		db = db.Where("created_at <= ?", in.EndTime)
	}
	ctx2 := ctx
	ctx2 = context.WithValue(ctx2, "where", db)
	lists, err := blog.GetByPage(ctx2)
	out := GetByPageOut{}
	out.Lists = lists
	out.Total = len(lists)
	out.Page = in.Page
	out.Row = in.Row
	out.TotalPages = int(math.Ceil(float64(out.Total) / float64(out.Row)))
	return out, err
}

//SaveOrUpdate 编辑
func (b BlogController) SaveOrUpdate(ctx context.Context) (rs interface{}, err error) {
	r := ctx.Value("request").(*http.Request)
	decoder := json.NewDecoder(r.Body)
	// var decoder = schema.NewDecoder()
	in := b.blog
	err = decoder.Decode(&in)
	if err != nil {
		return
	}

	validate := validator.New()
	err = validate.Struct(in)
	if err != nil {
		return
	}
	if in.ID > 0 {
		//编辑
		if errs := in.Save(); len(errs) > 0 {
			return nil, errs[0]
		}
	} else {
		//新增
		if errs := in.Create(); len(errs) > 0 {
			return nil, errs[0]
		}
	}
	return
}
