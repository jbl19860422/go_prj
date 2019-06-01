package main

import (
	"fmt"
	"time"
)

//测试select的用法
/*
打印内容
wait for done
wait for done
wait for done
wait for done
work done
work has been done, we exit
*/
func doWork(c chan bool) {
	time.Sleep(4 * time.Second)
	fmt.Println("work done")
	c <- true
}

func main() {
	c := make(chan bool)
	go doWork(c)
DONE:
	for {
		select {
		case <-c:
			fmt.Println("work has been done, we exit")
			break DONE //这里如果直接写break，那么break作用不到外面的for上，所以需要在这里写明break DONE才能跳出for
		default:
			fmt.Println("wait for done")
			time.Sleep(time.Second)
		}
		// break DONE //break是跳到for的后面执行，所以是会退出for的
	}
}
