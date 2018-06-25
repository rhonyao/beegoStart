/*
 * @Author: zhenwei zhang 
 * @Date: 2018-03-29 18:20:40 
 * @Last Modified by: zhenwei zhang
 * @Last Modified time: 2018-06-23 23:51:03
 */
 
package controllers

import (
	"strings"
    "net/http"
    "net/url"
    "encoding/json"
    "fmt"
    "bytes"
	"io/ioutil"
	"math/rand"
	"unsafe"
	"time"
	"github.com/astaxie/beego"
)

type MessageCodeController struct {
	BaseController
}

type MessageParams struct {
	Phone   string `json:"phone"`
	Captcha string `json:"captcha"`
}

func (this *MessageCodeController) Send() {
	var postData MessageParams
	body := this.Ctx.Input.RequestBody
	json.Unmarshal(body, &postData)
	if strings.Trim(postData.Phone," ") == "" || strings.Trim(postData.Captcha," ") == "" {
		this.DisplayJson(0, "缺少必要参数" , nil);
		return
	}
	if !cpt.Verify(captchaId,postData.Captcha){
		this.DisplayJson(0, "图形验证码错误，短信发送失败" , nil);
		return
	}
	randCode := rand.New(rand.NewSource(time.Now().UnixNano())) 
	messageCode := fmt.Sprintf("%06v", randCode.Int31n(1000000))
	this.SetSession("messageCode",messageCode)
	this.SetSession("messageCodePhone",postData.Phone)
	params := make(map[string]interface{})
    params["account"] = beego.AppConfig.String("message::account")
    params["password"] = beego.AppConfig.String("message::password")
	params["phone"] = string(postData.Phone)
    params["msg"] = url.QueryEscape(beego.AppConfig.String("message::sign") + this.TemplateCompiler(beego.AppConfig.String("message::template"), string(messageCode))) 
    params["report"] = "false"
    bytesData, err := json.Marshal(params)
    if err != nil { 
		fmt.Println(err.Error() )
		this.DisplayJson(0, "验证码发送失败" , nil);
        return
    }
    reader := bytes.NewReader(bytesData)
    url := "http://smssh1.253.com/msg/send/json"
    request, err := http.NewRequest("POST", url, reader)
    if err != nil {
		fmt.Println(err.Error())
		this.DisplayJson(0, "验证码发送失败" , nil);
        return
    }
    client := http.Client{}
    resp, err := client.Do(request)
    if err != nil {
		fmt.Println(err.Error())
		this.DisplayJson(0, "验证码发送失败" , nil);
        return
    }
    respBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
		fmt.Println(err.Error())
		this.DisplayJson(0, "验证码发送失败" , nil);
        return
    }
    str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)
	this.DisplayJson(1, "验证码发送成功" , nil);
}

func (this *MessageCodeController) TemplateCompiler(template string, code string) string{
	str := template
    str = strings.Replace(str, "${var}", code, -1)
	return str
}