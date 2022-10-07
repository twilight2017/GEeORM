package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" //注册sqlite3驱动
	"log"
)

func main() {
	db, _ := sql.Open("sqlite3", "gee.db") //db是一个sql.DB的指针
	defer func() {
		_ = db.Close()
	}()
	//Exec()用于执行SQL语句，如果是查询语句，不会返回相关的记录
	//查询语句通常使用Query()和QueryRow()
	_, _ = db.Exec("DROP TABLE IF EXIST User;")
	_, _ = db.Exec("CREATE TABLE User(Name text);")
	result, err := db.Exec("Insert INTO User(`Name`) values (?), (?)", "TOM", "Sam")
	if err == nil {
		affected, _ := result.RowsAffected()
		log.Println(affected)
	}
	row := db.QueryRow("SELECT Name FROM User LIMIT 1")
	var name string
	//row.Scan()可以获取到对应的列值
	if err := row.Scan(&name); err == nil {
		log.Println(name)
	}
}
