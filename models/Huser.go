package models

import(
	"github.com/astaxie/beego/orm"
	"time"
)
//Huser用户
type Huser struct {
	Id  int `form:Id`
	Username string `form:Username`
	Mobile   string  `form:Mobile` 
	addtime  time.Time
}