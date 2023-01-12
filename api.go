package groupie

import (
	"encoding/json"
	"fmt"
	gpd "groupie/datas"
	"io/ioutil"
	"net/http"
)

type API struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relations string `json:"relation"`
}

type ArtistInfos struct {
	Artist   interface{}
	Location interface{}
	Date     interface{}
	Relation interface{}
}

var Ap API

var Donnees gpd.DATAS

func GetDatas() (gpd.DATE, []gpd.ARTIST, gpd.GetLocation, gpd.RELATION) {
	response, _ := http.Get("https://groupietrackers.herokuapp.com/api")

	responseData, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(responseData, &Ap)
	fmt.Println(Ap.Artists)
	fmt.Println(Ap.Locations)
	fmt.Println(Ap.Dates)
	fmt.Println(Ap.Relations)

	/*SET Dates*/
	responseDates, _ := http.Get(Ap.Dates)
	fmt.Println(Ap.Dates)
	responseDataDates, _ := ioutil.ReadAll(responseDates.Body)
	Da := gpd.DATE{}
	json.Unmarshal(responseDataDates, &Da)

	//fmt.Println(Da)

	/*SET artists*/
	responseArtists, _ := http.Get(Ap.Artists)
	responseDataArtists, _ := ioutil.ReadAll(responseArtists.Body)
	Ar := []gpd.ARTIST{}
	json.Unmarshal(responseDataArtists, &Ar)

	//fmt.Println(Ar)

	/*SET Location*/
	responseLocation, _ := http.Get(Ap.Locations)
	responseDataLocation, _ := ioutil.ReadAll(responseLocation.Body)
	GL := gpd.GetLocation{}
	json.Unmarshal(responseDataLocation, &GL)
	//fmt.Println(GL)

	responseRelation, _ := http.Get(Ap.Relations)
	responseDataRelation, _ := ioutil.ReadAll(responseRelation.Body)
	Re := gpd.RELATION{}
	json.Unmarshal(responseDataRelation, &Re)

	return Da, Ar, GL, Re
}

func SetData(d gpd.DATE, a []gpd.ARTIST, l gpd.GetLocation, relation gpd.RELATION, donnes gpd.DATAS) gpd.DATAS {
	Donnees.Date = d.Index
	for i := 0; i < (len(a)); i++ {
		Donnees.Artist = append(Donnees.Artist, a[i])
	}
	Donnees.Location = l.Index
	Donnees.Relation = relation.Index

	return Donnees

}
