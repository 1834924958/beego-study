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
    //前台用户显示
    beego.Router("/admin/hlist", &controllers.AuserController{},"*:Hlist")
    //前台用户编辑页面
    beego.Router("/admin/hedit/?:id", &controllers.AuserController{},"Get:Hedit")
    //前台用户的编辑删除操作
    beego.Router("/home/hstatus", &controllers.AuserController{},"Post:Hstatus")
    //执行管理员用户的操作
    beego.Router("/admin/alist", &controllers.AuserController{},"*:Alist")
    beego.Router("/admin/aedit/?:id", &controllers.AuserController{},"Get:Aedit")
    beego.Router("/home/astatus", &controllers.AuserController{},"Post:Astatus")
}
