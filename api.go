package groupie

import (
	"encoding/json"
	gpd "groupie/datas"
	"io/ioutil"
	"net/http"
)

var Ap gpd.API

var Donnees gpd.DATAS

func GetDatas() (gpd.DATE, []gpd.ARTIST, gpd.GetLocation, gpd.RELATION) {
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

	/*SET Location*/
	responseLocation, _ := http.Get(Ap.Locations)
	responseDataLocation, _ := ioutil.ReadAll(responseLocation.Body)
	GL := gpd.GetLocation{}
	json.Unmarshal(responseDataLocation, &GL)

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

	All := make([][][]string, 52)
	for i := 0; i < len(Donnees.Location); i++ {
		All[i] = make([][]string, len(Donnees.Relation[i].DatesLocations))
		cpt := 0
		for loc, dates := range Donnees.Relation[i].DatesLocations {
			All[i][cpt] = append(All[i][cpt], loc)
			for j := 0; j < len(dates); j++ {
				All[i][cpt] = append(All[i][cpt], dates[j])
			}
			cpt++
		}
	}

	Donnees.Locs = All

	return Donnees

}
