package main

import (
	"fmt"
	gp "groupie"
	gpd "groupie/datas"
	"html/template"
	"net/http"
	"strconv"
)

var donnermoi gpd.DATAS
var Da gpd.DATE
var Ar []gpd.ARTIST
var Gl gpd.GetLocation
var Re gpd.RELATION

func main() {
	Da, Ar, Gl, Re = gp.GetDatas()
	donnermoi = gp.SetData(Da, Ar, Gl, Re)
	fmt.Println("Starting server on port 8080")
	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/filter", HandleFilter)
	http.HandleFunc("/infos", HandleInfos)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
	return
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/artistes.html"))
	tmpl.Execute(w, donnermoi)
	return
}

func HandleFilter(w http.ResponseWriter, r *http.Request) {
	buttons := r.FormValue("Member")
	creation := r.FormValue("creationdate")
	album := r.FormValue("albumdate")
	city := r.FormValue("city")
	fmt.Println(buttons)
	fmt.Println(creation)
	fmt.Println(album)
	fmt.Println(city)

	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/artistes.html"))
	tmpl.Execute(w, donnermoi)
	http.Redirect(w, r, "/", 302)
	return
}

func HandleInfos(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	Iid, _ := strconv.Atoi(id)

	Iid = Iid - 1

	loc := donnermoi.Location[Iid]
	art := donnermoi.Artist[Iid]
	dat := donnermoi.Date[Iid]
	a := donnermoi.Locs[Iid]
	donnerartist := gpd.ArtistInfos{}
	donnerartist.Artist = art
	donnerartist.Location = loc
	donnerartist.All = a

	fmt.Println(a)

	for i := 0; i < len(dat.Dates); i++ {
		if string(dat.Dates[i][0]) == "*" {
			dat.Dates[i] = dat.Dates[i][1:]
		}
	}

	donnerartist.Date = dat

	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/info.html"))
	tmpl.Execute(w, donnerartist)
	return
}
