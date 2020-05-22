package routers

import (
	"github.com/astaxie/beego"
	"zjko.vip/game/src/account/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/user", &controllers.UserController{})
}
