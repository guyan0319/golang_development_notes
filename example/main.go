package main

import (
	"database/sql"
	"log"
)

type Person struct {
	//结构也是一种类型
	Name string //定义struct的属性
	Age  int
}

func main() {
	var dbconn *sql.DB
	//fmt.Println(conf[operater])
	dbconn, err := sql.Open(conf["type"], conf[operater])
	//defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}
	//用于设置闲置的连接数。如果 <= 0, 则没有空闲连接会被保留
	dbconn.SetMaxIdleConns(0)
	//用于设置最大打开的连接数,默认值为0表示不限制。
	dbconn.SetMaxOpenConns(30)
	if err := dbconn.Ping(); err != nil {
		log.Fatalln(err)
	}
}
