package controllers

import (
	"encoding/json"
	"gameserver/src/account/models"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	userModel := &models.User{}
	res, _ := json.Marshal(userModel.GET())
	c.Ctx.WriteString(string(res))

}
