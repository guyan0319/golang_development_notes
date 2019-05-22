package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTimer(time.Second * 2)
	defer t.Stop()
	ch := make(chan bool)
	for {
		select {
		case <-t.C:
			fmt.Println("timer running...")
		case stop := <-ch:
			if stop {
				fmt.Println("timer Stop")
				return
			}
		}
		// 需要重置Reset 使 t 重新开始计时
		t.Reset(time.Second * 2)
	}
	time.Sleep(10 * time.Second)
	ch <- true
	close(ch)
	//ticker := time.NewTicker(2 * time.Second)
	//
	//ch := make(chan bool)
	//go func(ticker *time.Ticker) {
	//	defer ticker.Stop()
	//	for {
	//		select {
	//		case <-ticker.C:
	//			fmt.Println("timer....")
	//		case stop := <-ch:
	//			if stop {
	//				fmt.Println("Ticker2 Stop")
	//				return
	//			}
	//		}
	//	}
	//}(ticker)
	//
	//time.Sleep(10 * time.Second)
	//ch <- true
	//close(ch)
}
