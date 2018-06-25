/*
 * @Author: zhenwei zhang 
 * @Date: 2018-03-29 18:22:18 
 * @Last Modified by: zhenwei zhang
 * @Last Modified time: 2018-04-15 15:21:57
 */

package models

import (
	"time"
)

type Posts struct{
	Id int `json:"id"`
	Uid int `json:"uid"`
	CreateDate time.Time `json:"createDate"`
	Title string `json:"title"`
	Content string `json:"content"`
	Delete bool `json:"-"` /*是否删除*/
	ViewCounts int `json:"viewCounts"` /*阅读次数*/
	Mention string `json:"mention"` /*提醒谁观看*/
	Group int `json:"group"` /*关联的社群ID，可为空*/
	Public bool `json:"public"` /*是否公开*/
}