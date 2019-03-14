package controllers

import (
	"fmt"
)

type AindexController struct{
	BaseAdminController
}
//定义后台访问的操作
func (c *AindexController) Index(){
	c.TplName = "admin/index/index.html"
}

//定义后台首页
func (c *AindexController) Main(){
	fmt.Println("测试112")
	c.TplName = "admin/page/main.html"
}