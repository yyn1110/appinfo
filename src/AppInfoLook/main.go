package main

import (
	"AppInfoLook/controllers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/appinfo/:appid", &controllers.SearchController{})
	beego.Run()
}
