package main

import (
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(7 * time.Second)
	w.Write([]byte("message"))
}

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
