package main

import (
	"fmt"
	"time"
	"golang.org/x/net/context"
)

func main() {
	//创建一个可取消子context,context.Background():返回一个空的Context，这个空的Context一般用于整个Context树的根节点。
	ctx, cancel := context.WithCancel(context.Background())
	ctxTwo, cancelTwo := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			//使用select调用<-ctx.Done()判断是否要结束
			case <-ctx.Done():
				fmt.Println("goroutineA exit")
				return
			default:
				fmt.Println("goroutineA running.")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)
	go func(ctx context.Context) {
		for {
			select {
			//使用select调用<-ctx.Done()判断是否要结束
			case <-ctx.Done():
				fmt.Println("goroutineB exit")
				return
			default:
				fmt.Println("goroutineB running.")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)
	go func(ctxTwo context.Context) {
		for {
			select {
			//使用select调用<-ctx.Done()判断是否要结束
			case <-ctxTwo.Done():
				fmt.Println("goroutineC exit")
				return
			default:
				fmt.Println("goroutineC running.")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctxTwo)

	time.Sleep(4 * time.Second)
	fmt.Println("main fun exit")
	//取消context
	cancel()
	cancelTwo()
	time.Sleep(5 * time.Second)

}
