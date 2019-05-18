package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("goroutineA finish")
		wg.Done()
	}()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("goroutine finish")
		wg.Done()
	}()
	wg.Wait()
	time.Sleep(2 * time.Second)
	fmt.Println("main fun exit")
}
