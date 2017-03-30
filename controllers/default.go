package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.baidu1.com"
	c.Data["Email"] = "luis.chan@qq.com"
	c.TplName = "index.tpl"
}
