/*
 * @Author: zhenwei zhang 
 * @Date: 2018-03-29 18:20:23 
 * @Last Modified by: zhenwei zhang
 * @Last Modified time: 2018-06-24 00:22:56
 */
 
package controllers

import (
    "github.com/astaxie/beego"
)

type ErrorController struct {
    beego.Controller
}

func (this *ErrorController) Error404() {
    this.Data["content"] = "page not found"
    this.TplName = "error/404.html"
}

func (this *ErrorController) Error501() {
    this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
    this.Data["content"] = "server error"
    this.TplName = "error/501.html"
}

func (this *ErrorController) ErrorDb() {
    this.Data["content"] = "database is now down"
    this.TplName = "error/dberror.html"
}

func (this *ErrorController) IE() {
    this.Data["content"] = "for your security, internet explore is not supported yet"
    this.TplName = "error/dberror.html"
}