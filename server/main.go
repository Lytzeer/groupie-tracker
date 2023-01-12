package main

import (
	"fmt"
	gp "groupie"
	"html/template"
	"net/http"
)

var donnermoi gp.DATAS
var Da gp.DATE
var Ar []gp.ARTIST
var Gl gp.GetLocation
var Re gp.RELATION

func main() {
	fmt.Println("Starting server on port 8080")
	http.HandleFunc("/", HandleIndex)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
	return
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	Da, Ar, Gl, Re = gp.GetDatas()
	donnermoi := gp.SetData(Da, Ar, Gl, Re, donnermoi)
	// gp.SetData(Da, Ar, Gl, Re, donnermoi)
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/artistes.html"))
	tmpl.Execute(w, donnermoi)
	fmt.Println(donnermoi)
	return
}
