/*
 * @Author: zhenwei zhang 
 * @Date: 2018-03-29 18:22:42 
 * @Last Modified by: zhenwei zhang
 * @Last Modified time: 2018-06-24 20:19:53
 */
package routers

import (
	"github.com/astaxie/beego"
	"github.com/virskor/beegoStart/controllers"
)

func init() {
	beego.Router("/api/v1/oauth", &controllers.UserController{}, "get:Get")
	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/oss",
			beego.NSRouter("/upload", &controllers.OSSController{}, "post:Upload"),
		),
		beego.NSNamespace("/user",
			beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
			beego.NSRouter("/create", &controllers.UserController{}, "post:Create"),
			beego.NSRouter("/logout", &controllers.UserController{}, "get,post:Logout"),
			beego.NSRouter("/getinfo", &controllers.UserController{}, "get,post:GetUserInfo"),
			beego.NSRouter("/captcha", &controllers.UserController{}, "get:Captcha"),
			beego.NSRouter("/check", &controllers.UserController{}, "post:UserNameExits"),
			beego.NSRouter("/reset", &controllers.UserController{}, "post:Reset"),
			beego.NSRouter("/avatar/:id", &controllers.UserController{}, "get:Avatar"),
			beego.NSRouter("/avatar/update", &controllers.UserController{}, "post:UpdateAvatar"),
			beego.NSRouter("/bind", &controllers.UserController{}, "get,post:Bind"),
		),
		beego.NSNamespace("/security",
			beego.NSRouter("/getMessageCode", &controllers.MessageCodeController{}, "post:Send"),
		),
		beego.NSNamespace("/auth",
			beego.NSRouter("/github/login", &controllers.GithubController{}, "get:Login"),
		),
		beego.NSNamespace("/payment",
			beego.NSRouter("/create", &controllers.PaymentController{}, "post:Create"),
			beego.NSRouter("/list", &controllers.PaymentController{}, "post:List"),
			beego.NSRouter("/alipay", &controllers.AlipayController{}, "get:Native"),
			beego.NSRouter("/alipay/return", &controllers.AlipayController{}, "get:Return"),
			beego.NSRouter("/alipay/notify", &controllers.AlipayController{}, "post,get:Notify"),
		),
	)
	beego.AddNamespace(ns)
}