package controllers

import (
	"math/rand"
	"crypto/md5"
	"time"
	"os"
	"path"
	"fmt"
	"github.com/astaxie/beego"
)

type AtestController struct {
	BaseAdminController
}

//执行测试配置的操作
func (c *AtestController) Atest() {
	c.TplName = "admin/article/atest.html"
}
//执行上传配置的操作
func (c *AtestController) Aimage() {
	if c.Ctx.Request.Method == "POST" {
		//检测信息
		// a := c.GetString("Images")
		// a := c.GetString("file")
		// c.jsonResult("1","a",a)
		f, h, _ := c.GetFile("Images")//获取上传的文件
		    ext := path.Ext(h.Filename)
		    //验证后缀名是否符合要求
		    var AllowExtMap map[string]bool = map[string]bool{
		        ".jpg":true,
		        ".jpeg":true,
		        ".png":true,
		    }
		    if _,ok:=AllowExtMap[ext];!ok{
		        c.Ctx.WriteString( "后缀名不符合上传要求" )
		        return 
		    }
		    //创建目录
		    uploadDir := "static/upload/atest" + time.Now().Format("2006/01/02/")
		    err := os.MkdirAll( uploadDir , 777)
		    if err != nil {
		        c.Ctx.WriteString( fmt.Sprintf("%v",err) )
		        return 
		    }
		    //构造文件名称
		    rand.Seed(time.Now().UnixNano())
		    randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000 )
		    hashName := md5.Sum( []byte( time.Now().Format("2006_01_02_15_04_05_") + randNum ) )

		    fileName := fmt.Sprintf("%x",hashName) + ext
		    //this.Ctx.WriteString(  fileName )

		    fpath := uploadDir + fileName
		    defer f.Close()//关闭上传的文件，不然的话会出现临时文件不能清除的情况
		    err = c.SaveToFile("myfile", fpath)
		    if err != nil {
		        c.Ctx.WriteString( fmt.Sprintf("%v",err) )
		    }
		    c.Ctx.WriteString( "上传成功~！！！！！！！" )
	}else{
		c.jsonResult("0","无效的请求方式","")
	}
}
func (c *AtestController) Upload() {

	if c.Ctx.Request.Method == "POST" {
		filepath := "static/upload/" + beego.Date(time.Now(), "Ymd") + "/"
		_, err := os.Stat(filepath)
		if err != nil {
			err := os.MkdirAll(filepath, 0777)
			if err != nil {
				c.jsonResult("0","上传失败","")
			}
		}
		c.UploadFile("upfile", filepath)
	}else{
		c.jsonResult("0","无效的请求方式","")
	}
}
