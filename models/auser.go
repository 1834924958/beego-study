package models

import (
	"time"
)

type Auser struct {
	Id 			int
	Username 	string
	Mobile   	string
	Email    	string
	Password 	string
	Status 		int
	Ipaddress	string
	Createtime  time.Time
}
//管理员用户表
func (m *Auser) TableName() string {
	return TableName("auser")
}