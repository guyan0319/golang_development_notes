package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	filepath := "D:/gopath/src/example/example/log.txt"
	fi, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)

	chunks := make([]byte, 0)
	fmt.Println(chunks)
	buf := make([]byte, 1024) //一次读取多少个字节
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		fmt.Println(string(buf[:n]))
		break
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	fmt.Println(string(chunks))
}
func readone(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	return string(content)
}
