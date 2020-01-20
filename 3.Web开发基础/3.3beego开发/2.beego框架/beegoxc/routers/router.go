/*
 * @Descripttion:
 * @version:
 * @Author: 步荀仙
 * @Date: 2020-01-11 14:37:30
 * @LastEditors  : 步荀仙
 * @LastEditTime : 2020-01-14 20:44:58
 */
package routers

import (
	"github.com/TianRandai111/GoPkg/3.Web开发基础/3.3beego开发/2.beego框架/beegoxc/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// beego.Router("/index/?:id", &controllers.IndexController{}, "get:ShowGet;post:Post") //指定路由
	beego.Router("/index/*", &controllers.IndexController{}, "get:ShowGet;post:Post") //指定路由
	beego.Router("/mysql", &controllers.MysqlController{}, "get:ShowMysql;post:Post")
	beego.Router("/orm",&controllers.OrmController{},"get:ShowOrm")
	beego.Router("/insert",&controllers.OrmController{},"get:ShowInsert")
	beego.Router("/update",&controllers.OrmController{},"get:ShowUpdate")

}
