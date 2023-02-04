package main

import (
	"fmt"
	gp "groupie/api"
	gpd "groupie/datas"
	gpf "groupie/func"
	"html/template"
	"net/http"
	"strconv"
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

	//Récupération et conversion de la valeur entrée dans la barre de recherche
	search := r.FormValue("input")
	intSearch, _ := strconv.Atoi(search)

	//Appel de la fonction pour ne garder que les données souhaitées
	DataSearch := gpf.SearchDatas(search, intSearch, Alldatas)

	var tmpl *template.Template
	if len(DataSearch.Artist) == 0 {
		tmpl = template.Must(template.ParseFiles("./static/noresult.html"))
	} else {
		tmpl = template.Must(template.ParseFiles("./static/artistes.html"))
	}
	tmpl.Execute(w, DataSearch)
	return
}

func HandleFilter(w http.ResponseWriter, r *http.Request) {

	//On réupère et on converti en int les valeurs voulues
	buttonAll := r.FormValue("MemberAll")
	city := r.FormValue("city")
	intbutton1, _ := strconv.Atoi(r.FormValue("Member1"))
	intbutton2, _ := strconv.Atoi(r.FormValue("Member2"))
	intbutton3, _ := strconv.Atoi(r.FormValue("Member3"))
	intbutton4, _ := strconv.Atoi(r.FormValue("Member4"))
	intbutton5, _ := strconv.Atoi(r.FormValue("Member5"))
	intbutton6, _ := strconv.Atoi(r.FormValue("Member6"))
	intbutton7, _ := strconv.Atoi(r.FormValue("Member7"))
	intbutton8, _ := strconv.Atoi(r.FormValue("Member8"))
	intcreation, _ := strconv.Atoi(r.FormValue("creationdate"))
	intalbum, _ := strconv.Atoi(r.FormValue("albumdate"))

	//Tableau pour ajouter toutes les valeurs des bouttons plus facile à traiter
	var tabbutton []int
	tabbutton = append(tabbutton, intbutton1, intbutton2, intbutton3, intbutton4, intbutton5, intbutton6, intbutton7, intbutton8)

	//On appel<e la fonction pour rechercher les données que l'on souhaite
	Donnees := gpf.FilterDatas(Alldatas, buttonAll, tabbutton, city, intcreation, intalbum)

	var tmpl *template.Template
	if len(Donnees.Artist) == 0 {
		tmpl = template.Must(template.ParseFiles("./static/noresult.html"))
	} else {
		tmpl = template.Must(template.ParseFiles("./static/artistes.html"))
	}
	tmpl.Execute(w, Donnees)
	return
}

func HandleInfos(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	Iid, _ := strconv.Atoi(id)
	Iid = Iid - 1

	donneesArtist := gpf.InfoArtist(Alldatas, Iid)

	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("./static/info.html"))
	tmpl.Execute(w, donneesArtist)
	return
}
