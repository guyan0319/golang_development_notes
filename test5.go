package main

import (
	"fmt"
	"time"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func main() {

const (
seleniumPath = `C:\Program Files (x86)\chromedriver\chromedriver.exe`
port            = 9515
)

//如果seleniumServer没有启动，就启动一个seleniumServer所需要的参数，可以为空，示例请参见https://github.com/tebeka/selenium/blob/master/example_test.go
opts := []selenium.ServiceOption{}
//opts := []selenium.ServiceOption{
//    selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
//    selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
//}

//selenium.SetDebug(true)
service, err := selenium.NewChromeDriverService(seleniumPath, port, opts...)
if nil != err {
fmt.Println("start a chromedriver service falid", err.Error())
return
}
//注意这里，server关闭之后，chrome窗口也会关闭
defer service.Stop()

//链接本地的浏览器 chrome
caps := selenium.Capabilities{
"browserName": "chrome",
}

//禁止图片加载，加快渲染速度
imagCaps := map[string]interface{}{
"profile.managed_default_content_settings.images": 2,
}
chromeCaps := chrome.Capabilities{
Prefs: imagCaps,
Path:  "",
Args: []string{
//"--headless", // 设置Chrome无头模式，在linux下运行，需要设置这个参数，否则会报错
//"--no-sandbox",
"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36", // 模拟user-agent，防反爬
},
}
//以上是设置浏览器参数
caps.AddChrome(chromeCaps)


// 调起chrome浏览器
w_b1, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
if err != nil {
fmt.Println("connect to the webDriver faild", err.Error())
return
}
//关闭一个webDriver会对应关闭一个chrome窗口
//但是不会导致seleniumServer关闭
defer w_b1.Quit()
err = w_b1.Get("https://zhuanlan.zhihu.com/p/37752206")
if err != nil {
fmt.Println("get page faild", err.Error())
return
}



// 重新调起chrome浏览器
w_b2, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
if err != nil {
fmt.Println("connect to the webDriver faild", err.Error())
return
}
defer w_b2.Close()
//打开一个网页
err = w_b2.Get("https://www.toutiao.com/")
if err != nil {
fmt.Println("get page faild", err.Error())
return
}
//打开一个网页
err = w_b2.Get("https://www.baidu.com/")
if err != nil {
fmt.Println("get page faild", err.Error())
return
}
//w_b就是当前页面的对象，通过该对象可以操作当前页面了
//........
time.Sleep(5* time.Minute)
return
}