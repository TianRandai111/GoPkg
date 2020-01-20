/*
 * @Descripttion:
 * @version:
 * @Author: 步荀仙
 * @Date: 2020-01-14 20:47:54
 * @LastEditors  : 步荀仙
 * @LastEditTime : 2020-01-20 11:17:43
 */
package controllers

import (
	"database/sql"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlController struct {
	beego.Controller
}

// 1.打开数据库

// 2.操作数据库

// 3.关闭数据库
func (this *MysqlController) ShowMysql() {
	// 1.打开数据库
	// 1.1.驱动名称,用户名密码@协议(IP地址:端口)/库名
	conn, err := sql.Open("mysql", "root:Fnw12345@@tcp(10.39.52.10)/test?charset=utf8")
	if err != nil {
		beego.Info("连接错误", err)
		return
	}
	// 关闭数据库
	defer conn.Close()
	//_,err = conn.Exec("create table userInfo(id int,name varchar(11))")
	//if err != nil {
	//	beego.Info("建表错误",err)
	//	return
	//}
	//插入数据
	//_, err = conn.Exec("insert userInfo(id,name)VALUE (?,?)", 1, "BXX",1)
	//if err != nil {
	//	beego.Info("插入数据成失败",err)
	//	this.Data["errmsg"] = "404"
	//	this.TplName = "www.baidu.com"
	//	return
	//}
	// 查询数据
	//rows,err:=conn.Query("select id from userInfo")
	//if err != nil {
	//	beego.Info("查询报错",err)
	//}
	//var id int
	//for rows.Next(){
	//	rows.Scan(&id)
	//	beego.Info(id)
	//}
	this.Ctx.WriteString("插入数据成功")
}
