/*
 * @Author: zhenwei zhang 
 * @Date: 2018-03-29 18:22:25 
 * @Last Modified by:   zhenwei zhang 
 * @Last Modified time: 2018-03-29 18:22:25 
 */

package models

import (
	"time"
)

type Storage struct{
	Id int `json:"id"`
	Uid int `json:"uid"`
	FileName string `json:"fileName"`
	FileType string `json:"fileType"`
	FileSize string `json:"fileSize"`
	FilePath string `json:"filePath"`
	UploadTime time.Time `json:"registryDate"`
	StorageEngine string `json:"-"`
}