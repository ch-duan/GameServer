package main

import (
	_ "gameserver/src/account/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
