package models

import (
	"time"
)

type Huser struct {
	Id int
	Username string
	Mobile string 
	Status int
	Addtime time.Time `orm:"auto_now_add;type(datetime)"`
	Endtime time.Time `orm:"auto_now_add;type(datetime)"`
}

func (m *Huser) TableName() string {
	return TableName("huser")
}