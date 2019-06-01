//go语言打开浏览器
package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	//cmd := exec.Command("cmd", "/C", "start http://www.baidu.com")
	//cmd.Run()
	err := Open("https://blog.csdn.net/guyan0319/article/details/90450958")
	fmt.Println(err)
}

// Open calls the OS default program for uri
func Open(uri string) error {

	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("xdg-open", uri)
		return cmd.Start()
	case "darwin": //苹果系统
		cmd := exec.Command("open", uri)
		return cmd.Start()
	case "windows":
		cmd := exec.Command("cmd", "/C", "start", uri)

		cmd.Start()

		return nil
	}
	return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)

}
