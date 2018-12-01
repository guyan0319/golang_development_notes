package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {
	SetCookie(w, r)
	cookie, err := r.Cookie("name")
	if err == nil {
		fmt.Println(cookie.Value)
		fmt.Println(cookie.Domain)
		fmt.Println(cookie.Expires)
	}
	//fmt.Fprintf(w, "Hello world!\n") //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayHelloHandler) //	设置访问路由
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func SetCookie(w http.ResponseWriter, r *http.Request) {
	ck := &http.Cookie{
		Name:   "name",
		Value:  "hello",
		Path:   "/",
		Domain: "localhost",
		MaxAge: 120,
	}
	w.Header().Set("set-cookie", ck.String()) //设置cookie  value 值可以包含空格
	//http.SetCookie(w, ck)  //设置cookie  value 值可以不能包含空格
}
