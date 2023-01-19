package main

import (
	"fmt"
	gp "groupie"
	gpd "groupie/datas"
	"html/template"
	"net/http"
	"strconv"
	"strings"
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

	sdatas.All = donnermoi.All

	fmt.Println(sdatas)
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/artistes.html"))
	//http.Redirect(w, r, "/", 302)
	tmpl.Execute(w, sdatas)
	return
}

func HandleFilter(w http.ResponseWriter, r *http.Request) {
	buttons := r.FormValue("Member")
	fmt.Println(buttons)
	creation := r.FormValue("creationdate")
	album := r.FormValue("albumdate")
	city := r.FormValue("city")
	fmt.Println(creation)
	fmt.Println(album)
	fmt.Println(city)

	var Donnees gpd.DATAS
	intbutton, _ := strconv.Atoi(buttons)
	intcreation, _ := strconv.Atoi(creation)
	intalbum, _ := strconv.Atoi(album)
	var splitalbum []int
	for i := 0; i < (len(donnermoi.Artist)); i++ {
		cioucou := strings.Split(donnermoi.Artist[i].First_ablbum, "-")[2]
		coucou, _ := strconv.Atoi(cioucou)
		splitalbum = append(splitalbum, coucou)
	}
	// for i := 0; i < (len(donnermoi.Artist)); i++ {
	// 	cioucou := strings.Split(donnermoi.Artist[i].Rlations, "-")[1]
	// 	splitalbum = append(splitalbum, cioucou)
	// }

	for i := 0; i < (len(donnermoi.Artist)); i++ {
		if buttons != "All" {
			if len(donnermoi.Artist[i].Members) == intbutton && donnermoi.Artist[i].Creation_date >= intcreation && int(splitalbum[i]) >= intalbum {
				var Artist gpd.ARTIST
				Artist.Name = donnermoi.Artist[i].Name
				Artist.Image = donnermoi.Artist[i].Image
				Artist.Id = donnermoi.Artist[i].Id
				Donnees.Artist = append(Donnees.Artist, Artist)
			}
		} else {
			if donnermoi.Artist[i].Creation_date >= intcreation && int(splitalbum[i]) >= intalbum {
				var Artist gpd.ARTIST
				Artist.Name = donnermoi.Artist[i].Name
				Artist.Image = donnermoi.Artist[i].Image
				Artist.Id = donnermoi.Artist[i].Id
				Donnees.Artist = append(Donnees.Artist, Artist)
			}
		}

	}

	// for i := 0; i < (len(donnermoi.Artist)); i++ {
	// 	if donnermoi.Artist[i].Creation_date >= intcreation {
	// 		var Artist gpd.ARTIST
	// 		Artist.Name = donnermoi.Artist[i].Name
	// 		Artist.Image = donnermoi.Artist[i].Image
	// 		Artist.Id = donnermoi.Artist[i].Id
	// 		Donnees.Artist = append(Donnees.Artist, Artist)
	// 	}
	// }

	// for i := 0; i < (len(donnermoi.Artist)); i++ {
	// 	if int(splitalbum[i]) >= intalbum {
	// 		var Artist gpd.ARTIST
	// 		Artist.Name = donnermoi.Artist[i].Name
	// 		Artist.Image = donnermoi.Artist[i].Image
	// 		Artist.Id = donnermoi.Artist[i].Id
	// 		Donnees.Artist = append(Donnees.Artist, Artist)

	// 	}
	// }
	// Iid, _ := strconv.Atoi(id)

	// Iid = Iid - 1
	// loc := donnermoi.Location[Iid]
	// art := donnermoi.Artist[Iid]
	// dat := donnermoi.Date[Iid]
	// a := donnermoi.Locs[Iid]
	// donnerartist := gpd.ArtistInfos{}
	// donnerartist.Artist = art
	// donnerartist.Location = loc
	// donnerartist.All = a

	// for i := 0; i < len(dat.Dates); i++ {
	// 	if string(dat.Dates[i][0]) == "*" {
	// 		dat.Dates[i] = dat.Dates[i][1:]
	// 	}
	// }

	// donnerartist.Date = dat

	Donnees.All = donnermoi.All

	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/artistes.html"))
	tmpl.Execute(w, Donnees)
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
