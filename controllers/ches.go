package controllers

import (
	"github.com/astaxie/beego"
)

type ChesController struct{
	beego.Controller
}
// type User struct {
// 	Name string
// 	Age  int
// 	Sex  string
// }

func (c *ChesController) Get(){
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	// user:= &User{"晨曦",20,"m"}
	c.Ctx.WriteString("hello,执行demo测试")
	// c.Data["xml"] =user 
	// c.ServeXML()
}