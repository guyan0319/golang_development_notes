package main

import (
	"os"
	"io/ioutil"
	"fmt"
)

func main() {
	filepath := "D:/gopath/src/golang_development_notes/example/log.txt"
	//readone(filepath)
	content ,err :=ioutil.ReadFile(filepath)
	if err !=nil {
		panic(err)
	}
	fmt.Println(string(content))
}
func readone ( filepath string )  string{
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	return string(content)
}
