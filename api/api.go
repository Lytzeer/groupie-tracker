package groupie

import (
	"encoding/json"
	gpd "groupie/datas"
	gpf "groupie/func"
	gps "groupie/search-bar"
	"io/ioutil"
	"net/http"
	"strings"
)

var Ap gpd.API

var Donnees gpd.DATAS

func GetDatas() (gpd.DATE, []gpd.ARTIST, gpd.GetLocation, gpd.RELATION) {
	//On récupère toutes les données et on les stockent dans leurs structures avant de les ajouter dans une seule et même strcuture
	response, _ := http.Get("https://groupietrackers.herokuapp.com/api")

	responseData, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(responseData, &Ap)

	/*SET Dates*/
	responseDates, _ := http.Get(Ap.Dates)
	responseDataDates, _ := ioutil.ReadAll(responseDates.Body)
	Da := gpd.DATE{}
	json.Unmarshal(responseDataDates, &Da)

	/*SET artists*/
	responseArtists, _ := http.Get(Ap.Artists)
	responseDataArtists, _ := ioutil.ReadAll(responseArtists.Body)
	Ar := []gpd.ARTIST{}
	json.Unmarshal(responseDataArtists, &Ar)

	/*SET Locations*/
	responseLocation, _ := http.Get(Ap.Locations)
	responseDataLocation, _ := ioutil.ReadAll(responseLocation.Body)
	GL := gpd.GetLocation{}
	json.Unmarshal(responseDataLocation, &GL)

	/*SET Relations*/
	responseRelation, _ := http.Get(Ap.Relations)
	responseDataRelation, _ := ioutil.ReadAll(responseRelation.Body)
	Re := gpd.RELATION{}
	json.Unmarshal(responseDataRelation, &Re)

	return Da, Ar, GL, Re
}

func SetData(d gpd.DATE, a []gpd.ARTIST, l gpd.GetLocation, relation gpd.RELATION) gpd.DATAS {

	Donnees.Date = d.Index
	for i := 0; i < (len(a)); i++ {
		Donnees.Artist = append(Donnees.Artist, a[i])
	}
	Donnees.Location = l.Index
	Donnees.Relation = relation.Index

	//Tableau avec toutes les données que l'on peut chercher dans la barre de recherche
	Search := make([][][]string, 52)
	for i := 0; i < len(Donnees.Location); i++ {
		Search[i] = make([][]string, len(Donnees.Relation[i].DatesLocations))
		cpt := 0
		for loc, dates := range Donnees.Relation[i].DatesLocations {
			Search[i][cpt] = append(Search[i][cpt], loc+" : ")
			for j := 0; j < len(dates); j++ {
				if j == 0 {
					Search[i][cpt] = append(Search[i][cpt], dates[j])
				} else if j >= 1 {
					Search[i][cpt] = append(Search[i][cpt], dates[j])
				} else {
					Search[i][cpt] = append(Search[i][cpt], dates[j])
				}

			}
			cpt++
		}
	}

	Donnees.Locs = Search

	for i := 0; i < (len(Donnees.Artist)); i++ {
		Donnees.NbMembers = append(Donnees.NbMembers, len(Donnees.Artist[i].Members))
	}

	Donnees.All = gps.GetAll(Donnees)

	var countries []string
	for i := 0; i < (len(Donnees.Location)); i++ {
		for j := 0; j < (len(Donnees.Location[i].Locations)); j++ {
			country := strings.Split(Donnees.Location[i].Locations[j], "-")[1]
			if !gpf.Isin(country, countries) {
				countries = append(countries, country)
			}

		}
	}
	Donnees.Country = countries

	return Donnees

}
