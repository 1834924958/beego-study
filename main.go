package main

import (
	_ "beego-study/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"beego-study/models"
)

func init() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
}
func main() {
	beego.Run()
}

