package controllers

import (
	"fmt"
)

type IndexController struct{
	BaseHomeController
}
// func (c *IndexController) Get(){
// 	c.TplName = "home/index/index.html"
// }


//定义首页访问操作
func (c *IndexController) Index(){
	c.TplName = "home/index/index.html"
}
//测试数据库的操作
func (c *IndexController) Ches(){
	fmt.Println("车上")
	c.Ctx.WriteString("hello,执行数据测试的操作")
}