package controllers

import (
	"github.com/astaxie/beego"
	"beego-study/models"
)
type BaseAdminController struct{
	beego.Controller
}
//定义返回信息
func (p *BaseAdminController) jsonResult(code,msg string,obj interface{}){
	r := &models.JsonResult{code,msg,obj}
	p.Data["json"] = r
	p.ServeJSON()
	p.StopRun()
}