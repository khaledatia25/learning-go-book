package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func fileLen(name string) (int, error) {
	f, err := os.Open(name)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var count int
	data := make([]byte, 2048)
	for {
		c, err := f.Read(data)
		count += c
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			break
		}
	}
	return count, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	length, err := fileLen(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(length)
}
