package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
)

const BUFFER_SIZE int64 = 0

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":8081")

	conn, _ := ln.Accept()
	for {
		message, _ := bufio.NewReader(conn).ReadBytes('\n')
		fmt.Println(message)
		fmt.Println(string(message))
		err := ioutil.WriteFile("/tmp/data", message, 0644)
		if err != nil {
			fmt.Println("Error")
		}
		newmessage := "Received bytes"
		conn.Write([]byte(newmessage + "\n"))
		return
	}
}
