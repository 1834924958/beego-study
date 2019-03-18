package controllers

import (
	"github.com/astaxie/beego"
	"beego-study/models"
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