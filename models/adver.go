package models

import (
	"time"
)

type Adver struct{
	Id		int
	Title	string
	Url 	string
	Createtime time.Time
	Status   int
	images   string
}

func (m *Adver) TableName() string {

	return TableName("adver")
}