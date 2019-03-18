package controllers

import (
	"beego-study/models"
	"beego-study/util"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type AloginController struct {
	beego.Controller
}
//定义返回信息
func (p *AloginController) jsonResult(code,msg string,obj interface{}){
	r := &models.JsonResult{code,msg,obj}
	p.Data["json"] = r
	p.ServeJSON()
	p.StopRun()
}
//定义后台登录也访问的操作
func (c *AloginController) Login(){
	if c.Ctx.Request.Method == "POST" {
		username := c.GetString("username")
		o := orm.NewOrm()
		auser := models.Auser{}
		//执行检测用户存在不存在
		o.QueryTable(new(models.Auser).TableName()).Filter("Username",username).One(&auser)
		//执行检测的操作
		if auser.Id == 0 {
			c.jsonResult("0","无效的用户","")
		}else if util.Md5(c.GetString("password")) != auser.Password{
			c.jsonResult("0","用户名或密码错误","")
		}else if auser.Username != username{
			c.jsonResult("0","用户名或密码错误","")
		}else if auser.Status != 1{
			c.jsonResult("0","该账户已冻结","")
		}else{
			c.SetSession("auser", auser)
			c.jsonResult("1","登录成功11","")
		}
		// c.Data["json"] = auser
		// c.ServeJSON()
	}else{
		c.TplName = "admin/login/index.html"
	}
}
//执行后台退出的操作
func (c *AloginController) Logout(){
	c.DestroySession()
	c.TplName = "admin/login/index.html"
}