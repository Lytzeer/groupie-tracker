package main

import (
	"encoding/json"
	"fmt"
	gp "groupie"
	gpd "groupie/datas"
	"html/template"
	"io/ioutil"
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
	members := []string{}
	search := r.FormValue("input")
	intSearch, _ := strconv.Atoi(search)

	var sdatas gpd.DATAS
	if intSearch == 0 {
		for i := 0; i < (len(Alldatas.Artist)); i++ {
			for _, jsp := range Alldatas.Location[i].Locations {
				if (jsp == search || Alldatas.Artist[i].Name == search || Alldatas.Artist[i].First_ablbum == search) && !gp.Isin(Alldatas.Artist[i].Name, members) {
					var Artist gpd.ARTIST
					Artist.Name = Alldatas.Artist[i].Name
					Artist.Image = Alldatas.Artist[i].Image
					Artist.Id = Alldatas.Artist[i].Id
					sdatas.Artist = append(sdatas.Artist, Artist)
					members = append(members, Artist.Name)
				}
			}
			for _, member := range Alldatas.Artist[i].Members {
				if member == search && !gp.Isin(member, members) {
					var Artist gpd.ARTIST
					Artist.Name = Alldatas.Artist[i].Name
					Artist.Image = Alldatas.Artist[i].Image
					Artist.Id = Alldatas.Artist[i].Id
					sdatas.Artist = append(sdatas.Artist, Artist)
					members = append(members, Artist.Name)
					fmt.Println("b")
				}
			}
		}
	}

	if intSearch != 0 {
		for i := 0; i < len(Alldatas.Artist); i++ {
			if Alldatas.Artist[i].Creation_date == intSearch && !gp.Isin(Alldatas.Artist[i].Name, members) {
				var Artist gpd.ARTIST
				Artist.Name = Alldatas.Artist[i].Name
				Artist.Image = Alldatas.Artist[i].Image
				Artist.Id = Alldatas.Artist[i].Id
				sdatas.Artist = append(sdatas.Artist, Artist)
				members = append(members, Artist.Name)
				fmt.Println("c")
			}
		}
	}

	sdatas.All = Alldatas.All
	sdatas.Country = Alldatas.Country

	var tmpl *template.Template
	if len(sdatas.Artist) == 0 {
		tmpl = template.Must(template.ParseFiles("./static/noresult.html"))
	} else {
		tmpl = template.Must(template.ParseFiles("./static/artistes.html"))
	}
	tmpl.Execute(w, sdatas)
	return
}
func Displaydata(i int, Donnees gpd.DATAS) []gpd.ARTIST {
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
	name := []string{"rien"}
	for i := 0; i < (len(Alldatas.Artist)); i++ {
		splitalbumel := strings.Split(Alldatas.Artist[i].First_ablbum, "-")[2]
		splitalbumstr, _ := strconv.Atoi(splitalbumel)
		splitalbum = append(splitalbum, splitalbumstr)
	}

	for i := 0; i < len(Alldatas.Artist); i++ {
		for _, jsp := range Alldatas.Location[i].Locations {
			capi := strings.Split(jsp, "-")[1]
			if buttons != "All" && city != "All" {
				if len(Alldatas.Artist[i].Members) == intbutton && Alldatas.Artist[i].Creation_date >= intcreation && int(splitalbum[i]) >= intalbum && capi == city && !gp.Isin(Alldatas.Artist[i].Name, name) {
					Donnees.Artist = Displaydata(i, Donnees)
					name = append(name, Alldatas.Artist[i].Name)
				}
			} else if buttons == "All" && city != "All" {
				if Alldatas.Artist[i].Creation_date >= intcreation && int(splitalbum[i]) >= intalbum && capi == city && !gp.Isin(Alldatas.Artist[i].Name, name) {
					Donnees.Artist = Displaydata(i, Donnees)
					name = append(name, Alldatas.Artist[i].Name)
				}
			} else if buttons == "All" && city == "All" {
				if Alldatas.Artist[i].Creation_date >= intcreation && int(splitalbum[i]) >= intalbum && !gp.Isin(Alldatas.Artist[i].Name, name) {
					Donnees.Artist = Displaydata(i, Donnees)
					name = append(name, Alldatas.Artist[i].Name)
				}
			} else if buttons != "All" && city == "All" {
				if len(Alldatas.Artist[i].Members) == intbutton && Alldatas.Artist[i].Creation_date >= intcreation && int(splitalbum[i]) >= intalbum && !gp.Isin(Alldatas.Artist[i].Name, name) {
					Donnees.Artist = Displaydata(i, Donnees)
					name = append(name, Alldatas.Artist[i].Name)
				}
			}
		}
	}

	Donnees.All = Alldatas.All
	Donnees.Country = Alldatas.Country

	var tmpl *template.Template
	if len(Donnees.Artist) == 0 {
		tmpl = template.Must(template.ParseFiles("./static/noresult.html"))
	} else {
		tmpl = template.Must(template.ParseFiles("./static/artistes.html"))
	}
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

	var test2 []gpd.FeatureCollection

	for _, loc := range loc.Locations {

		var test gpd.FeatureCollection
		loca := ""
		url_loc := ""
		tiret := false
		for _, letter := range loc {
			if string(letter) != "-" && tiret == false {
				loca += string(letter)
			} else if string(letter) == "-" {
				tiret = true
			}
		}
		url_loc = "https://api.mapbox.com/geocoding/v5/mapbox.places/" + loca + ".json?access_token=pk.eyJ1IjoibWF0c3VlbGwiLCJhIjoiY2xkbjNoMTgzMGZseDN1bHgybjgwbmFnOCJ9.qUR-JuwsRM69PeuHEcVo4A"

		data, _ := http.Get(url_loc)
		responseData, _ := ioutil.ReadAll(data.Body)
		json.Unmarshal(responseData, &test)

		test2 = append(test2, test)
	}

	cartes := []string{}
	carte := "https://api.mapbox.com/styles/v1/mapbox/streets-v12/static/"

	for i, l := range test2 {
		coordonnees1 := strconv.FormatFloat(l.Features[0].Center[0], 'g', 9, 32)
		coordonnees2 := strconv.FormatFloat(l.Features[0].Center[1], 'g', 9, 32)
		if i == len(test2)-1 {
			carte += "pin-l-music+f74e4e(" + coordonnees1 + "," + coordonnees2 + ")" + "/20,0,0/500x500?access_token=pk.eyJ1IjoibWF0c3VlbGwiLCJhIjoiY2xkbjNoMTgzMGZseDN1bHgybjgwbmFnOCJ9.qUR-JuwsRM69PeuHEcVo4A"
		} else {
			carte += "pin-l-music+f74e4e(" + coordonnees1 + "," + coordonnees2 + "),"
		}
		cartes = append(cartes, carte)
	}

	donnerartist.Carte = carte
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/info.html"))
	tmpl.Execute(w, donnerartist)
	return
}
