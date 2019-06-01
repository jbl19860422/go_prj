package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	/*******测试lookuphost*********/
	addrs, err := net.LookupHost("www.baidu.com")
	if err != nil {
		fmt.Println("LookupHost of www.baidu.com failed")
		return
	}
	fmt.Println("addrs of baidu.com is:")
	fmt.Println(addrs)
	//测试ResolveIp
	ipAddr, err := net.ResolveIPAddr("ip", "www.baidu.com")
	if err != nil {
		fmt.Println("ResolveIPAddr failed")
		return
	}
	fmt.Fprintf(os.Stdout, "%s IP:%s, Zone:%s\n", ipAddr.IP.String(), ipAddr.Zone)
	//查看应用使用的端口
	port, err := net.LookupPort("tcp", "ssh")
	if err != nil {
		fmt.Fprintf(os.Stderr, "LookupPort failed\n")
		return
	}
	fmt.Fprintf(os.Stdout, "ssh's port=%d\n", port)

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
		// reader := bufio.NewReader(conn)
		fmt.Println("new ReaderSize")
		reader := bufio.NewReaderSize(conn, 30)
		scanner := bufio.NewScanner()
		// _, err := reader.Read(recvData)

		// _, err := conn.Read(recvData)
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
