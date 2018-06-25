package controllers

import (
	"strings"
	"alipay"
	"strconv"
	"github.com/astaxie/beego"
	"github.com/virskor/beegoStart/models"
)

type AlipayController struct {
	BaseController
}

func newClient() *alipay.Client {
	return &alipay.Client{
		Partner:   beego.AppConfig.String("payment::alipartner"),
		Key:       beego.AppConfig.String("payment::alikey"),
		ReturnUrl: beego.AppConfig.String("payment::domainurl") + "/alipay/return",
		NotifyUrl: beego.AppConfig.String("payment::domainurl") + "/alipay/notify",
		Email:     beego.AppConfig.String("payment::aliemail"),
	}
}

func (this *AlipayController) Native() {
	this.ShouldLogin()
	user,err := this.GetUser()
	orderNo := this.GetString("orderNo")
	orderFee := this.GetString("orderFee")
	if strings.Trim(orderNo, " ") == "" || strings.Trim(orderFee, " ") == "" {
		this.Ctx.WriteString("orderNo与orderFee不能为空")
		this.StopRun()
	}
	if err == nil{
		schemestr := this.Ctx.Input.Site()	
		alipayClient := newClient()
		fee, _ := strconv.ParseFloat(orderFee, 32)
		ots := alipay.Options{
				OrderId:            orderNo,
				Fee:                float32(fee),
				NickName:           user.Username,
				Subject:            "为账户加值",
				Extra_common_param: schemestr,
		}
		form := alipayClient.Form(ots)
		this.Data["form"] = form
		this.TplName = "payment/alipay.html"
	}else{
		this.Ctx.WriteString("用户信息有误，请您核实您的登录状态是否正常")
	}
}

func (this *AlipayController) Return() {
	alipayClient := newClient()
	result := alipayClient.Return(&this.Controller)
	//beego.Debug("notify", result)
	if result.Status == 1 { /* payment success*/
		if result.Extra_common_param != "" {
			this.TplName = "payment/success.html"
		}
	} else {
		this.TplName = "payment/fail.html"
	}
}

func (this *AlipayController) Notify() {
	alipayClient := newClient()
	result := alipayClient.Notify(&this.Controller)
	if result.Status == 1 {
		if result.Extra_common_param != "" {
			//update payment order
			var payment models.Payment
			fee, _ := strconv.ParseFloat(result.TotalFee, 64)
			uid,err := payment.UpadateToPaid(
				result.OrderNo,
				float64(fee),
				result.BuyerEmail,
			)
			if err == nil {
				/*account balance increase*/
				var updateUser models.Users
				uid, _ := strconv.Atoi(uid)
				err := updateUser.UpdateMoney(uid, float64(fee))
				if err == nil {
					this.Ctx.WriteString("success")
				}else{
					this.Ctx.WriteString(err.Error())
				}
			}else{
				this.Ctx.WriteString(err.Error())
			}
		}else{
			this.Ctx.WriteString("来源验证失败")
		}
	}
}