package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"bufio"
	"io"
)

func main() {
	filepath := "D:/gopath/src/golang_development_notes/example/log.txt"
	fi, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)

	chunks := make([]byte, 1024, 10240)

	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}

		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n])
	}
	fmt.Println(string(chunks))



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
