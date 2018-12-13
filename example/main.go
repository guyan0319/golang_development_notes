package main

import (
	"fmt"
	"log"
	"net/http"
	//_ "net/http/pprof"
)

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Fprintf(w, "Hello world!\n") //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayHelloHandler) //	设置访问路由
	//http.HandleFunc("/debug/pprof/", pprof.Index)
	//http.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	//http.HandleFunc("/debug/pprof/profile", pprof.Profile)
	//http.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	//http.HandleFunc("/debug/pprof/trace", pprof.Trace)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
