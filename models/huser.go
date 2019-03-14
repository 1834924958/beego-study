package models

import (
	"time"
)

type Huser struct {
	Id int
	Username string
	Mobile string 
	Status int
	Addtime time.Time
}

func (m *Huser) TableName() string {
	return TableName("huser")
}