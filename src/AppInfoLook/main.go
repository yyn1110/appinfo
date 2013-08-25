package main

import (
	"AppInfoLook/controllers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/appinfo/:appid", &controllers.SearchController{})
	beego.Router("/appinfo/", &controllers.SearchController{})
	beego.Router("/appinfo", &controllers.SearchController{})
	beego.Router("/appstore", &controllers.AppStoreController{})
	beego.Router("/roomlist", &controllers.ListController{})
	beego.Run()
}
