package usercontroller

import (
	"blog/app/controllers"
	"blog/models"
	"blog/utils/jwt"
	"net/http"

	"github.com/gorilla/schema"
	"gopkg.in/go-playground/validator.v9"
)

//UserController 用户控制器
type UserController struct {
	user models.Users
}

//Login 登陆
func (c UserController) Login(w http.ResponseWriter, r *http.Request) {
	rs := controllers.Result{Code: 20000}
	defer controllers.ReturnJson(w, &rs)
	//1.数据接收
	r.ParseForm()
	var decoder = schema.NewDecoder()
	in := LoginIn{}
	err := decoder.Decode(&in, r.Form)
	if err != nil {
		rs.Code = 400
		rs.Data = err.Error()
		return
	}
	//2.数据验证
	validate := validator.New()
	err = validate.Struct(in)
	if err != nil {
		rs.Code = 400
		rs.Data = err.Error()
		return
	}
	user := c.user
	user.Username = in.Username
	user.Password = in.Password
	user.Get()
	if user.ID < 1 {
		rs.Code = 400
		rs.Data = "username or password is wrong"
		return
	}
	token, err := jwt.GetToken(in.Username, in.Password)
	if err != nil {
		rs.Code = 400
		rs.Data = err.Error()
		return
	}
	//3.返回值
	rs.Data = LoginOut{
		Token: token,
	}
}

//TestAdd 添加
func (c UserController) TestAdd(w http.ResponseWriter, r *http.Request) {
	rs := controllers.Result{Code: 60204}
	defer controllers.ReturnJson(w, &rs)
	usermodel := c.user
	if usermodel.AddTest() == false {
		rs.Code = 400
	}
}

// GetByToken 获取记录
func (u UserController) GetByToken(w http.ResponseWriter, r *http.Request) {
	rs := controllers.Result{Code: 20000, Data: GetByTokenOut{Roles: "admin", Introduction: "a manager", Avatar: "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif", Name: "liuqi"}}
	defer controllers.ReturnJson(w, &rs)
	return
}

// Logout 注销
func (u UserController) Logout(w http.ResponseWriter, r *http.Request) {
	rs := controllers.Result{Code: 20000, Data: "success"}
	defer controllers.ReturnJson(w, &rs)
	return
}
