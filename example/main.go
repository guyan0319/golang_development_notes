package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	get_act := query["uid"][0]
	fmt.Println(get_act)
	//queryForm, err := url.ParseQuery(r.URL.RawQuery)
	//fmt.Println(queryForm["uid"])
	//if err == nil && len(queryForm["uid"]) > 0 {
	//	fmt.Println(queryForm["uid"][0])
	//}
	//r.ParseForm()
	//uid := r.Form["uid"]
	//fmt.Println(uid)
	//fmt.Fprintf(w, "Hello world!\n") //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayHelloHandler) //	设置访问路由
	log.Fatal(http.ListenAndServe(":8080", nil))
}
