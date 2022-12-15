package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port 8080")
	http.HandleFunc("/", HandleIndex)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
	return
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/index.html"))
	tmpl.Execute(w, nil)
	return
}
