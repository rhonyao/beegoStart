/*
 * @Author: zhenwei zhang 
 * @Date: 2018-03-29 18:22:03 
 * @Last Modified by: zhenwei zhang
 * @Last Modified time: 2018-04-16 19:37:52
 */
 
 package models

 type Oauth struct{
	 Id int `json:"id"`
	 Uid int `json:"uid"`
	 GithubId string `orm:"size(255);null" json:"githubId"`
	 Tencent string `orm:"size(255);null" json:"tencent"`
	 Wechat string `orm:"size(255);null" json:"wechat"`
	 Gitee string `orm:"size(255);null" json:"gitee"`
	 Google string `orm:"size(255);null" json:"google"`
	 Facebook string `orm:"size(255);null" json:"facebook"`
	 Weibo string `orm:"size(255);null" json:"weibo"`
	 Baidu string `orm:"size(255);null" json:"baidu"`
 }