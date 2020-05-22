package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"hello/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	userModel := &models.User{}
	res, _ := json.Marshal(userModel.GET())
	c.Ctx.WriteString(string(res))

}
