package controllers

import (
	"github.com/astaxie/beego"
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