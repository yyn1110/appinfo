package main

import (
	"AppInfoLook/controllers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/appinfo/:lg/:appid", &controllers.SearchController{})
	beego.Router("/appinfo/:lg", &controllers.SearchController{})
	beego.Router("/appinfo/:lg/", &controllers.SearchController{})
	beego.Router("/appinfo/", &controllers.SearchController{})
	beego.Router("/appinfo", &controllers.SearchController{})
	beego.Router("/appstore", &controllers.AppStoreController{})
	beego.Run()
}
