package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, _ := os.Open("messages.txt")

	channel := getLinesChannel(file)

	for line := range channel {
		fmt.Printf("read: %s\n", line[:])
	}

}

func getLinesChannel(f io.ReadCloser) <-chan string {
	lineChannel := make(chan string)

	var currLine []byte
	var buf [8]byte

	go func() {
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

		f.Close()
		close(lineChannel)
	}()

	return lineChannel
}
