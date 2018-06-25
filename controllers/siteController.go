/*
 * @Author: zhenwei zhang 
 * @Date: 2018-03-29 18:20:44 
 * @Last Modified by: zhenwei zhang
 * @Last Modified time: 2018-06-24 20:00:47
 */
 
package controllers

type SiteController struct {
	BaseController
}

func (this *SiteController) Get() {
	this.TplName = "api/index.html"
	user,_ := this.GetUser()
	this.Data["user"] = user
	this.Render()
}