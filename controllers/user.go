package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Ctx.WriteString("hello world")
}

type my struct {
	Test string `json:"test"`
}

func (c *UserController) Post() {
	test := c.GetString("test")
	mystruct := my{Test: test + "2"}
	c.Data["json"] = &mystruct
	c.ServeJSON()
}
