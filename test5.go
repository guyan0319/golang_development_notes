package main

import (
	"fmt"
	"time"
	"golang.org/x/net/context"
)

func main() {
	//创建一个可取消子context
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine exit")
				return
			default:
				fmt.Println("goroutine running.")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("main fun exit")
	cancel()
	time.Sleep(5 * time.Second)

}
