package controllers

import (
	"github.com/astaxie/beego"
	"beego-study/models"
	"beego-study/util"
	"path"
)
type BaseAdminController struct{
	beego.Controller
}
//执行初始化检测信息
func  (c *BaseAdminController) Prepare() {
	//获取session信息
	ausers := c.GetSession("auser")
	// c.jsonResult("1","打印session",ausers)
	if ausers == nil{
		// c.jsonResult("1","zhengc","")
		c.Redirect("/admin/login/",302)
	}
}
//定义返回信息
func (p *BaseAdminController) jsonResult(code,msg string,obj interface{}){
	r := &models.JsonResult{code,msg,obj}
	p.Data["json"] = r
	p.ServeJSON()
	p.StopRun()
}
//上传文件
func (c *BaseAdminController) UploadFile(filename string, filepath string) {
	f, h, err := c.GetFile(filename)

	out := make(map[string]interface{})
	if err != nil {
		out["msg"] = "文件读取错误"
	}
	var fileSuffix, newFile string
	fileSuffix = path.Ext(h.Filename) //获取文件后缀
	//验证后缀名是否符合要求
	var AllowExtMap map[string]bool = map[string]bool{
		".mp4":true,
		".mov":true,
		".rmvb":true,
	}
	if _,ok:=AllowExtMap[fileSuffix];!ok{
		c.jsonResult("0","后缀名不符合上传要求","")
		return 
	}
	//验证文件的大小是否符合要求
	if h.Size > 52428800 {
		c.jsonResult("0","文件不能超过50MB","")
		return
	}
	newFile = util.GetRandomString(8) + fileSuffix
	err = c.SaveToFile("upfile", filepath+newFile)
	if err != nil {
		out["msg"] = "文件保存错误"
	}
	defer f.Close()
	out["state"] = "SUCCESS"
	// out["fileSuffix"] = fileSuffix
	out["url"] = filepath + newFile
	// out["title"] = newFile
	// out["original"] = h.Filename
	// out["size"] = h.Size
	out["msg"] = "ok"
	c.Data["json"] = out
	c.jsonResult("1","上传成功",out)
	// c.ServeJSON()
	c.StopRun()
}
func (c *BaseAdminController) UploadFile111(filename string, filepath string) {
	f, h, err := c.GetFile(filename)

	out := make(map[string]interface{})
	if err != nil {
		out["msg"] = "文件读取错误"
	}
	var fileSuffix, newFile string
	fileSuffix = path.Ext(h.Filename) //获取文件后缀
	newFile = util.GetRandomString(8) + fileSuffix
	err = c.SaveToFile("upfile", filepath+newFile)
	if err != nil {
		out["msg"] = "文件保存错误"
	}
	defer f.Close()
	out["state"] = "SUCCESS"
	out["fileSuffix"] = fileSuffix
	out["url"] = filepath + newFile
	out["title"] = newFile
	out["original"] = h.Filename
	out["size"] = h.Size
	out["msg"] = "ok"
	c.Data["json"] = out
	c.ServeJSON()
	c.StopRun()
}