/*
 * @Author: zhenwei zhang 
 * @Date: 2018-03-29 18:22:13 
 * @Last Modified by: zhenwei zhang
 * @Last Modified time: 2018-06-24 00:16:21
 */

package models

import (
	"errors"
	"strconv"
	"time"
	"math/rand"
	"github.com/astaxie/beego/orm"
)

type Payment struct{
	Id int `json:"id"` /*支付单ID*/
	Pid string `json:"pid"` /*Pid是唯一的订单支付流水账号，没有使用id存储*/
	Uid string `json:"uid"` /*创建人的UID*/
	CreateDate time.Time `json:"createDate"` /*创建时间*/
	FinishDate time.Time `orm:"size(255);null" json:"finishDate"` /*支付完成时间*/
	Method string `json:"method"` /*支付渠道 aliapy wexin paypal*/
	Fee float64 `json:"fee"` /*需支付费用*/
	CurrentFee float64 `orm:"size(255);null" json:"currentfee"` /*实际支付的费用*/
	TotalFee float64 `json:"totalFee"` /*支付渠道收取的手续费用*/
	Status string `json:"status"` /*是否已经支付 create finished quit*/
	ContactEmail string `orm:"size(255);null" json:"ContactEmail"` /*联系邮箱，用户支付后支付渠道回调得来的信息*/
	CallBackData string `orm:"size(255);null" json:"callBackData"` /*支付渠道回调信息*/
	Action string `json:"action"` /*increase 充值， custom 消费， other 其他*/
}

func (m *Payment) TableName() string {
	return "payment"
}

func (m *Payment) Create(uid string, fee float64, totalFee float64, action string, method string) (string, error){
	/*create a new order*/
	o := orm.NewOrm()
	uuid := m.Rand_generator(1000000000)
	timeUnix := time.Now().Unix()
	timeUnixInt32 := strconv.FormatInt(timeUnix, 10)
	pid := uid + strconv.Itoa( <-uuid ) + timeUnixInt32

	status := "create"
	if action != "increase" {
		status = "finished"
	}
	m = &Payment{
		Uid: uid,
		Pid: pid,
		CreateDate: time.Now(),
		Method: method,
		Fee: fee,
		TotalFee: totalFee,
		Action: action,
		Status: status,
	}
	if _, err := o.Insert(m); err != nil {
		return "" , err
	} else {
		return pid , nil
	}
}

func (m *Payment) UpadateToPaid(pid string,currentFee float64, contactEmail string) (string, error){
	if pid == "" {
		return "", errors.New("订单编号不能为空")
	}
	o := orm.NewOrm()
	update := Payment{Pid:pid}
	if o.Read(&update, "pid") == nil {
		if update.Status != "create" {
			return update.Uid, errors.New("已经完成过过充值")//errors.New("订单已经完成充值，无法继续加值")
		}else{
			update.FinishDate = time.Now()
			update.CurrentFee = currentFee
			update.ContactEmail = contactEmail
			update.Status = "finished"
			if _, err := o.Update(&update, "FinishDate", "CurrentFee", "ContactEmail", "Status"); err == nil {
				return update.Uid, nil
			}else{
				return "", err
			}
		}
	}else{
		return "", errors.New("找不到订单号")
	}
}

func (m *Payment) List(uid int,pageSize int, pageNumber int) ([]*Payment, int64, error) {
	var payment []*Payment
	o := orm.NewOrm()
	startFrom := (pageNumber -1) * pageSize
	_,err := o.QueryTable(m.TableName()).OrderBy("-Id").Filter("Uid", uid).Limit(pageSize,startFrom).All(&payment)
	count,err :=  o.QueryTable(m.TableName()).Filter("Uid", uid).Count()
	return payment, count, err
 }

func (m *Payment) Rand_generator(n int) chan int {
    rand.Seed(time.Now().UnixNano())
    out := make(chan int)
    go func(x int) {
        for {
            out <- rand.Intn(x)
        }
    }(n)
    return out
}