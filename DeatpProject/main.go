package main

import (
	_ "DeatpProject/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
}

