package main

/*
这个例子，演示了只读和只写通道，
并且奇怪的是，因为recv只接收一次，但是后面send再发没有接收的goroutine，居然不挂。
*/
import (
	"fmt"
	"time"
)

func send(c chan<- int) {
	c <- 1
	fmt.Println("send 1")
	c <- 2
	fmt.Println("send 2")
	c <- 3
	fmt.Println("send 3")
}

func recv(c <-chan int) {
	i := <-c
	fmt.Println("recv ", i)
	_ = i
	//range 接收chan的用法
	// for i := range c {
	// 	fmt.Println("recv ", i)
	// }
}

func main() {
	c := make(chan int)
	// c <- 1  //这样没人接收，则会挂，在main中发送，没人接收，才会导致程序挂，否则只是阻塞那个发送协程而已
	go send(c)
	go recv(c)

	for {
		time.Sleep(1 * time.Second)
	}
}
