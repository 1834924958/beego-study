package controllers

import (
	"github.com/astaxie/beego"
	"beego-study/models"
)
type BaseHomeController struct{
	beego.Controller
}

func (p *BaseHomeController) History(msg string, url string) {
	if url == ""{
		p.Ctx.WriteString("<script>alert('"+msg+"');window.history.go(-1);</script>")
		p.StopRun()
	}else{
		p.Redirect(url,302)
	}
}
//定义返回信息
func (p *BaseHomeController) jsonResult(code,msg string,obj interface{}){
	r := &models.JsonResult{code,msg,obj}
	p.Data["json"] = r
	p.ServeJSON()
	p.StopRun()
}