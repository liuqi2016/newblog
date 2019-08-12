package usercontroller

import (
	"blog/app/controllers"
	"blog/models"
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
	rs := controllers.Result{Code: 200, Message: "success"}
	defer controllers.ReturnJson(w, &rs)
	//1.数据接收
	r.ParseForm()
	var decoder = schema.NewDecoder()
	in := LoginIn{}
	err := decoder.Decode(&in, r.Form)
	if err != nil {
		rs.Code = 400
		rs.Message = "fail"
		rs.Data = err.Error()
		return
	}
	//2.数据验证
	validate := validator.New()
	err = validate.Struct(in)
	if err != nil {
		rs.Code = 400
		rs.Message = "fail"
		rs.Data = err.Error()
		return
	}
	//3.返回值
	rs.Data = LoginOut{
		Token: "qwewqeweqwewewqeqweqweqweccscscsd",
	}
}

//Add 添加
func (c UserController) TestAdd(w http.ResponseWriter, r *http.Request) {
	rs := controllers.Result{Code: 200, Message: "success"}
	defer controllers.ReturnJson(w, &rs)
	usermodel := c.user
	if usermodel.AddTest() == false {
		rs.Code = 400
		rs.Message = "fail"
	}
}
