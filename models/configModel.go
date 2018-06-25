/*
 * @Author: zhenwei zhang 
 * @Date: 2018-03-29 18:22:09 
 * @Last Modified by:   zhenwei zhang 
 * @Last Modified time: 2018-03-29 18:22:09 
 */

 package models

type Config struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Value string `json:"value"`
}