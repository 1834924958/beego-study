package controllers

import (
	"time"
	"strconv"
	"beego-study/models"
	"github.com/astaxie/beego/orm"
)
type AuserController struct{
	BaseAdminController
}
//定义前端用户详情
func (c *AuserController) Hlist() {
	o := orm.NewOrm()
	hlists := [] *models.Huser{}
	query := o.QueryTable(new(models.Huser).TableName())
	//执行限制条件的操作
	search := c.GetString("search")
	// c.Ctx.WriteString(search)
	if search != ""{
		query.Filter("Username",search).All(&hlists)
	}else{
		query.OrderBy("-id").All(&hlists)
	}
	
	// lists.OrderBy("-id").All(&hlists)
	c.Data["hlists"] =  hlists
	c.TplName = "admin/user/hlist.html"
		// c.Ctx.WriteString(search)
}
//定义前端用户编辑操作
func (c *AuserController) Hedit() {
	id := c.Input().Get("id")
	if id != "" {
		o := orm.NewOrm()
		//执行字符串转换成整型
		intId,_ := strconv.Atoi(id)
		hlist := models.Huser{Id:intId}
		o.Read(&hlist)
		c.Data["hlist"] = hlist
	}
	c.TplName = "admin/user/hedit.html"
}
//执行前端用户修改的操作(编辑,删除)
func (c *AuserController) Hstatus() {
	if c.Ctx.Request.Method == "POST" {
		types := c.GetString("types")
		id := c.GetString("Id")
		o := orm.NewOrm()
		cases := models.Huser{}
		intId,_ := strconv.Atoi(id)
		if (types != "" && id != ""){
			switch types {
				case "hedit":
					// cases.Status   = 0
					cases.Username = c.GetString("Username")
					cases.Mobile = c.GetString("Mobile")
					// cases.Addtime  = time.Now()
					cases.Endtime  = time.Now()
					cases.Id       = intId
					status,_ := strconv.Atoi(c.GetString("Status"))
					cases.Status   = status
					if _,err := o.Update(&cases);err != nil{
						// c.Ctx.WriteString("数据更新失败")
						c.jsonResult("0","数据更新失败","")

					}else{
						// c.Ctx.WriteString("更新成功")
						c.jsonResult("1","更新成功111","")

					}
					// Huser := models.Huser{}
					// Huser.Id = 6
					// Huser.Username = "chenxi123"
					// Huser.Mobile   = "17839906238"
					// Huser.Status   = 0
					// Huser.Addtime   = time.Now()
					// o.Update(&cases)
					// c.Ctx.WriteString("修改操作")
				case "hdelete":
					cases.Id       = intId
					if _,err := o.Delete(&cases);err != nil{
							c.jsonResult("0","删除失败","")
						}else{
							c.jsonResult("1","删除成功","")
						}
				default:
					c.jsonResult("0","操作失败","")
			}
		}else{
			c.jsonResult("0","参数错误","")

		}
	}else{
		 c.jsonResult("0","无效的请求方式","")
	}
}
//管理员列表的处理
func (c *AuserController) Alist() {
	o := orm.NewOrm()
	alists := [] *models.Auser{}
	o.QueryTable(new(models.Auser).TableName()).OrderBy("-id").All(&alists)
	c.Data["alists"] =  alists
	c.TplName = "admin/user/alist.html"
}
//执行编辑修改的页面
func (c *AuserController) Aedit() {
	id := c.Input().Get("id")
	if id != "" {
		o := orm.NewOrm()
		//执行字符串转换成整型
		intId,_ := strconv.Atoi(id)
		alist := models.Auser{Id:intId}
		o.Read(&alist)
		c.Data["alist"] = alist
	}
	c.TplName = "admin/user/aedit.html"
}
//执行管理员的添加，删除，编辑的操作
func (c *AuserController) Astatus() {
	if c.Ctx.Request.Method == "POST" {
		types := c.GetString("types")
		id := c.GetString("Id")
		intId,_ := strconv.Atoi(id)
		o := orm.NewOrm()
		cases := models.Auser{}
		cases.Username = c.GetString("Username")
		cases.Mobile = c.GetString("Mobile")
		status,_ := strconv.Atoi(c.GetString("Status"))
		cases.Status   = status
		if (types != ""){
			switch types {
				case "aedit":
					if id != ""{
						cases.Id       = intId
						cases.Createtime   = time.Now()
						if _,err := o.Update(&cases);err != nil{
							c.jsonResult("0","修改失败","")

						}else{
							c.jsonResult("1","修改成功","")
						}
					}else{
							c.jsonResult("0","无效的ID","")
					}
				case "aregister":
					cases.Createtime   = time.Now()
					if _,err := o.Insert(&cases);err != nil{
						c.jsonResult("0","管理员添加失败","")
					}else{
						c.jsonResult("1","管理员添加成功","")
					}
				case "hdelete":
					cases.Id       = intId
					if _,err := o.Delete(&cases);err != nil{
							c.jsonResult("0","删除失败","")
						}else{
							c.jsonResult("1","删除成功","")
						}
				default:
					c.jsonResult("0","操作失败","")
			}
		}else{
			c.jsonResult("0","参数错误","")

		}
	}else{
		 c.jsonResult("0","无效的请求方式","")
	}
}