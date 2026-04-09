package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://localhost:8080/" + os.Args[1])

	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)

		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println(string(body))

	}

}
