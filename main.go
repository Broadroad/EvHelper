package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//	_ "sfapi/docs"
	"sfapi/models"
	_ "sfapi/routers"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:zk@/sf?charset=utf8&loc=Local")
	orm.RegisterModel(&models.SfOrder{})
}

func main() {
	orm.Debug = true
	_, err := orm.GetDB()
	if err != nil {
		fmt.Println("get default DataBase")
	}
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
