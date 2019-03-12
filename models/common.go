package models

//定义类
type JsonResult struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Obj	 interface{} `json:"obj"`
}