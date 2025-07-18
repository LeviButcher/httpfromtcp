package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	address, _ := net.ResolveUDPAddr("udp", "127.0.0.1:42069")
	conn, _ := net.DialUDP("udp", nil, address)
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(">")
		line, err := reader.ReadString('\n')

		if err == nil {
			fmt.Println(err)
		}

		_, err = conn.Write([]byte(line))

		if err == nil {
			fmt.Println(err)
		}
	}
}
