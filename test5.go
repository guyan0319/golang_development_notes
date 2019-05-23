package main

import (
	"time"
	"fmt"
)

func main() {

	ticker := time.NewTicker(2 * time.Second)
	ch := make(chan bool)
	go func(ticker *time.Ticker) {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				fmt.Println("Ticker running...")
			case stop := <-ch:
				if stop {
					fmt.Println("Ticker Stop")
					return
				}
			}
		}
	}(ticker)
	time.Sleep(10 * time.Second)
	ch <- true
	close(ch)
}
