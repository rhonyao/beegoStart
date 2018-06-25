/*
 * @Author: zhenwei zhang 
 * @Date: 2018-03-29 18:22:03 
 * @Last Modified by: zhenwei zhang
 * @Last Modified time: 2018-04-24 18:29:33
 */
 
 package models

 import (
	"time"
 )

 type Verify struct{
	Id int `json:"id"`
	Uid int `json:"uid"`
	TrueName string `json:"trueName"`
	IdNo string `json:"idNo"`
	Images string `json:"images"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	Status string `json:"status"` /*0审核中，1审核完成，2审核失败*/
 }