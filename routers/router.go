package routers

import (
	"github.com/coolyrat/bcs-feeds/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
