package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", "localhost:42069")
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		fmt.Println("Accepted Connection")

		channel := getLinesChannel(conn)

		for line := range channel {
			fmt.Printf("%s\n", line)
		}

		conn.Close()
		fmt.Println("Connection Closed")
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	lineChannel := make(chan string)

	var currLine []byte
	var buf [8]byte

	go func() {
		defer close(lineChannel)

		for {
			count, error := f.Read(buf[:])

			if error == io.EOF {
				break
			}

			for _, element := range buf[:count] {
				if element == '\n' {
					lineChannel <- string(currLine)
					currLine = make([]byte, 0)
					continue
				}
				currLine = append(currLine, element)
			}
		}

		if len(currLine) > 0 {
			lineChannel <- string(currLine)
		}

	}()

	return lineChannel
}
