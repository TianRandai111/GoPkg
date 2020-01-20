package controllers

import (
	_ "github.com/TianRandai111/GoPkg/3.Web开发基础/3.3beego开发/2.beego框架/beegoxc/routers"
	_ "github.com/TianRandai111/GoPkg/3.Web开发基础/3.3beego开发/2.beego框架/beegoxc/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type OrmController struct{
	beego.Controller
}

func (this *OrmController)ShowOrm(){
	// 1.连接数据库
	orm.RegisterDataBase("default","mysql","root:Fnw12345@@tcp(10.39.52.10:3306)/test?charset=utf8")
	// 2.注册表
	orm.RegisterModel(new(models.User))
	// 3.生成表
	orm.RunSyncdb("default",false,true)
	this.Ctx.WriteString("创建表成功")
}

func (this *OrmController)ShowInsert(){

	//插入数据
	//	1.获取链接对
	//  	象
	//o:=orm.NewOrm()
	//
	//// 2. 获取插入对象
	//user:=User{
	//	Name:"BXX",
	//}
	//// 3.插入数据
	//_,err:=o.Insert(&user)
	//if err != nil {
	//	beego.Info("插入失败",err)
	//	return
	//}
	//this.Ctx.WriteString("插入成功")
	user:=models.User{
		Id :1,
	}
	//	1.获取链接对象
	o := orm.NewOrm()
	//  2.获取查询对象

	//  3.执行查询操作
	err := o.Read(&user)
	if err != nil {
		beego.Info("查询失败",err)
	}
	beego.Info(user)
	this.Ctx.WriteString("查询成功")
}
func (this *OrmController)ShowUpdate(){
	//1.获取orm对象
	o := orm.NewOrm()
	//2.获取更新对象
	user:= models.User{Id:1}
	//3.查询对象
	err := o.Read(&user)
	if  err != nil {
		beego.Info("查询数据失败",err)
		return
	}
	//4.查询到的对象复制
	user.Name="KYZ"
	//5.更新操作
	_,err = o.Update(&user)
	if err != nil {
		beego.Info("更新数据错误",err )
		return
	}
	this.Ctx.WriteString("更新成功")
}
