package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"sfapi/models"
	"time"
)

// Operations about Order
type SfOrderController struct {
	beego.Controller
}

// @Title Get
// @Description get all Orders
// @Success 200 {object} models.Order
// @router / [get]
func (u *SfOrderController) GetAll() {
	orders := &models.SfOrder{}
	users, _ := orders.GetAllOrders()
	u.Data["json"] = users
	u.ServeJSON()
}

//@router / [post]
func (this *SfOrderController) Add() {
	jsonResult := &models.JsonResult{}
	var order models.SfOrder
	json.Unmarshal(this.Ctx.Input.RequestBody, &order)
	userid := order.Userid
	senderid := order.Senderid
	expresslocation := order.Expresslocation
	senderlocation := order.Senderlocation

	if userid == 0 {
		jsonResult.Message = "接单者id不能为空"
	} else if senderid == 0 {
		jsonResult.Message = "发单者id不能为空"
	} else if expresslocation == "" {
		jsonResult.Message = "快递员地址不能为空"
	} else if senderlocation == "" {
		jsonResult.Message = "发单者地址不能为空"
	}

	if jsonResult.Message == "" {
		order.Sendertime = time.Now()
		order.IsDeleted = 0
		err := order.SaveOrder()

		if err != nil {
			jsonResult.Message = "保存失败，请重试"
		} else {
			jsonResult.Message = "保存成功"
			jsonResult.Data = order
			jsonResult.Success = true
		}
	}

	this.Data["json"] = jsonResult
	this.ServeJSON()
}
