package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	//DB 传递变量
	DB *sql.DB
)

func init() {
	database, err := sql.Open("mysql", "root:Fnw12345@@tcp(10.39.52.10:3306)/test")
	if err != nil {
		log.Fatalln("链接失败", err)
	}

	err = database.Ping()
	if err != nil {
		log.Fatalln("数据库链接失败", err)
		return
	}
	log.Println("连接成功")
	DB = database
}

func main() {
	defer DB.Close()

	//基础操作
	// sql := "insert into stu values (1,'素还真')"
	// data1, err := DB.Exec(sql)
	// if err != nil {
	// 	log.Fatalln("data1数据写入失败", err)
	// }
	// n, _ := data1.RowsAffected()
	// fmt.Println("data1受影响的记录", n)

	// data2, err := DB.Exec("insert into stu values (?,?)", 2, "一页书")
	// if err != nil {
	// 	log.Fatalln("data2数据写入失败", err)
	// }
	// n, _ = data2.RowsAffected()
	// fmt.Println("data2受影响的记录", n)

	//预处理
	// sql := [2][2]string{{"2", "南宫恨"}, {"3", "史艳文"}}
	// stmt, err := DB.Prepare("insert into stu value (?,?)")
	// if err != nil {
	// 	log.Fatalln("操作数据失败", err)
	// }
	// for _, v := range sql {
	// 	stmt.Exec(v[0], v[1])
	// }

	//单行查询
	var (
		id   int
		name string
	)
	rows := DB.QueryRow("select * from stu where id=1")
	rows.Scan(&id, &name) //将rows中的数据存放到id,name中
	fmt.Println(id, "--", name)

	//多行查询
	allrows, _ := DB.Query("select * from stu")
	for allrows.Next() {
		allrows.Scan(&id, &name)
		fmt.Println(id, "--", name)
	}

}
