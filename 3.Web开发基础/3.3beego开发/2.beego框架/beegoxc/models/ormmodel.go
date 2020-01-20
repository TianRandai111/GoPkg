package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

// 1.存放结构体

// 2.表设计

type User struct {
	Id int
	Name string
}

func init() {
	orm.RegisterDataBase("default","mysql","root:Fnw12345@@tcp(10.39.52.10:3306)/test?charset=utf8")
	// 2.注册表
	orm.RegisterModel(new(User))
	// 3.生成表
	orm.RunSyncdb("default",false,true)
}