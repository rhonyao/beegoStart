/*
 * @Author: zhenwei zhang 
 * @Date: 2018-03-29 18:20:13 
 * @Last Modified by: zhenwei zhang
 * @Last Modified time: 2018-06-24 13:51:56
 */
 
package controllers

import (
	"net/url"
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/astaxie/beego/orm"
	"github.com/virskor/beegoStart/models"
)

type ApiResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type BaseController struct {
	beego.Controller
	o orm.Ormer
}

type Page struct {
	PageSize int
	PageNumer int
}

var (
	userInfo string
	cpt *captcha.Captcha
	env string
	captchaId string
)

func init() {
	env = beego.AppConfig.String("runmode")
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/api/v1/captcha/get", store)
	cpt.ChallengeNums = 4
	cpt.StdWidth = 100
	cpt.StdHeight = 40
	if env == "dev" {
		orm.Debug = true
	}
}

func (this *BaseController) Prepare() {
	/*beeGo construct*/
	//controllerName, actionName := this.GetControllerAndAction()
	this.o = orm.NewOrm();
}

func (this *BaseController) DisplayJson(code int, message string, data interface{}) {
	/*display json data*/
	this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	apiResult := &ApiResponse {
		code,
		message,
		data, 
	}
	this.Data["json"] = &apiResult
	this.ServeJSON()
	this.StopRun()
} 

func (this *BaseController) ApiLogin(){
	/*check if user status*/
	user:= this.GetSession("user")
	if user == nil {
		this.DisplayJson(401, "您还没有登录或登录过期，请您登录后再进行操作" , nil);  
	}
}

func (this *BaseController) ShouldLogin() {
	user:= this.GetSession("user")
	if user == nil {
		Redirect := this.Ctx.Request.RequestURI
		this.Ctx.Redirect(302,"/api/oauth/#/login?redirect=" + url.QueryEscape(Redirect))
		this.StopRun()
		return
	}
}

func (this *BaseController) GetUser() (*models.Users, error){
	/*get the information of user*/
	var errors error
	userSession := this.GetSession("user")
	if userSession == nil {
		return nil, errors
	}
	result := userSession
	user := result.(*models.Users)
	/*update user balance for session*/
	money,err := user.GetMoney(user.Id)
	user.Money = money
	return user, err
}

func (this *BaseController) GetClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}