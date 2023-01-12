package main

import (
	"fmt"
	gp "groupie"
	"html/template"
	"net/http"
	"strconv"
)

var donnermoi gp.DATAS
var Da gp.DATE
var Ar []gp.ARTIST
var Gl gp.GetLocation
var Re gp.RELATION

func main() {
	Da, Ar, Gl, Re = gp.GetDatas()
	donnermoi = gp.SetData(Da, Ar, Gl, Re, donnermoi)
	fmt.Println("Starting server on port 8080")
	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/infos", HandleInfos)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
	return
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {

	//gp.SetData(Da, Ar, Gl, Re, donnermoi)

	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/artistes.html"))
	tmpl.Execute(w, donnermoi)
	fmt.Println(donnermoi)
	return
}
func HandleInfos(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	Iid, _ := strconv.Atoi(id)

	Iid = Iid - 1

	loc := donnermoi.Location[Iid]
	art := donnermoi.Artist[Iid]
	rel := donnermoi.Relation[Iid]
	dat := donnermoi.Date[Iid]
	donnerartist := gp.ArtistInfos{}
	donnerartist.Artist = art
	donnerartist.Location = loc
	donnerartist.Date = dat
	donnerartist.Relation = rel

	fmt.Println(loc)
	fmt.Println(art)
	fmt.Println(dat)
	fmt.Println(rel)
	fmt.Println(donnerartist)
	//gp.SetData(Da, Ar, Gl, Re, donnermoi)

	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/info.html"))
	tmpl.Execute(w, donnerartist)
	// fmt.Println(donnermoi)
	return
}
