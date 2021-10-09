package main

import (
	"fmt"
)

var defaultStuffClient = stuffClient{
	retries: 3,
	timeout: 2,
}
type StuffClientOption func(*stuffClient)
func WithRetries(r int) StuffClientOption {
	return func(o *stuffClient) {
		o.retries = r
	}
}
func WithTimeout(t int) StuffClientOption {
	return func(o *stuffClient) {
		o.timeout = t
	}
}
type StuffClient interface {
	DoStuff() error
}
type stuffClient struct {
	conn    Connection
	timeout int
	retries int
}
type Connection struct{}
func NewStuffClient(conn Connection, opts ...StuffClientOption) StuffClient {
	client := defaultStuffClient
	for _, o := range opts {
		o(&client)
	}

	client.conn = conn
	return client
}
func (c stuffClient) DoStuff() error {
	return nil
}
func main() {
	sum := 0
	for i := 0; i <= 100; i++ {
		fmt.Println("fasdfasfsd",i)
		sum += i
	}

	//time.Sleep(time.Millisecond)
}