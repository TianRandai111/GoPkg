/*
 * @Descripttion:
 * @version:
 * @Author: 步荀仙
 * @Date: 2020-01-21 13:59:18
 * @LastEditors  : 步荀仙
 * @LastEditTime : 2020-01-22 17:07:00
 */
package controllers

import (
	"github.com/TianRandai111/GoPkg/3.Web开发基础/3.3beego开发/2.beego框架/beegoxc/models"
	_ "github.com/TianRandai111/GoPkg/3.Web开发基础/3.3beego开发/2.beego框架/beegoxc/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RegController struct {
	beego.Controller
}

func (this *RegController) ShowReg() {
	this.TplName = "register.html"
}

func (this *RegController) HandleReg() {
	name := this.GetString("userName")
	passwd := this.GetString("password")
	if name == "" || passwd == "" {
		beego.Info("用戶名或密碼不能為空")
		this.TplName = "register.html"
		return
	}

	beego.Info(name, passwd)

	//	1.獲取orm對象
	o := orm.NewOrm()
	//	2.獲取插入對象
	user := models.User{
		UserName: name,
		Passwd:   passwd,
	}
	//	3.插入操作
	_, err := o.Insert(&user)
	if err != nil {
		this.TplName = "register.html"
		beego.Info("数据插入失败", err)
	}

	//	4.插入成功返回登陆失败返回注册
	this.Redirect("/",302)
}

type LoginController struct {
	beego.Controller
}

func (this *LoginController) ShowLogin() {
	this.TplName = "login.html"
}

func (this *LoginController) HandLogin() {
	name := this.GetString("userName")
	passwd := this.GetString("password")
	if name == "" && passwd == "" {
		beego.Info("用戶名或密碼不能為空")
		return
	}

	o := orm.NewOrm()

	user := models.User{
		UserName: name,
	}

	err := o.Read(&user, "UserName")
	if err != nil {
		beego.Info("登陆失败，用户名错误", err)
		this.TplName = "login.html"
		return
	}

	if user.Passwd != passwd {
		beego.Info("登陆失败，密码错误", err)
		this.TplName = "login.html"
		return
	}

	this.Ctx.WriteString("登陆成功")

}
