 /*
 * @Descripttion:
 * @version:
 * @Author: 步荀仙
 * @Date: 2020-01-11 14:37:30
 * @LastEditors  : 步荀仙
 * @LastEditTime : 2020-01-14 20:29:54
 */
package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["test"] = "abccc"
	c.TplName = "test.html"
}

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Post() {
	this.Data["test"] = "这是个啥"
	this.TplName = "test.html"
}

func (this *IndexController) ShowGet() {
	//不确定的*.*用path和ext接受
	// path := this.GetString(":path")
	// ext := this.GetString(":ext")
	splat := this.GetString(":splat")
	// beego.Info(id)
	// beego.Info("Path=", path, "ext=", ext)
	beego.Info("splat=", splat)
	this.Data["test"] = "查看Get方法"
	this.TplName = "test.html"
}
