package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"time"
)

type UserInfo struct {
	Id         	int64   	`orm:"auto"`
	UserName   	string  	`orm:"size(255)"`
	Password   	string  	`orm:"size(255)"`
	Name 	   	string		`orm:"size(255)"`
	BirthDate 	string 		`orm:"size(255)"`
	Gender		int8
	Email		string		`orm:"size(255)"`
	Phone		string		`orm:"size(255)"`
	Status		int 
	CreateTime  time.Time
	UpdateTime  time.Time
}
//执行获取数据
func QueryAllUserInfo() (dataList [] interface{},err error) {
	var list []UserInfo
	o := orm.NewOrm()
	o.Using("beegotest")
	qs := o.QueryTable(new(UserInfo))
	//查询数据
	if _,err = qs.All(&list);err == nil {
		for _,v := range list {
			dataList = append(dataList,v)
		}
		return dataList,nil
	}
	return nil,err
}