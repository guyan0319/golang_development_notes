package main

import "fmt"

func multiParam(args ...string) {
	for _, e := range args {
		fmt.Println(e)
	}
}
func main() {
	names := []string{"jerry", "herry"}
	multiParam(names...)
}
