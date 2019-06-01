package main

import (
	"config"
	"fmt"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	var data []byte = make([]byte, 1024, 1024)
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second)) // 100ms
	_, err := conn.Read(data)
	if err != nil {
		fmt.Println("read data failed, err=" + err.Error())
		return
	}

	str := string(data)
	fmt.Println("server recv:" + str)
	sendData := []byte(conn.RemoteAddr().String())
	conn.Write(sendData)
}

func main() {
	fmt.Println("hellow")
	config.LoadConfig()

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("listen to 8080 error")
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("conn accept failed")
			return
		}

		go handleConnection(conn)
	}
}
