/*
 * @Author: zhenwei zhang 
 * @Date: 2018-03-29 18:20:33 
 * @Last Modified by: zhenwei zhang
 * @Last Modified time: 2018-06-24 00:31:23
 */
 
package controllers

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"regexp"
	"bytes"  
    "github.com/mingzhehao/goutils/filetool"
    "io"  
    "log"  
    "mime/multipart"  
    "net/url"  
    "os"  
    "path/filepath"  
    "strings"  
    "time"
    "github.com/astaxie/beego"
)

const (  
    LOCAL_FILE_DIR    = "uploads/file" 
    MIN_FILE_SIZE     = 1
    MAX_FILE_SIZE     = 5000000
    IMAGE_TYPES       = "(zip|rar|bmp|exe|doc|docx|csv|txt|mp3|mp4|jpg|gif|p?jpeg|(x-)?png)"  
    ACCEPT_FILE_TYPES = IMAGE_TYPES  
    EXPIRATION_TIME   = 300
    THUMBNAIL_PARAM   = "=s80"  
)  

var (
    aliyunBucket    = beego.AppConfig.String("oss::aliyunBucket")
	aliyunSecretKey = beego.AppConfig.String("oss::aliyunSecretKey")
	aliyunAccessId  = beego.AppConfig.String("oss::aliyunAccessId")
	aliyunEndpoint = beego.AppConfig.String("oss::aliyunEndpoint")
	aliyunInternalEndpoint = beego.AppConfig.String("oss::aliyunInternalEndpoint")
    imageTypes      = regexp.MustCompile(IMAGE_TYPES)  
    acceptFileTypes = regexp.MustCompile(ACCEPT_FILE_TYPES)  
) 

type OSSController struct {
	BaseController
}

type ResponseData struct{
	ImageUrl string `json:"imageUrl"`
}

type Sizer interface {  
    Size() int64  
}   

type FileInfo struct {  
    Url          string `json:"url,omitempty"`  
    ThumbnailUrl string `json:"thumbnailUrl,omitempty"`  
    Name         string `json:"name"`  
    Type         string `json:"type"`  
    Size         int64  `json:"size"`  
    Error        string `json:"error,omitempty"`  
    DeleteUrl    string `json:"deleteUrl,omitempty"`  
    DeleteType   string `json:"deleteType,omitempty"`  
}  
  
func (this *OSSController) Upload() {  
    this.ApiLogin()
    f, h, err := this.GetFile("file")  
    defer f.Close()  
    if err != nil {  
        fmt.Println("getfile err ", err)   
        this.DisplayJson(0, "未选择文件" , nil);
        return  
    } else {  
        var imageUrl string  
        ext := filetool.Ext(h.Filename)  
        fi := &FileInfo{  
            Name: h.Filename,  
            Type: ext,  
        }  
        if !fi.ValidateType() {  
            this.DisplayJson(0, "文件类型暂不支持上传" , nil);
            return  
        }  
        var fileSize int64  
        if sizeInterface, ok := f.(Sizer); ok {  
            fileSize = sizeInterface.Size()  
            fmt.Println(fileSize)  
        }  
        fileExt := strings.TrimLeft(ext, ".")  
		fileSaveName := fmt.Sprintf("%s_%d%s", fileExt, time.Now().Unix(), ext)  
		t := time.Now()
	    fileStamp := "/" + fmt.Sprintf("%d",t.Year()) + fmt.Sprintf("%d",t.Month()) + fmt.Sprintf("%d",t.Day())
		imgPath := fmt.Sprintf("%s/%s", LOCAL_FILE_DIR + fileStamp, fileSaveName)  
        filetool.InsureDir(LOCAL_FILE_DIR + fileStamp)
		this.SaveToFile("file", imgPath)
		if this.OSSUpload(imgPath) == true {
			imageUrl = "/" + imgPath
		}else{
			if err == nil {
				if filetool.IsExist(imgPath) {
					filetool.Unlink(imgPath)
                    this.DisplayJson(0, "OSS同步失败，文件暂被移除" , nil);
				}else{
                    this.DisplayJson(0, "文件移除失败" , nil);
				}
			}
			imageUrl = ""
        }
        var Res ResponseData
        Res.ImageUrl = imageUrl
        this.DisplayJson(1, "上传成功" , &Res);  
        return  
    }  
}

func (this *OSSController) OssDelete(object string) error {
	client, err := oss.New(aliyunEndpoint, aliyunAccessId, aliyunSecretKey)
    if err != nil {
        fmt.Print(err)
    }
    bucket, err := client.Bucket(aliyunBucket)
    if err != nil {
        fmt.Print(err)
    }
    return bucket.DeleteObject(object)
}
  
func (this *FileInfo) ValidateType() (valid bool) {  
    if acceptFileTypes.MatchString(this.Type) {  
        return true  
    }  
    this.Error = "Filetype not allowed"  
    return false  
}  
  
func (this *FileInfo) ValidateSize() (valid bool) {  
    if this.Size < MIN_FILE_SIZE {  
        this.Error = "File is too small"  
    } else if this.Size > MAX_FILE_SIZE {  
        this.Error = "File is too big"  
    } else {  
        return true  
    }  
    return false  
}  

func (this *OSSController) List() {
	this.DisplayJson(0, "列表请求暂不开放" , nil);
}

func (c *OSSController) OSSUpload(imgPath string) bool {
	result := false
	client, err := oss.New(aliyunEndpoint, aliyunAccessId, aliyunSecretKey)
    if err != nil {
		fmt.Println(err)
		fmt.Println("发生错误的文件地址：" + imgPath)
		return result
    }
    bucket, err := client.Bucket(aliyunBucket)
    if err != nil {
		fmt.Println(err)
		fmt.Println("发生错误的文件地址：" + imgPath)
		return result
    }
    err = bucket.PutObjectFromFile(imgPath, imgPath)
    if err != nil {
		fmt.Println(err)
		fmt.Println("发生错误的文件地址：" + imgPath)
		return result
	}
	result = true
	return result
}
  
func check(err error) {  
    if err != nil {  
        panic(err)  
    }  
}  
  
func escape(s string) string {  
    return strings.Replace(url.QueryEscape(s), "+", "%20", -1)  
}  
  
func getFormValue(p *multipart.Part) string {  
    var b bytes.Buffer  
    io.CopyN(&b, p, int64(1<<20)) 
    return b.String()  
}  

func substr(s string, pos, length int) string {  
    runes := []rune(s)  
    l := pos + length  
    if l > len(runes) {  
        l = len(runes)  
    }  
    return string(runes[pos:l])  
}  
  
func getParentDirectory(dirctory string) string {  
    return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))  
}  
  
func getCurrentDirectory() string {  
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))  
    if err != nil {  
        log.Fatal(err)  
    }  
    return strings.Replace(dir, "\\", "/", -1)  
}  