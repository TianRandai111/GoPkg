package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm")

type User struct {
	ID int
	UserName string
	Passwd string
}

func init() {
	orm.RegisterDataBase("default","mysql","root:Fnw12345@@tcp(10.39.52.10:3306)/Users")
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default",false,true)

}
