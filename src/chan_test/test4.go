package main

import (
	"fmt"
	"time"
)

/*
测试chan关闭，所有监听的goroutine可以收到事件

最终打印：
recv1 is close
recv3 is close
recv2 is close
*/
func recv1(c <-chan int) {
	_, ok := <-c
	if !ok {
		fmt.Println("recv1 is close")
	}
}

func recv2(c <-chan int) {
	_, ok := <-c
	if !ok {
		fmt.Println("recv2 is close")
	}
}

func recv3(c <-chan int) {
	_, ok := <-c
	if !ok {
		fmt.Println("recv3 is close")
	}
}

func main() {
	c := make(chan int, 3)
	go recv1(c)
	go recv2(c)
	go recv3(c)

	close(c)
	for {
		time.Sleep(time.Second)
	}
}
