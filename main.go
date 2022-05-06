package main

import (
	_ "bee_go_web/routers"
	"github.com/astaxie/beego"
)

func main() {
	// 对应api类型将默认渲染改为false，即不渲染模板
	beego.BConfig.WebConfig.AutoRender = false
	beego.Run()
}

