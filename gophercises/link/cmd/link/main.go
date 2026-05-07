package main

import (
	"fmt"
	"link"
	"log"
	"os"
)

func main() {
	filename := "testdata/index.html"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	links, err := link.Parse(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(links)
}
