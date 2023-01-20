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

var Alldatas gpd.DATAS
var Da gpd.DATE
var Ar []gpd.ARTIST
var Gl gpd.GetLocation
var Re gpd.RELATION

func main() {
	Da, Ar, Gl, Re = gp.GetDatas()
	Alldatas = gp.SetData(Da, Ar, Gl, Re)
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
	tmpl.Execute(w, Alldatas)
	return
}

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("input")
	intSearch, _ := strconv.Atoi(search)

	
	var sdatas gpd.DATAS
	cpt := 0
	if intSearch == 0 {
		for i := 0; i < len(Alldatas.Artist); i++ {
			if Alldatas.Artist[i].Name == search || Alldatas.Artist[i].First_ablbum == search {
				var Artist gpd.ARTIST
				Artist.Name = Alldatas.Artist[i].Name
				Artist.Image = Alldatas.Artist[i].Image
				Artist.Id = Alldatas.Artist[i].Id
				sdatas.Artist = append(sdatas.Artist, Artist)
				cpt++
			}
			for _, members := range Alldatas.Artist[i].Members {
				if members == search {
					var Artist gpd.ARTIST
					Artist.Name = Alldatas.Artist[i].Name
					Artist.Image = Alldatas.Artist[i].Image
					Artist.Id = Alldatas.Artist[i].Id
					sdatas.Artist = append(sdatas.Artist, Artist)
					cpt++
				}
			}
		}
	}

	if intSearch != 0 {
		for i := 0; i < len(Alldatas.Artist); i++ {
			if Alldatas.Artist[i].Creation_date == intSearch {
				var Artist gpd.ARTIST
				Artist.Name = Alldatas.Artist[i].Name
				Artist.Image = Alldatas.Artist[i].Image
				Artist.Id = Alldatas.Artist[i].Id
				sdatas.Artist = append(sdatas.Artist, Artist)
				cpt++
			}
		}
	}

	sdatas.All = Alldatas.All
	sdatas.Country = Alldatas.Country

	
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/artistes.html"))
	//http.Redirect(w, r, "/", 302)
	tmpl.Execute(w, sdatas)
	return
}
func Displaydata(i int, Donnees gpd.DATAS) []gpd.ARTIST{
	var Artist gpd.ARTIST
	Artist.Name = Alldatas.Artist[i].Name
	Artist.Image = Alldatas.Artist[i].Image
	Artist.Id = Alldatas.Artist[i].Id
	Donnees.Artist = append(Donnees.Artist, Artist)
	return (Donnees.Artist)
}

func HandleFilter(w http.ResponseWriter, r *http.Request) {
	buttons := r.FormValue("Member")
	creation := r.FormValue("creationdate")
	album := r.FormValue("albumdate")
	city := r.FormValue("city")

	var Donnees gpd.DATAS
	intbutton, _ := strconv.Atoi(buttons)
	intcreation, _ := strconv.Atoi(creation)
	intalbum, _ := strconv.Atoi(album)
	var splitalbum []int
	for i := 0; i < (len(Alldatas.Artist)); i++ {
		splitalbumel := strings.Split(Alldatas.Artist[i].First_ablbum, "-")[2]
		splitalbumstr, _ := strconv.Atoi(splitalbumel)
		splitalbum = append(splitalbum, splitalbumstr)
	}

	for i := 0; i < (len(Alldatas.Artist)); i++ {
		j := 0
		capi := strings.Split(Alldatas.Location[i].Locations[j], "-")[1]
		if buttons != "All" && city != "All" {
			if len(Alldatas.Artist[i].Members) == intbutton && Alldatas.Artist[i].Creation_date >= intcreation && int(splitalbum[i]) >= intalbum && capi == city {
				Donnees.Artist = Displaydata(i , Donnees)
			}
			j++
		} else if buttons == "All" && city != "All" {
			if Alldatas.Artist[i].Creation_date >= intcreation && int(splitalbum[i]) >= intalbum && capi == city {
				Donnees.Artist = Displaydata(i , Donnees)
			}
			j++
		} else if buttons == "All" && city == "All" {
			if Alldatas.Artist[i].Creation_date >= intcreation && int(splitalbum[i]) >= intalbum {
				Donnees.Artist = Displaydata(i , Donnees)
			}
		} else if buttons != "All" && city == "All" {
			if len(Alldatas.Artist[i].Members) == intbutton && Alldatas.Artist[i].Creation_date >= intcreation && int(splitalbum[i]) >= intalbum {
				Donnees.Artist = Displaydata(i , Donnees)
			}
		}

	}

	Donnees.All = Alldatas.All
	Donnees.Country = Alldatas.Country

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

	loc := Alldatas.Location[Iid]
	art := Alldatas.Artist[Iid]
	dat := Alldatas.Date[Iid]
	a := Alldatas.Locs[Iid]
	donnerartist := gpd.ArtistInfos{}
	donnerartist.Artist = art
	donnerartist.Location = loc
	donnerartist.All = a

	

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
