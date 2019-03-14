package routers

import (
	"beego-study/controllers"
	"github.com/astaxie/beego"
)


func init() {
    // beego.Router("/", &controllers.MainController{})
    beego.Router("/", &controllers.IndexController{},"*:Index")
    beego.Router("/dbches", &controllers.IndexController{},"*:Ches")
    //自定义路由，便于显示的操作
    // beego.Router("/home",&home.IndexController{})
    beego.Router("/ches",&controllers.ChesController{})
    beego.Router("/info",&controllers.MainController{})
    //前台用户登录注册
    beego.Router("/home/register", &controllers.HLoginController{},"*:Register")
    beego.Router("/home/login", &controllers.HLoginController{},"*:Login")
    beego.Router("/home/regform", &controllers.HLoginController{},"Post:Regform")
    //后台用户处理的操作
    beego.Router("/admin/index", &controllers.AindexController{},"*:Index")
    beego.Router("/admin/main", &controllers.AindexController{},"*:Main")
    beego.Router("/admin/info", &controllers.AindexController{},"*:Info")

}
