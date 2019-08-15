package user

import (
	"blog/models"
	"blog/utils/jwt"
	"context"
	"errors"
	"net/http"

	"github.com/gorilla/schema"
	"gopkg.in/go-playground/validator.v9"
)

//UserController 用户控制器
type UserController struct {
	user models.User
}

//Login 登陆
func (c UserController) Login(ctx context.Context) (rs interface{}, err error) {
	r := ctx.Value("request").(*http.Request)
	//1.数据接收
	r.ParseForm()
	var decoder = schema.NewDecoder()
	in := LoginIn{}
	err = decoder.Decode(&in, r.Form)
	if err != nil {
		return
	}
	//2.数据验证
	validate := validator.New()
	err = validate.Struct(in)
	if err != nil {
		return
	}
	user := c.user
	user.Username = in.Username
	user.Password = in.Password
	user.Get()
	if user.ID < 1 {
		return nil, errors.New("用户名或密码错误")
	}
	token, err := jwt.GetToken(in.Username, in.Password)
	if err != nil {
		return
	}
	//3.返回值
	rs = LoginOut{
		Token: token,
	}
	return
}

//TestAdd 添加
func (c UserController) TestAdd(ctx context.Context) (rs interface{}, err error) {
	// r := ctx.Value("request").(*http.Request)
	usermodel := c.user
	if usermodel.AddTest() == false {
		return nil, errors.New("失败")
	}
	return
}

// GetInfo 获取记录
func (u UserController) GetInfo(ctx context.Context) (rs interface{}, err error) {
	// r := ctx.Value("request").(*http.Request)

	userinfo := ctx.Value("userinfo").(jwt.UserInfo)
	// fmt.Printf("%+v", userinfo)
	rs = GetInfoOut{Roles: "admin", Introduction: "a manager", Avatar: "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif", Name: userinfo.Username}
	return
}

// Logout 注销
func (u UserController) Logout(ctx context.Context) (rs interface{}, err error) {
	return "success", nil
}

// SearchByName  根据名字搜索
func (u UserController) SearchByName(ctx context.Context) (rs interface{}, err error) {
	r := ctx.Value("request").(*http.Request)
	//1.数据接收
	r.ParseForm()
	var decoder = schema.NewDecoder()
	in := SearchByNameIn{}
	err = decoder.Decode(&in, r.Form)
	if err != nil {
		return
	}
	return
}
