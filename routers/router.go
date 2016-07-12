// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"sfapi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),

		beego.NSRouter("/order", &controllers.SfOrderController{}, "get:GetAll"),
		beego.NSRouter("/order", &controllers.SfOrderController{}, "post:Add"),
		beego.NSRouter("/user", &controllers.UserController{}, "post:Post"),
		beego.NSRouter("/user/:telephone", &controllers.UserController{}, "get:Get"),
		beego.NSRouter("/image", &controllers.ImagesController{}, "post:Post;get:GetAll"),
		beego.NSRouter("/image/:imgId", &controllers.ImagesController{}, "get:Get"),
	)
	beego.AddNamespace(ns)
}
