package main

import (
	"log"
	"os"
)

func main() {
	// 按照所需读写权限创建文件
	f, err := os.OpenFile("filename.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// 完成后延迟关闭，而不是习惯!
	//defer f.Close()
	//设置日志输出到 f
	log.SetOutput(f)
	//测试用例
	log.Println("check to make sure it works")
}
