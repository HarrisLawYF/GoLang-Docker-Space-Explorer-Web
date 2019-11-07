package conf

import (
	router "SpaceApp/router"
	"github.com/astaxie/beego";
)

func Init_restfulAPI_service(method string){
	beego.BConfig.RunMode = "dev" 
	beego.BConfig.WebConfig.AutoRender = true
	beego.BConfig.WebConfig.EnableDocs = true
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.SetStaticPath("/images","images")
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.Listen.HTTPPort = 8080
	beego.BConfig.Listen.HTTPSPort = 8081
	if method == "HTTP" {
		beego.BConfig.Listen.EnableHTTP = true
		beego.BConfig.Listen.EnableHTTPS = false
	} else if method == "HTTPS" {
		beego.BConfig.Listen.EnableHTTP = false
		beego.BConfig.Listen.EnableHTTPS = true
	}
	router.Init()
	beego.Run()
}

