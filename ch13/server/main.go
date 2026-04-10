package main

import (
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
)

const url = "https://dummyjson.com/"

type todo struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Title     string `json:"todo"`
	Completed bool   `json:"completed"`
}

var form = `
<h1>Todo #{{.ID}}</h1>
<div>{{printf "User %d" .UserID}}</div>
<div>{{printf "%s (completed: %t)" .Title .Completed}}</div>
`

func getTodo(path string) (todo, error) {
	resp, err := http.Get(url + path)

	if err != nil {
		return todo{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		var item todo
		if err = json.NewDecoder(resp.Body).Decode(&item); err != nil {
			return item, err
		}
		return item, nil
	} else {
		return todo{}, errors.New("error")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, err := getTodo(r.URL.Path[1:])
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	tmpl := template.New("khaled")
	tmpl.Parse(form)
	tmpl.Execute(w, t)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
