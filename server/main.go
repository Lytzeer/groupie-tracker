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
	http.HandleFunc("/search", HandleSearch)
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

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("input")
	intSearch, err := strconv.Atoi(search)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(intSearch)
	fmt.Println(search)
	var sdatas gpd.DATAS
	cpt := 0
	if intSearch == 0 {
		for i := 0; i < len(donnermoi.Artist); i++ {
			if donnermoi.Artist[i].Name == search || donnermoi.Artist[i].First_ablbum == search {
				var Artist gpd.ARTIST
				Artist.Name = donnermoi.Artist[i].Name
				Artist.Image = donnermoi.Artist[i].Image
				Artist.Id = donnermoi.Artist[i].Id
				sdatas.Artist = append(sdatas.Artist, Artist)
				cpt++
			}
			for _, members := range donnermoi.Artist[i].Members {
				if members == search {
					var Artist gpd.ARTIST
					Artist.Name = donnermoi.Artist[i].Name
					Artist.Image = donnermoi.Artist[i].Image
					Artist.Id = donnermoi.Artist[i].Id
					sdatas.Artist = append(sdatas.Artist, Artist)
					cpt++
				}
			}
		}
	}

	if intSearch != 0 {
		for i := 0; i < len(donnermoi.Artist); i++ {
			if donnermoi.Artist[i].Creation_date == intSearch {
				var Artist gpd.ARTIST
				Artist.Name = donnermoi.Artist[i].Name
				Artist.Image = donnermoi.Artist[i].Image
				Artist.Id = donnermoi.Artist[i].Id
				sdatas.Artist = append(sdatas.Artist, Artist)
				cpt++
			}
		}
	}

	fmt.Println(sdatas)
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/artistes.html"))
	//http.Redirect(w, r, "/", 302)
	tmpl.Execute(w, sdatas)
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
