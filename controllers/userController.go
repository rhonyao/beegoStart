/*
 * @Author: zhenwei zhang 
 * @Date: 2018-03-29 18:19:49 
 * @Last Modified by: zhenwei zhang
 * @Last Modified time: 2018-06-24 19:17:05
 */
package controllers

import (
	"time"
	"fmt"
	"strings"
	"strconv"
	"encoding/json"
	"math/rand"
	"crypto/md5"
	"path/filepath"
	"os"
	"github.com/virskor/beegoStart/models"
	"github.com/mingzhehao/goutils/filetool"
	"github.com/astaxie/beego"
	"github.com/virskor/beegoStart/oauth"
	."github.com/virskor/beegoStart/utils"
)

type UserController struct {
	OSSController
}

type UserInfoStruct struct {  
	Uid int `json:"uid"`
	UserName string `json:"username"`
	TrueName string `json:"truename"` 
	UserGroup int `json:"usergroup"`
	Email string `json:"email"` 
	Status int `json:"status"`
	Credits int `json:"credits"`  
	Money int `json:"money"`  
}  

type UserLogin struct{
	UserName string `json:"username"`
	Password string `json:"password"`
	Captcha string `json:"captcha"`
}

type UserCreate struct{
	UserName string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Phone string `json:"phone"`
	MessageCode string `json:"messageCode"`
}

type CheckUserName struct{
	UserName string `json:"username"`
}

type FindAvatar struct{
	Uid string
}

type PasswordEncrpt struct{
	Salt string
	Password string
}

type PassworCompare struct{
	Password string
	Salt string
	Generated string
}

type PasswordReset struct{
	Phone string
	Password string
	MessageCode string
}

type AvatarUpload struct{
	ImageUrl string `json:"imageUrl"`
}

const pwHashBytes = 64

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (this *UserController) Get() {
	//this.TplName = "user/"+ env + ".html"
	this.TplName = "user/compress.html"
	this.Render()
}

func (this *UserController) UpdateAvatar(){
	this.ApiLogin()
	var avatar AvatarUpload
	var insertAvatar models.Avatar
	body := this.Ctx.Input.RequestBody
	json.Unmarshal(body,&avatar)
	if avatar.ImageUrl == ""{
		this.DisplayJson(0, "缺少必要参数", nil)
		return
	}
	user, _ := this.GetUser()
	targetAvatar := models.Avatar{Uid:user.Id}
	this.o.Read(&targetAvatar,"uid")
	if targetAvatar.Id == 0 {
		/*new insert*/
		insertAvatar.Uid = user.Id
		insertAvatar.AvatarUrl = avatar.ImageUrl
		if _, err := this.o.Insert(&insertAvatar); err != nil {
			this.DisplayJson(0, "更新头像失败，请联系管理员" , nil)
		} else {
			this.DisplayJson(1, "更新头像成功", nil)
		}
	}else{
		/*update avatar*/
		oldAvatarUrl := targetAvatar.AvatarUrl
		targetAvatar.AvatarUrl = avatar.ImageUrl
		_, err := this.o.Update(&targetAvatar,"avatarUrl")
		if err == nil{
			/*delete old avatar*/
			err := this.OssDelete(oldAvatarUrl)
			if err == nil{
				this.DisplayJson(1, "更新头像成功", nil)
			}else{
				this.DisplayJson(1, "更新头像成功，但OSS出现故障。错误级别低，无影响", nil)
			}
		}else{
			this.DisplayJson(0, "更新头像失败，请联系管理员", nil)
		}
	}
}

func (this *UserController) Avatar(){
	/*get user avatar*/
	var (
		appPath string
		err error
		file []byte
	)
	this.Ctx.Output.Header("Content-Type", "image/png")
	uid := this.Ctx.Input.Param(":id")
	if appPath, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {  
        panic(err)  
	} 
	convertUid,_ := strconv.Atoi(uid)
	avatar := models.Avatar{Uid:convertUid}
	this.o.Read(&avatar, "uid")
	if avatar.AvatarUrl == "" {
		file, err = filetool.ReadFileToBytes(appPath + "/static/avatar.png")
		if err != nil{
			this.Ctx.WriteString("无法匹配用户" + string(uid) + "的头像，本地静态文件出现问题")
			this.StopRun()
		}
		this.Ctx.Output.Body(file)
	}else{
		ossUrl := beego.AppConfig.String("ossdomain_" + env)
		this.Ctx.Redirect(302,ossUrl + avatar.AvatarUrl)
	}
}

func (this *UserController) Bind() {
	oauthType := ""
	findGithubUser := this.GetSession("githubUser")
	if findGithubUser != nil && findGithubUser != "" {
		oauthType = "github"
	}

	if oauthType == "" {
		this.Ctx.WriteString("Session已经过期或系统无法定义OAUTH来源类型，如果遇到这个问题请联系开发人员。")
		this.StopRun()
	}

	if oauthType == "github"{
		githubUser := findGithubUser.(*oauth.GithubUser)
		user := models.Oauth{GithubId: strconv.Itoa(githubUser.Id)}
		this.o.Read(&user, "GithubId")
		if user.Uid == 0 {
			/*notice to user for no account binding*/
			this.Data["oauthType"] = oauthType
			this.TplName = "user/bind.html"
			this.Render()
		}else{
			login := this.LoginByUid(user.Uid)
			if !login {
				this.Ctx.WriteString("登录失败，用户未找到，请稍后再试")
			}else{
				/*clean session stroage after login from github*/
				this.SetSession("githubUser","")
				this.Redirect("/", 302)
			}
		}
		
	}
	
}

func (this *UserController) LoginByUid(uid int) bool{
	/*do not set this function to router*/
	login := models.Users{Id: uid}
	this.o.Read(&login, "id")
	fmt.Println(login)
	if login.Password == "" {
		return false
	}else{
		this.SetSession("user",&login)
		return true
	}
}

func (this *UserController) Login(){
	var user UserLogin
	body := this.Ctx.Input.RequestBody
	json.Unmarshal(body,&user)
	if env != "dev"{
		if !cpt.Verify(captchaId,user.Captcha){
			this.DisplayJson(0, "图形验证码错误" , nil)
		}
	}
	result := models.Users{Username:user.UserName}
	this.o.Read(&result, "username")
	if result.Password == "" {
		this.DisplayJson(0, "登录失败，用户不存在" , nil)
	}else{
		var compare PassworCompare
		compare.Password = user.Password
		compare.Generated = result.Password
		compare.Salt = result.Salt
		compare_result := this.ComparePassword(&compare)
		if compare_result == false{
			this.DisplayJson(0, "密码错误" , nil)
		}else{
			/*if user trying to bind Oauth account*/
			this.SetSession("user",&result)
			this.AutoBind()
			this.DisplayJson(1, "登录成功" , nil)
		}
	}
}

func (this *UserController) AutoBind(){
	oauthType := ""
	findGithubUser := this.GetSession("githubUser")
	if findGithubUser != nil && findGithubUser != "" {
		oauthType = "github"
	}
	
	if oauthType != "" {
		/*bind account if oauth data exits*/
		if oauthType == "github" {
			var insertUser models.Oauth
			user,err := this.GetUser()
			if err == nil{
				githubUser := findGithubUser.(*oauth.GithubUser)
				insertUser.Uid = user.Id
				insertUser.GithubId = strconv.Itoa(githubUser.Id)
				if _, err := this.o.Insert(&insertUser); err == nil {
					this.SetSession("githubUser","")
				}
			}
		}
	}
}

func (this *UserController) Logout(){
	this.DestroySession()
	this.DisplayJson(1, "退出成功" , nil)
}

func (this *UserController) Create(){
	var createData UserCreate
	body := this.Ctx.Input.RequestBody
	json.Unmarshal(body, &createData)
	if 
		strings.Trim(createData.Phone," ") == "" ||
		strings.Trim(createData.Email," ") == "" || 
		strings.Trim(createData.Password," ") == "" || 
		strings.Trim(createData.MessageCode," ") == "" || 
		strings.Trim(createData.UserName," ") == "" {
		this.DisplayJson(0, "缺少必要参数" , nil)
		return
	}
	if this.GetSession("messageCode") != createData.MessageCode {
		this.DisplayJson(0, "手机短信验证码输入错误" , nil)
	}
	if createData.Phone != this.GetSession("messageCodePhone") {
		this.DisplayJson(0, "获取验证码的手机号与当前提交手机号不一致" , nil)
	}
	salt:= this.GenerateSalt(5)
	password := this.GeneratePassword(createData.Password, salt)
	var user CheckUserName
	json.Unmarshal(body,&user)
	result := models.Users{Username:user.UserName}
	this.o.Read(&result, "username")
	if result.Password == "" {
		if IsMailString(createData.Email) == true { /*email validation*/
			if !this.PhoneNumberExits(createData.Phone){ /*phone number validation*/
				/*create account*/
				var encrypt PasswordEncrpt
				encrypt.Salt = salt
				encrypt.Password = password
				insert := this.InsertUser(&createData,&encrypt)
				if insert == false{
					this.DisplayJson(0, "账号创建失败，请重试" , nil)
				}else{
					this.DisplayJson(1, "您的账号创建成功", nil)
				}
			}else{
				this.DisplayJson(0, "您的手机号已经被注册，请更换手机号" , nil)  
			}
		}else{
			this.DisplayJson(0, "请填写正确的邮箱账号" , nil)  
		} 
	}else{
		this.DisplayJson(0, "用户名存在" , nil)  
	}
}

func (this *UserController) Reset(){
	var reset PasswordReset
	body := this.Ctx.Input.RequestBody
	json.Unmarshal(body,&reset)
	if this.GetSession("messageCode") != reset.MessageCode {
		this.DisplayJson(0, "手机短信验证码输入错误" , nil)
	}
	findUser := models.Users{Phone:reset.Phone}
	this.o.Read(&findUser, "phone")
	if findUser.Password == "" {
		this.DisplayJson(0, "未找到对应手机号" , nil)
	}else{
		/*update pass word*/
		salt := this.GenerateSalt(5);
		password := this.GeneratePassword(reset.Password, salt)
		findUser.Salt = salt
		findUser.Password = password
		_, err := this.o.Update(&findUser,"password","salt")
		if err == nil {
			this.DestroySession()
			this.DisplayJson(1, "密码更新成功" , nil)
		}else{
			this.DisplayJson(0, "密码更新失败" , err)
		}
	}
}

func (this *UserController) GetUserInfo(){
	user := this.GetSession("user")
	if user == nil{
		this.DisplayJson(0, "您还没有登录" , nil)
	}else{
		userLatestInfo,err := this.GetUser()
		if err == nil{
			this.DisplayJson(1, "成功" , &userLatestInfo)
		}else{
			this.DisplayJson(0, "用户已经登录，但用户信息获取失败" , nil)
		}
	}
}

func (this *UserController) InsertUser(data *UserCreate, encrpt *PasswordEncrpt) bool{
	user := models.Users{}
	user.Username = data.UserName
	user.Email = data.Email
	user.Phone = data.Phone
	user.Salt = encrpt.Salt
	user.Password = encrpt.Password
	user.Status = 0
	user.Usergroup = 1
	user.Verify = 1
	user.Registrydate = time.Now()
	if _, err := this.o.Insert(&user); err != nil {
		return false
	} else {
		return true
	}
}

func (this *UserController) UserNameExits(){
	var user CheckUserName
	body := this.Ctx.Input.RequestBody
	json.Unmarshal(body,&user)
	result := models.Users{Username:user.UserName}
	this.o.Read(&result, "username")
	if result.Password == "" {
		this.DisplayJson(1, "用户名可用" , nil)  
	}else{
		this.DisplayJson(0, "用户名存在" , nil)  
	}
}

func (this *UserController) PhoneNumberExits(number string) bool{
	result := models.Users{Phone:number}
	this.o.Read(&result, "phone")
	if result.Password == "" {
		return false; 
	}else{
		return true;
	}
}

func (this *UserController) Captcha(){
	cptId, err := cpt.CreateCaptcha()
	captchaId = cptId
	if err == nil{
		captchaUrl := cpt.URLPrefix + cptId + ".png"
		this.Redirect(captchaUrl,301)
	}
}

func (this *UserController) GenerateSalt(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func (this *UserController) GeneratePassword(password string, salt string)(generated string){
	password_encryt := md5.Sum([]byte(password))
	return fmt.Sprintf("%x",md5.Sum([]byte(fmt.Sprintf("%x", password_encryt) + salt)))
}

func (this *UserController) ComparePassword(comp *PassworCompare) bool{
	var result bool
	if comp.Generated == "" || comp.Password == "" || comp.Salt == ""{
		result = false
	}
	password_encryt := md5.Sum([]byte(comp.Password))
	password_generated := fmt.Sprintf("%x",md5.Sum([]byte(fmt.Sprintf("%x", password_encryt) + comp.Salt)))
	if password_generated == comp.Generated {
		result = true
	}
	return result
}