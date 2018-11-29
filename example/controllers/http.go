package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Fprintf(w, "Hello world!\n") //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayHelloHandler) //	设置访问路由
	log.Fatal(http.ListenAndServe(":8080", nil))
}
