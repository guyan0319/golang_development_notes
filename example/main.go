package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}

	request, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		fmt.Println(err)
	}

	cookie := &http.Cookie{Name: "JERRY", Value: "dkfsf"}
	request.AddCookie(cookie) //向request中添加cookie

	//设置request的header
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()
	fmt.Println(response.StatusCode)
	if response.StatusCode == 200 {
		r, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(r))
	}
}
