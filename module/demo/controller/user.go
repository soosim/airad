package controller

import (
	"airad/common/base"
	"airad/common/util"
	"airad/module/demo/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Operations about Users
type UserController struct {
	base.BaseController
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (c *UserController) Post() {
	var v models.User
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if errorMessage := util.CheckNewUserPost(v.Username, v.Password,
			v.Age, v.Gender, v.Address, v.Email); errorMessage != "ok" {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = base.BaseResponse{403, 403, errorMessage, ""}
			c.ServeJSON()
			return
		}
		if models.CheckUserName(v.Username) {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = base.BaseResponse{403, 403, "用户名称已经注册了", ""}
			c.ServeJSON()
			return
		}
		if models.CheckEmail(v.Email) {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = base.BaseResponse{403, 403, "邮箱已经注册了", ""}
			c.ServeJSON()
			return
		}

		/*if user, err := models.AddUser(&v); err == nil {
		      c.Ctx.Output.SetStatus(201)
		      var returnData = &UserSuccessLoginData{user.Token, user.Username}
		      c.Data["json"] = &base.BaseResponse{0, 0, "ok", returnData}
		  } else {
		      c.Data["json"] = &base.BaseResponse{1, 1, "用户注册失败", err.Error()}
		  }*/
	} else {
		c.Data["json"] = &base.BaseResponse{1, 1, "用户注册失败", err.Error()}
	}
	c.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (c *UserController) GetAll() {
	token := strings.TrimSpace(c.Ctx.Input.Header("token"))
	_, file, line, _ := runtime.Caller(0)
	logs.Debug("[%s.%d] %s", file, line, "token param:"+token)
	logs.Debug("[%s.%d] %s", file, line, "logID:"+c.Ctx.Input.GetData("commonLogId").(string))
	et := util.NewEasyToken()
	ok, err := et.ValidateToken(token)
	if !ok {
		c.Ctx.ResponseWriter.WriteHeader(401)
		c.Data["json"] = base.BaseResponse{401, 401, fmt.Sprintf("%s", err), ""}
		c.ServeJSON()
		return
	}
	l, err := models.GetUserAll()
	if err != nil {
		c.Success(err.Error())
	} else {
		c.Success(l)
	}
	return
}

// GetOne ...
// @Title GetOne
// @Description get User by id
// @Param	id		path 	string	true "The key for static block"
// @Success 200 {object} models.AirAd
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UserController) GetOne() {
	token := c.Ctx.Input.Header("token")
	//idStr := c.Ctx.Input.Param("id")
	idStr := c.Ctx.Input.Param(":id")
	//token := c.Ctx.Input.Param(":token")
	et := util.NewEasyToken()
	//token := strings.TrimSpace(c.Ctx.Request.Header.Get("Authorization"))
	valido, err := et.ValidateToken(token)
	if !valido {
		c.Ctx.ResponseWriter.WriteHeader(401)
		c.Data["json"] = base.BaseResponse{401, 401, fmt.Sprintf("%s", err), ""}
		c.ServeJSON()
		return
	}

	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUserById(id)
	if v == nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()

}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (c *UserController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.User{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateUserById(&v); err == nil {
			c.Data["json"] = base.SuccessReturn
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (c *UserController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteUser(id); err == nil {
		c.Data["json"] = base.SuccessReturn
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [POST]
func (c *UserController) Login() {
	var reqData struct {
		Username string `valid:"Required"`
		Password string `valid:"Required"`
	}
	var token string

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &reqData); err == nil {
		if errorMessage := util.CheckUsernamePassword(reqData.Username, reqData.Password); errorMessage != "ok" {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = base.BaseResponse{403, 403, errorMessage, ""}
			c.ServeJSON()
			return
		}
		if ok, user := models.Login(reqData.Username, reqData.Password); ok {
			et := util.NewEasyToken()
			validation, err := et.ValidateToken(user.Token)
			if !validation {
				et.Username = user.Username
				et.Uid = int64(user.Id)
				et.Expires = time.Now().Unix() + 2*3600
				token, err = et.GetToken()
				if token == "" || err != nil {
					c.Data["json"] = base.ErrExpired
					c.ServeJSON()
					return
				} else {
					models.UpdateUserToken(user, token)
				}
			} else {
				token = user.Token
			}
			models.UpdateUserLastLogin(user)

			var returnData = &UserSuccessLoginData{token, user.Username}
			c.Data["json"] = &base.BaseResponse{0, 0, "ok", returnData}
		} else {
			logs.Info(ok)
			logs.Info(user)
			c.Data["json"] = &base.ErrNoUserOrPass
		}
	} else {
		c.Data["json"] = &base.ErrNoUserOrPass
	}
	c.ServeJSON()
}

// @Title 认证测试
// @Description 测试错误码
// @Success 200 {object} base.BaseResponse
// @Failure 401 unauthorized
// @router /auth [get]
func (c *UserController) Auth() {
	et := util.NewEasyToken()
	token := strings.TrimSpace(c.Ctx.Request.Header.Get("Authorization"))
	validation, err := et.ValidateToken(token)
	if !validation {
		c.Ctx.ResponseWriter.WriteHeader(401)
		c.Data["json"] = base.BaseResponse{401, 401, fmt.Sprintf("%s", err), ""}
		c.ServeJSON()
		return
	}

	c.Data["json"] = base.BaseResponse{0, 0, "is login", ""}
	c.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = base.SuccessReturn
	u.ServeJSON()
}
