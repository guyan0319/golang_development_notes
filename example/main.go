package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id    int64
	Name  string
	Ctime int64
}

func main() {
	arr := []string{"hello", "apple", "python", "golang", "base", "peach", "pear"}
	lang, err := json.Marshal(arr)
	if err == nil {
		fmt.Println("================array 到 json str==")
		fmt.Println(string(lang))
	}

	jsonStr := "[\"test\",\"testb\"]"
	var dat []string
	err = json.Unmarshal([]byte(jsonStr), &dat)
	if err == nil {
		fmt.Println(dat)
	}
	fmt.Println(err)
}
