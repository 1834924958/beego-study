package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type HLoginController struct{
	BaseHomeController
}

type Huser struct {
    Id          int       `orm:"auto"`
    UserName    string      `orm:"size(150)"`
    Mobile      string      `orm:"size(100)"`
    Status      int     
    AddTime     int
}
func init() {
	orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/beegotest?charset=utf8",30)
	orm.RegisterModel(new(Huser))
	orm.RunSyncdb("default",false,true)
}
// func (c *IndexController) Get(){
// 	c.TplName = "home/index/index.html"
// 引入包:// "beego-study/models"
// }
//定义注册的操作
func (c *HLoginController) Register(){
	o := orm.NewOrm()
	// huser := Huser{UserName:"qwe12",Mobile:"9527",Status:2}
	////insert
	// id,err := o.Insert(&huser)
	////update
	// huser.UserName = "chenxi"
	// huser.Id = 3
	// id,err := o.Update(&huser)
	////delete
	hu := Huser{Id:4}
	// id,err := o.Delete(&hu)
	////read 
	err := o.Read(&hu)
	// fmt.Println(err)
	// user:= Huser{UserName:"qwe12123",Mobile:"9527",Status:2}
	user:= err
	c.Data["json"] =user  
	c.ServeJSON()
	// dataList,err  := models.QueryAllUserInfo()
	// if err == nil {
	// 	c.Data["json"] = dataList
	// }else{
	// c.Ctx.WriteString("hello,执行数据测试的操作11")
	// }
	c.Ctx.WriteString("ches ")
	// c.TplName = "home/login/register.html"
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
// func (c *HLoginController) Regform(){
// 	m := models.Huser{}
// 	o := orm.NewOrm()
// 	types    := c.GetString("type")
// 	if c.Ctx.Request.Method == "POST";types == "register"{
// 		m.addtime = time.Now()
// 		if m.Id == 0{
// 			if err := c.ParseForm(&m);err == nil{
// 				if _,err = o.insert(&m);err == nil{
// 					c.Data['json'] = "添加成功"
// 				}else{
// 					c.Data['json'] = "添加失败"
// 				}
// 			}else{
// 				c.Data["json"] = "error"
// 			}
// 		}
// 		//执行添加数据
// 	}else{
// 		c.Ctx.WriteString("无效的提交方式")
// 	}
// }
////执行注册时添加数据的操作
func (c *HLoginController) Regform(){
	//查询数据
	types    := c.GetString("type")
	if (c.Ctx.Request.Method == "POST" && types == "register") {
		c.Ctx.WriteString("post提交方式")
	}else{
		c.Ctx.WriteString("无效的提交方式")
	}
	// if c.Ctx.Request.Method == "POST";types == "register"{
	// 	c.Ctx.WriteString("post提交方式")
	// 	//执行添加数据
	// }else{
	// 	c.Ctx.WriteString("无效的提交方式")
	// }

}