package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GetLocation struct {
	Index []struct {
		Id        int      `json:"id"`
		Locations []string `json:"locations"`
	} `json:"index"`
}

type RELATION struct {
	Index []struct {
		Id             int `json:"id"`
		DatesLocations struct {
			Location string
			Dates    []string
		} `json:"datesLocations"`
	} `json:"index"`
}

type API struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relations string `json:"relation"`
}

type DATE struct {
	Index []struct {
		Id    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type ARTIST struct {
	Id            int      `json:"id"`
	Image         string   `json:"image"`
	Name          string   `json:"name"`
	Members       []string `json:"members"`
	Creation_date int      `json:"creationDate"`
	First_ablbum  string   `json:"firstAlbum"`
	Locations     string   `json:"locations"`
	Concert_dates string   `json:"concertDates"`
	Rlations      string   `json:"relations"`
}

type DATAS struct {
	Date []struct {
		Id    int      "json:\"id\""
		Dates []string "json:\"dates\""
	}
	Artist   []ARTIST
	Location []struct {
		Id        int      "json:\"id\""
		Locations []string "json:\"locations\""
	}
}

var Ap API

var Donnees DATAS

func main() {
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
	Da := DATE{}
	json.Unmarshal(responseDataDates, &Da)

	//fmt.Println(Da)

	/*SET artists*/
	responseArtists, _ := http.Get(Ap.Artists)
	responseDataArtists, _ := ioutil.ReadAll(responseArtists.Body)
	Ar := []ARTIST{}
	json.Unmarshal(responseDataArtists, &Ar)

	//fmt.Println(Ar)

	/*SET Location*/
	responseLocation, _ := http.Get(Ap.Locations)
	responseDataLocation, _ := ioutil.ReadAll(responseLocation.Body)
	GL := GetLocation{}
	json.Unmarshal(responseDataLocation, &GL)
	//fmt.Println(GL)

	SetData(Da, Ar, GL)
}

func SetData(d DATE, a []ARTIST, l GetLocation) {
	Donnees.Date = d.Index
	for i := 0; i < (len(a)); i++ {
		Donnees.Artist = append(Donnees.Artist, a[i])
	}
	Donnees.Location = l.Index

	fmt.Println(Donnees)
}
