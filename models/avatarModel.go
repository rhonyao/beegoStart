/*
 * @Author: zhenwei zhang 
 * @Date: 2018-03-29 18:22:03 
 * @Last Modified by:   zhenwei zhang 
 * @Last Modified time: 2018-03-29 18:22:03 
 */
 
package models

type Avatar struct{
	Id int `json:"id"`
	Uid int `json:"uid"`
	AvatarUrl string `json:"avatarUrl"`
	CoverUrl string `json:"coverUrl"`
}