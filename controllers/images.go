package controllers

import (
	//	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
	"sfapi/models"
)

var (
	//设置上传到的空间
	bucket = "broadroad"
)

//构造返回值字段
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

// Operations about Imagess
type ImagesController struct {
	beego.Controller
}

func downloadUrl(domain, key string) string {
	//调用MakeBaseUrl()方法将domain,key处理成http://domain/key的形式
	baseUrl := kodo.MakeBaseUrl(domain, key)
	policy := kodo.GetPolicy{}
	//生成一个client对象
	c := kodo.New(0, nil)
	//调用MakePrivateUrl方法返回url
	return c.MakePrivateUrl(baseUrl, &policy)
}

func (u *ImagesController) Get() {
	id, _ := u.GetInt64(":imgId")
	images := &models.Images{Id: id}
	if id != 0 {
		images, err := images.GetImage()
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = images
		}
	}
	u.ServeJSON()
}

func (u *ImagesController) GetAll() {
	images := &models.Images{}
	image, err := images.GetImages()
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = image
	}
	u.ServeJSON()
}

func (u *ImagesController) Post() {
	f, h, _ := u.GetFile("file")
	fmt.Println(h.Filename)
	u.SaveToFile("file", "/tmp/"+h.Filename)
	f.Close()
	jsonResult := &models.JsonResult{}
	picname := h.Filename
	userid, _ := u.GetInt64("userid")
	point, _ := u.GetInt64("point")
	description := u.GetString("description")
	key := u.GetString("key")
	var images models.Images
	images.Picname = picname
	images.Userid = userid
	images.Key = key
	images.Point = point
	images.Description = description
	err := images.UploadImage()
	if err != nil {
		jsonResult.Message = "保存失败"
	} else {
		jsonResult.Message = "保存成功"
		jsonResult.Data = images
		jsonResult.Success = true
	}
	//初始化AK，SK
	conf.ACCESS_KEY = "3TPl1uGxc3fT31jsNFQQ0_V0N2_mjsh_XnJbV5qh"
	conf.SECRET_KEY = "Qlgx61Y-_J7wUuge2yUXsqKnqzJdRSd01oiusHNB"
	c := kodo.New(0, nil)
	policy := &kodo.PutPolicy{
		Scope: bucket,
		//设置Token过期时间
		Expires: 3600,
	}
	//生成一个上传token
	token := c.MakeUptoken(policy)

	fmt.Println("start")
	//构建一个uploader
	zone := 0
	uploader := kodocli.NewUploader(zone, nil)
	var ret PutRet
	//设置上传文件的路径
	filepath := "/tmp/" + h.Filename
	//调用PutFileWithoutKey方式上传，没有设置saveasKey以文件的hash命名
	res := uploader.PutFile(nil, &ret, token, h.Filename, filepath, nil)
	//打印返回的信息
	fmt.Println(ret)
	//打印出错信息
	if res != nil {
		jsonResult.Message = "保存失败"
	}

	u.Data["json"] = jsonResult
	u.ServeJSON()
}
