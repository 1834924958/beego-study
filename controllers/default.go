package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	fmt.Println(c)
	c.TplName = "home/index/index.html"
}

func (c *MainController) info(){
	c.Data["ches"] = "ches123"
	fmt.Println(c)
}