package controllers

import (
	"fmt"
	"time"
	"path"
	"os"
	"math/rand"
	"crypto/md5"
	"strconv"
	"beego-study/models"
	"github.com/astaxie/beego/orm"
)

type ArticleController struct {
	BaseAdminController
}

//执行配置广告管理的操作
func (c *ArticleController) Adver(){
	////原始数据
	// o := orm.NewOrm()
	// adver := [] *models.Adver{}
	// o.QueryTable(new(models.Adver).TableName()).OrderBy("-id").All(&adver)
	// c.Data["adver"] =  adver
	/*
	 *执行第一种分页数据的处理(有点问题)
	 *
	*/
	// var (
	// 	page 		int
	// 	pagesize 	int =3
	// 	offset		int
	// )
	// keyword := c.GetString("title")
	// list :=  [] *models.Adver{}
	// if page,_ = c.GetInt("page");page < 1{
	// 	page = 1
	// }
	// offset = (page - 1) * pagesize
	// o := orm.NewOrm()
	// query := o.QueryTable(new(models.Adver).TableName())
	// //查询所有数据
	// count,_ := query.Count()
	// if count > 0 {
	// 	query.OrderBy("-id").Limit(pagesize,offset).All(&list)
	// }
	// c.Data["keyword"] = keyword
	// c.Data["count"] = count
	// c.Data["adver"] = list
	// c.Data["pagebar"] = util.NewPager1(page,int(count),pagesize,fmt.Sprintf("/admin/index.html?keyword=%s", keyword), true)
	/*
	 *执行第二种分页
	 *
	*/
	o := orm.NewOrm()
	adver := [] *models.Adver{}
	pa,err := c.GetInt("page")
	if err != nil{
		println(err)
	}
	//限制每页个数
	pre_page := 2
	//获取分页数据
	query := o.QueryTable(new(models.Adver).TableName())
	totals,_ := query.Count()
	res := models.Paginator(pa,pre_page,totals)
	// userlist := models.LimitList(3,pa)
	query.OrderBy("-id").Limit(2,pa).All(&adver)
	c.Data["adver"] = adver
	c.Data["paginator"] = res
	c.Data["totals"] = totals
	// c.Data["json"] = res
	// c.ServeJSON()
	c.TplName = "admin/article/adver.html"
}
//执行配置广告新增编辑的页面的操作
func (c *ArticleController) AdverEdit(){
	id := c.Input().Get("id")
	if id != "" {
		o := orm.NewOrm()
		//执行字符串转换成整型
		intId,_ := strconv.Atoi(id)
		alist := models.Adver{Id:intId}
		o.Read(&alist)
		c.Data["alist"] = alist
	}
	c.TplName ="admin/article/adveredit.html"
}
//执行配置广告管理状态的操作
func (c *ArticleController) AdverStatus() {
		if c.Ctx.Request.Method == "POST" {

		types := c.GetString("types")
		id := c.GetString("Id")
		intId,_ := strconv.Atoi(id)
		o := orm.NewOrm()
		cases := models.Adver{}
		// o.QueryTable(new(models.Adver).TableName()).Filter("Id", intId).One(&cases)
		// c.Data["json"] = cases
		// c.ServeJSON()
		cases.Title = c.GetString("Title")
		cases.Url = c.GetString("Url")
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
					//执行检测用户是否存在不存在
					o.QueryTable(new(models.Adver).TableName()).Filter("Id", intId).One(&cases)
					if (c.GetString("Ttile") == "" && cases.Id != 0) {
						c.jsonResult("0","标题不能为空","")
					}else if (cases.Title == c.GetString("Title") && cases.Id != 0) {
						c.jsonResult("0","标题已存在","")
					}else{
						cases.Createtime   = time.Now()
						if _,err := o.Insert(&cases);err != nil{
							c.jsonResult("0","广告添加失败","")
						}else{
							c.jsonResult("1","广告添加成功","")
						}
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
//执行上传广告图的操作
func (c *ArticleController) AdverImage(){
	// f, h, err := c.GetFile("fileImageUrl")
	// f, h, err := c.GetFile("Images")
	a := c.GetString("Formdata")
	c.jsonResult("1","a",a)
	// c.jsonResult("1","f",f)
	// c.jsonResult("1","h",h)
	// c.jsonResult("1","err",err)
}
//执行上传页面的操作
func (c *ArticleController) Aimage(){
	if c.Ctx.Request.Method == "POST" {
		///第一种：
		// f, h, err := c.GetFile("Images")
		// if err != nil {
		// 	c.jsonResult("0","上传失败", "")
		// }
		// defer f.Close()
		// filePath := "static/upload/adver/" + h.Filename
		// // 保存位置在 static/upload, 没有文件夹要先创建
		// c.SaveToFile("Images", filePath)
		// c.jsonResult("1","上传成功", "/"+filePath)
		////第二种
		    f, h, _ := c.GetFile("Images")//获取上传的文件
		    ext := path.Ext(h.Filename)
		    //验证后缀名是否符合要求
		    var AllowExtMap map[string]bool = map[string]bool{
		        ".jpg":true,
		        ".jpeg":true,
		        ".png":true,
		    }
		    if _,ok:=AllowExtMap[ext];!ok{
		        c.Ctx.WriteString( "后缀名不符合上传要求" )
		        return 
		    }
		    //创建目录
		    uploadDir := "static/upload/adver" + time.Now().Format("2006/01/02/")
		    err := os.MkdirAll( uploadDir , 777)
		    if err != nil {
		        c.Ctx.WriteString( fmt.Sprintf("%v",err) )
		        return 
		    }
		    //构造文件名称
		    rand.Seed(time.Now().UnixNano())
		    randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000 )
		    hashName := md5.Sum( []byte( time.Now().Format("2006_01_02_15_04_05_") + randNum ) )

		    fileName := fmt.Sprintf("%x",hashName) + ext
		    //this.Ctx.WriteString(  fileName )

		    fpath := uploadDir + fileName
		    defer f.Close()//关闭上传的文件，不然的话会出现临时文件不能清除的情况
		    err = c.SaveToFile("myfile", fpath)
		    if err != nil {
		        c.Ctx.WriteString( fmt.Sprintf("%v",err) )
		    }
		    c.Ctx.WriteString( "上传成功~！！！！！！！" )

	}else{
		c.TplName = "admin/article/aimage.html"	
	}
}
