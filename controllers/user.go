package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"sfapi/models"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

func (u *UserController) Get() {
	jsonResult := &models.JsonResult{}
	telephone := u.GetString(":telephone")
	fmt.Printf(telephone)
	user := &models.User{Telephone: telephone}
	if telephone != "" {
		user, err := user.GetUser()
		if err != nil {
			jsonResult.Status = 1
			jsonResult.Message = err.Error()
			jsonResult.Success = false
		} else {
			jsonResult.Status = 0
			jsonResult.Message = "OK"
			jsonResult.Data = user
			jsonResult.Success = true
		}
	}
	u.Data["json"] = jsonResult
	u.ServeJSON()
}

func (u *UserController) Post() {
	jsonResult := &models.JsonResult{}
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	err := user.AddUser()
	if err != nil {
		jsonResult.Message = "保存失败"
	} else {
		jsonResult.Message = "保存成功"
		jsonResult.Data = user
		jsonResult.Success = true
	}

	u.Data["json"] = jsonResult
	u.ServeJSON()
}
