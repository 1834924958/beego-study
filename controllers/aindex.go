package controllers


type AindexController struct{
	BaseAdminController
}
//定义后台访问的操作
func (c *AindexController) Index(){
	c.TplName = "admin/index/index.html"
}

//定义后台首页
func (c *AindexController) Main(){
	c.TplName = "admin/page/main.html"
}
//定义管理员用户详情
func (c * AindexController) Info() {
	c.TplName = "admin/page/user/info.html"
}