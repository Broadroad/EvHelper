package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["sfapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["sfapi/controllers:ObjectController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["sfapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["sfapi/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["sfapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["sfapi/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["sfapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["sfapi/controllers:ObjectController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["sfapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["sfapi/controllers:ObjectController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["sfapi/controllers:SfOrderController"] = append(beego.GlobalControllerRouter["sfapi/controllers:SfOrderController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["sfapi/controllers:SfOrderController"] = append(beego.GlobalControllerRouter["sfapi/controllers:SfOrderController"],
		beego.ControllerComments{
			"Add",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["sfapi/controllers:UserController"] = append(beego.GlobalControllerRouter["sfapi/controllers:UserController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["sfapi/controllers:UserController"] = append(beego.GlobalControllerRouter["sfapi/controllers:UserController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["sfapi/controllers:UserController"] = append(beego.GlobalControllerRouter["sfapi/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/:uid`,
			[]string{"get"},
			nil})

}
