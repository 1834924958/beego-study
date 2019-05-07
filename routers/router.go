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
    //管理员登录和退出
    beego.Router("/admin/login",&controllers.AloginController{},"*:Login")
    beego.Router("/admin/logout",&controllers.AloginController{},"*:Logout")
    //内容管理的操作
    //广告列表
    beego.Router("/admin/adver",&controllers.ArticleController{},"*:Adver")
    //广告新增,编辑操作
    beego.Router("/admin/adveredit",&controllers.ArticleController{},"Get:AdverEdit")
    //广告状态值修改的操作
    beego.Router("/admin/adverstatus",&controllers.ArticleController{},"Post:AdverStatus")
    beego.Router("/admin/adverimage",&controllers.ArticleController{},"Post:AdverImage")
    beego.Router("/admin/aimage",&controllers.ArticleController{},"*:Aimage")
    //执行测试配置的操作
    beego.Router("/admin/atest",&controllers.AtestController{},"*:Atest")
    beego.Router("/admin/atimage",&controllers.AtestController{},"*:Aimage")
    beego.Router("/admin/aupload",&controllers.AtestController{},"*:Upload")

}
