package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"urlshort"
)

type SourceFileType int

const (
	_ SourceFileType = iota
	YAML
	JSON
)

func parseArgs() (string, SourceFileType) {
	var filename string
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "--file=") {
			parts := strings.Split(arg, "=")
			filename = parts[1]
		}
	}
	var sourceFileType SourceFileType
	switch {
	case strings.HasSuffix(filename, ".json"):
		sourceFileType = JSON
	case strings.HasSuffix(filename, ".yaml"):
		sourceFileType = YAML
	default:
		panic("Only json and yaml files are supported")
	}
	return filename, sourceFileType
}

func main() {
	filename, srcType := parseArgs()
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	var fileHandler http.HandlerFunc
	switch srcType {
	case YAML:
		fileHandler, err = urlshort.YAMLHandler([]byte(data), mapHandler)
	case JSON:
		fileHandler, err = urlshort.JSONHandler([]byte(data), mapHandler)
	}
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", fileHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
