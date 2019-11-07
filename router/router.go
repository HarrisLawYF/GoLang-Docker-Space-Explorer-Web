package router

import (
	"SpaceApp/controllers"
	"github.com/astaxie/beego"
)

func Init() {
	beego.Router("/", &controllers.RobotController{})
	beego.Router("/calculate", &controllers.RobotController{})
}
