package main

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
)

/*
package sql
import "database/sql"
Drivers that do not support context cancelation will not return until after the query is completed.
Package sql provides a generic interface around SQL (or SQL-like) databases.
The sql package must be used in conjunction with a database driver. See https://golang.org/s/sqldrivers for a list of drivers.
*/
const (
    selectClass = "SELECT * FROM class "
)

func QueryFromDB(db *sql.DB, cmd string) {
	// 数据库查询
	rows, err := db.Query(cmd)
	if err != nil {
		fmt.Println("error:", err)
	}
	for rows.Next() {
		var clsNO string
		var clsName string
		var Director string
		var Specially string
		err = rows.Scan(&clsNO, &clsName, &Director, &Specially)
		fmt.Printf("%s\t%s\t%s\t%s\n", clsNO, clsName, Director, Specially)
	}
}

func checkErr(isOpen bool) {
	defer func() {
		if r := recover(); r != nil{
			fmt.Printf("捕获的异常：%v\n", r)
		}
	}()
    if !isOpen{
    	panic("数据库打开失败")
	}else{
		// 对数据库执行的操作
	}
}

func main() {
	// Go MySQL Driver实现了sql的interface，只需要导入驱动就可实现所有的sql API
	//（ You only need to import the driver and can use the full database/sql API then.）
	// form https://github.com/go-sql-driver/mysql#dsn-data-source-name
	// 连接数据库，sql.Open(driverName, daraSourceName)
	// driverName: "mysql"
	// dataSourceName(DSN):username:password@protocol(address)/dbname?param=value
	// username: 所有者名字
	// password：数据库密码
	// @protocol: 协议
	// address: 地址
	// dbname: 数据库名字
	// param: [example]charset
	// value: [example]<name>
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/new_schema?charset=utf8")
	// 错误处理
	var isOpen bool
	if err != nil {
		isOpen = false
	}else{
		isOpen = true
	}
	checkErr(isOpen)
	// 对数据执行的操作
	QueryFromDB(db, selectClass)

	defer db.Close()
    /*
	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
    */
}