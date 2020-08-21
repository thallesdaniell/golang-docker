package main

import (
	"net/http"
	"html/template"
)

type Item struct{
	Name string
}


func hello(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("hello.html"))

	itens := []Item{
		{"cars"},
		{"balls"},
		{"fruits"},
	}

	
	tmpl.Execute(w,itens)
}

func main() {
	 http.HandleFunc("/",hello)
	 http.ListenAndServe(":8080",nil)
}
