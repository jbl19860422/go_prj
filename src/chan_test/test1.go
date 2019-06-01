package main

import (
	"fmt"
	"time"
)

/*
带3个缓冲区，第一个被i接收，后面发了4个，发到第四个开始阻塞，检查发现没goroutine接收，报告死锁
如果不带缓冲区，相当于是0个缓冲区，发送的时候就阻塞并检查是否有goroutine接收
*/
func main() {
	c := make(chan int, 3)

	go func() {
		i := <-c
		fmt.Println("i=", i)
		_ = i
	}()
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5

	for {
		time.Sleep(1 * time.Second)
	}
}
