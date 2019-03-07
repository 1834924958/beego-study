package controllers

import (
	"fmt"
)

type HLoginController struct{
	BaseHomeController
}
// func (c *IndexController) Get(){
// 	c.TplName = "home/index/index.html"
// }


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
			c.Ctx.WriteString("使用Post进行提交")
	// if c.Ctx.Rquest.Method == "POST"{
	// 		c.Ctx.WriteString("使用Post进行提交")
	// 	}else{
	// 		c.Ctx.WriteString("无效的提交方式")
	// 	}
}