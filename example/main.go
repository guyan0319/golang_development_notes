package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {
	hellowold(10000)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Fprintf(w, "Hello world!\n") //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayHelloHandler) //	设置访问路由
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hellowold(times int) {
	time.Sleep(time.Second)
	var counter int
	for i := 0; i < times; i++ {
		for j := 0; j < times; j++ {
			counter++
		}
	}
}
