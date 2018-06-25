/*
 * @Author: zhenwei zhang 
 * @Date: 2018-03-29 18:20:37 
 * @Last Modified by: zhenwei zhang
 * @Last Modified time: 2018-06-24 13:52:16
 */
 
 package controllers

 import (
	"strings"
	"strconv"
	"encoding/json"
	"github.com/virskor/beegoStart/models"
 )

 type PaymentController struct {
	BaseController
 }

 type CreatePaymentParams struct{
	Fee string
	Method string
 }

 type ListPayment struct{
	PageSize int `json:"pageSize"`
	PageNumber int `json:"pageNumber"`
	PageCount int64 `json:"pageCount"`
	Table interface{} `json:"table"`
 }
 
 func (this *PaymentController) Create() {
	this.ApiLogin()
	var pay CreatePaymentParams
	post := this.Ctx.Input.RequestBody
	json.Unmarshal(post,&pay)
	if strings.Trim(pay.Fee, " ") == "" || strings.Trim(pay.Method, " ") == "" {
		this.DisplayJson(0, "所需参数不能为空", nil)
	}
	var payment models.Payment
	user,_ := this.GetUser()
	convFee,convErr := strconv.ParseFloat(pay.Fee, 32)
	zeroFee,_ := strconv.ParseFloat("0.00", 32)
	if convErr != nil{
		this.DisplayJson(0, "支付价格出现错误，请填写数字不是字符串", nil)
	}
	pid,err := payment.Create(
		strconv.Itoa(user.Id),
		convFee,
		zeroFee,
		"increase",
		pay.Method,
	)
	if err == nil{
		this.DisplayJson(1, "获取支付订单成功", map[string]string{"pid":pid })
	}else{
		this.DisplayJson(0, "生成订单号出现错误", nil)
	}
 }

 func (this *PaymentController) List() { 
	this.ApiLogin()
	var listPayment ListPayment
	var con models.Payment
	post := this.Ctx.Input.RequestBody
	json.Unmarshal(post,&listPayment)
	if listPayment.PageNumber == 0{
		listPayment.PageNumber = 1
	}
	if listPayment.PageSize == 15{
		listPayment.PageSize = 15
	}
	user,_ := this.GetUser()
	result,count,err := con.List(user.Id, listPayment.PageSize, listPayment.PageNumber)
	if err == nil{
		listPayment.Table = &result
		listPayment.PageCount = (count / int64(listPayment.PageSize)) + 1
		this.DisplayJson(1,"成功",listPayment)
	}else{
		this.DisplayJson(0,"读取数据出错，请重试",nil)
	}
 }