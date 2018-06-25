/*
 * @Author: zhenwei zhang 
 * @Date: 2018-03-29 18:22:29 
 * @Last Modified by: zhenwei zhang
 * @Last Modified time: 2018-04-17 16:41:23
 */

 package models
 
 import (
	"time"
 )

 type API struct{
	 Id int `json:"id"`
	 Uid int `json:"uid"`
	 AppId string `json:"appId"`
	 AppSecret string `json:"appSecret"`
	 GenerateDate time.Time `json:"generateDate"`
	 Status int `json:"status"`
 }