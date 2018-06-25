/*
 * @Author: zhenwei zhang 
 * @Date: 2018-06-23 23:58:09 
 * @Last Modified by: zhenwei zhang
 * @Last Modified time: 2018-06-24 20:23:12
 */
package main

import (
	_ "github.com/virskor/beegoStart/routers"
	_ "database/sql"
	_ "github.com/astaxie/beego/session"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/go-sql-driver/mysql"
	"encoding/gob"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego"
	"github.com/virskor/beegoStart/models"
	"github.com/virskor/beegoStart/oauth"
	"github.com/virskor/beegoStart/controllers"
)

func init(){
	/*init database config*/
	dbUser := beego.AppConfig.String("database::dbuser")
	dbPassword := beego.AppConfig.String("database::dbpassword")
	dbName := beego.AppConfig.String("database::dbname")
	ormAddress := dbUser + ":" + dbPassword + "@/" + dbName + "?charset=utf8&parseTime=true&charset=utf8&loc=Asia%2FShanghai"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", ormAddress, 30)
	orm.RegisterModel(
		new(models.Users),
		new(models.Config),
		new(models.Storage),
		new(models.Avatar),
		new(models.Payment),
		new(models.Oauth),
	)
	orm.RunSyncdb("default", false, true)
	gob.Register(new(models.Users))
	gob.Register(new(oauth.GithubUser))
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin","Content-Type","Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	beego.ErrorController(&controllers.ErrorController{})
	beego.AddTemplateExt(".html")
	beego.Run()
}