package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	}
	fileName := os.Args[1:2]
	var name string = fileName[0]
	if file, err := os.Open(name); err == nil {
		defer file.Close()
		conn, err := net.Dial("tcp", "127.0.0.1:8081")
		if err != nil {
			fmt.Println("Error, cannot connect")
			return
		}
		scanner := bufio.NewScanner(file)
		for {
			reading := scanner.Scan()
			if reading == false {
				return
			}
			line := scanner.Bytes()
			fmt.Println(line)
			line = append(line, []byte(string('\n'))...)

			fmt.Println(line)
			fmt.Println(string(line))
			conn.Write(line)
		}
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
