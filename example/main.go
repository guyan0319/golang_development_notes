package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"time"
)

var db *sql.DB

func main() {
	var err error
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=true&&loc=Local")

	//插入
	insert()
	//更新
	update()
	//查询
	query()
	//删除
	delete()
}
func insert() {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	ctime := time.Now().Unix()
	defer tx.Rollback()
	stmt, err := tx.Prepare("INSERT INTO user (name,ctime) VALUES (?,?)")
	if err != nil {
		return
	}
	rs, err := stmt.Exec("李四", ctime)
	if err != nil {
		return
	}
	err = tx.Commit()
	if err != nil {
		return
	}
	id, err := rs.LastInsertId()
	fmt.Println(id)
	defer stmt.Close() //runs here!
	return
}
func update() {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("UPDATE user SET name=?  WHERE id=?")
	if err != nil {
		return
	}
	_, err = stmt.Exec("zhangsan", 1)
	if err != nil {
		return
	}
	err = tx.Commit()
	if err != nil {
		return
	}
	defer stmt.Close()
	return
}
func query() {
	rows, err := db.Query("SELECT id,name,ctime FROM user  ")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id, ctime int64
		var name string
		rows.Scan(&id, &name, &ctime)
		fmt.Println(id, name, ctime)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}
func delete() {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("DELETE FROM user  WHERE id=?")
	if err != nil {
		return
	}
	_, err = stmt.Exec(1)
	if err != nil {
		return
	}
	err = tx.Commit()
	if err != nil {
		return
	}
	defer stmt.Close()
	return
}
