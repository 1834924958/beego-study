package controllers

import (
	"fmt"
	"github.com/beego-study/models"
)

type HLoginController struct{
	BaseHomeController
}
// func (c *IndexController) Get(){
// 	c.TplName = "home/index/index.html"
// }
type user struct{
	Name string  `form:"username"`
	Phone string `form:"phone"`
	Type  string `form:"register"`
}
//定义注册的操作
func (c *HLoginController) Register(){
	c.TplName = "home/login/register.html"
}
//定义登录的操作
func (c *HLoginController) Login(){
	c.TplName = "home/login/login.html"
}
//测试数据库的操作
func (c *HLoginController) Ches(){
	fmt.Println("车上")
	c.Ctx.WriteString("hello,执行数据测试的操作11")
}
//执行注册添加数据的操作
func (c *HLoginController) Regform(){
	m := models.Huser{}
	if c.Ctx.Request.Method == "POST"{
		c.Ctx.WriteString("Post提交111")
		// username := c.GetString("username")
		// mobile   := c.GetString("mobile")
		// types    := c.GetString("type")
		// var huser = user
		// huser.Username = username
		// huser.Mobile = mobile
		// id,err := o.Insert(&huser)
		//执行添加数据
	}else{
		c.Ctx.WriteString("无效的提交方式")
	}
}