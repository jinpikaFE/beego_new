package routers

import (
	"bee_go_web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/test", &controllers.MainController{}, "get:Test")
}
