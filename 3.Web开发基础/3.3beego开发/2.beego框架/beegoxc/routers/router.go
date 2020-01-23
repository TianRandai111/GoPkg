/*
 * @Descripttion:
 * @version:
 * @Author: 步荀仙
 * @Date: 2020-01-11 14:37:30
 * @LastEditors  : 步荀仙
 * @LastEditTime : 2020-01-21 10:39:33
 */
package routers

import (
	"github.com/TianRandai111/GoPkg/3.Web开发基础/3.3beego开发/2.beego框架/beegoxc/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/", &controllers.LoginController{},"get:ShowLogin;post:HandLogin")
	beego.Router("/register",&controllers.RegController{},"get:ShowReg;post:HandleReg")
}
