package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func HandleWelcome() func(c echo.Context) error {
	return func(c echo.Context) error {
		res := c.Response()
		w := res.Writer
		r := c.Request()
		http.Redirect(w, r, "http://www.baidu.com", http.StatusFound) //跳转到百度
		return nil
	}
}

func main() {
	e := echo.New()
	e.GET("/hello", HandleWelcome())
	e.Logger.Fatal(e.Start(":1323"))
}
