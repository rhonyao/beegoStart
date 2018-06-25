/*
 * @Author: zhenwei zhang 
 * @Date: 2018-03-29 18:22:29 
 * @Last Modified by: zhenwei zhang
 * @Last Modified time: 2018-04-20 22:08:23
 */

package models

import (
	"time"
	"errors"
	"github.com/astaxie/beego/orm"
)

type Users struct{
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Salt string	`json:"-"`
	Truename string `json:"truename"`
	Idnumber string `json:"-"`
	Verify int `json:"verify"`
	Usergroup int `json:"userGroup"`
	Registrydate time.Time `json:"registryDate"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Credits int `json:"credits"`
	Money float64 `json:"money"`
	Status int `json:"status"`
}

func (m *Users) TableName() string {
	return "users"
}

func (m *Users) UpdateMoney(uid int, increase float64) error {
	if uid == 0 {
		return errors.New("uid不能为空")
	}
	o := orm.NewOrm()
	user := Users{Id: uid}
	if o.Read(&user, "Id") == nil {
		oldMoneyBalance := user.Money
		user.Money = oldMoneyBalance + increase
		if _, err := o.Update(&user, "Money"); err == nil {
			return nil
		}else{
			return err
		}
	}else{
		return errors.New("找不到用户")
	}
}

func (m *Users) CutMoney(uid int, cutFee float64) error{
	/*更新用户的余额，method = 0 cut = 1 add*/
	if uid == 0 {
		return errors.New("uid不能为空")
	}
	o := orm.NewOrm()
	user := Users{Id: uid}
	if o.Read(&user, "Id") == nil {
		oldMoneyBalance := user.Money
		user.Money = oldMoneyBalance - cutFee
		if _, err := o.Update(&user, "Money"); err == nil {
			return nil
		}else{
			return err
		}
	}else{
		return errors.New("找不到用户")
	}
}

func (m *Users) GetMoney(uid int) (float64, error) {
	if uid == 0 {
		return 0, errors.New("uid不能为空")
	}
	o := orm.NewOrm()
	user := Users{Id: uid}
	if o.Read(&user, "Id") == nil {
		return user.Money, nil
	}else{
		return 0, errors.New("找不到用户")
	}
}
