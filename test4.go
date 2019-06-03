package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://www.baidu.com", http.StatusFound) //重定向
	content := []byte("hello world")
	err := ioutil.WriteFile("test.txt", content, 0644)
	if err != nil {
		panic(err)
	}
}
func main() {
	http.HandleFunc("/", sayHelloHandler) //   设置访问路由
	log.Fatal(http.ListenAndServe(":8080", nil))
}
