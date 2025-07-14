package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, _ := os.Open("messages.txt")

	var currLine []byte
	var buf [8]byte

	for {
		count, error := file.Read(buf[:])

		if error == io.EOF {
			break
		}

		for _, element := range buf[:count] {
			if element == '\n' {
				fmt.Printf("read: %s\n", currLine[:])
				currLine = make([]byte, 0)

				continue
			}
			currLine = append(currLine, element)
		}
	}

	if len(currLine) > 0 {
		fmt.Printf("read: %s\n", currLine[:])
	}

	file.Close()
}
