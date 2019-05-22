package main

import (
	"context"
	"fmt"
)

func Route(ctx context.Context) {
	ret, ok := ctx.Value("id").(int)
	if !ok {
		ret = 1
	}
	fmt.Printf("id:%d\n", ret)
	s, _ := ctx.Value("name").(string)
	fmt.Printf("name:%s\n", s)
}

func main() {
	ctx := context.WithValue(context.Background(), "id", 123)
	ctx = context.WithValue(ctx, "name", "jerry")
	Route(ctx)
}
