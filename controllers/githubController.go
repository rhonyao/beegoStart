/*
 * @Author: zhenwei zhang 
 * @Date: 2018-03-29 18:20:28 
 * @Last Modified by: zhenwei zhang
 * @Last Modified time: 2018-06-24 13:17:09
 */
 
package controllers

import (
	"strings"
	"errors"
	"github.com/astaxie/beego"
	"github.com/virskor/beegoStart/oauth"
)

type GithubController struct {
	BaseController
}

func (this *GithubController) Login() {
	var (
		ierr error
	)
	githubClientId := beego.AppConfig.String("oauth::githubClientId_" + env )
	githubClientCallback := beego.AppConfig.String("oauth::githubClientCallback_" + env)
	code := this.GetString("code")
	if strings.Trim(code, "") == ""{
		redirectUrl := "https://github.com/login/oauth/authorize?client_id=" + githubClientId + "&redirect_uri=" + githubClientCallback
		this.Redirect(redirectUrl, 302)
		return
	}
	if token, err := oauth.GetGithubAccessToken(code, env); err != nil {
		ierr = err
	} else {
		if info, err := oauth.GetGithubUserInfo(token.AccessToken); err != nil {
			ierr = err
		} else {
			if info.Id > 0 {
				this.SetSession("githubUser",&info)
			} else {
				ierr = errors.New("获取github用户数据失败")
			}
		}
	}
	if ierr != nil {
		this.Ctx.WriteString(ierr.Error())
	}else{
		this.Redirect("../../user/bind/", 302)
	}
}