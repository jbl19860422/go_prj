package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("dial failed")
	}
	data := []byte("hello i'm client")
	n, errConn := conn.Write(data)
	defer conn.Close()
	if errConn != nil {
		fmt.Println("write failed, err=" + errConn.Error())
		return
	}

	fmt.Println("write succeed, count=" + strconv.Itoa(n))
	conn.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second))
	for {
		recvData := make([]byte, 100, 1024)
		_, err := conn.Read(recvData)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read failed, err=" + err.Error())
		}
		fmt.Println("read data:" + string(recvData))
	}
	fmt.Println("end request")
}
