package blog

import (
	"blog/models"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
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
	r.ParseForm()
	in := GetByPageIn{}
	if err = schema.NewDecoder().Decode(&in, r.Form); err != nil {
		return
	}

	if err = validator.New().Struct(in); err != nil {
		return
	}
	blog := b.blog
	if in.ID > 0 {
		blog.ID = in.ID
	} else if in.Title != "" {
		blog.Title = in.Title
	} else if in.Author != "" {
		blog.Author = in.Author
	}
	if rs, errs := blog.GetByPage(); len(errs) > 0 {
		return rs, errs[0]
	}
	return
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
