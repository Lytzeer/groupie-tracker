package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var api DATAS

type API struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relations string `json:"relation"`
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

type LOCATION struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type DATE struct {
	Index []struct {
		Id    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

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

var Ap API

func SetData() {
	/*SET artists*/
	responseArtists, _ := http.Get(Ap.Artists)
	responseDataArtists, _ := ioutil.ReadAll(responseArtists.Body)
	Ar := []ARTIST{}
	json.Unmarshal(responseDataArtists, &Ar)

	/*SET Dates*/
	responseDate, _ := http.Get(Ap.Dates)
	responseDataDate, _ := ioutil.ReadAll(responseDate.Body)
	Da := []DATE{}
	json.Unmarshal(responseDataDate, &Da)

	/*SET Location*/
	responseLocation, _ := http.Get(Ap.Locations)
	responseDataLocation, _ := ioutil.ReadAll(responseLocation.Body)
	GL := []GetLocation{}
	json.Unmarshal(responseDataLocation, &GL)
}

func main() {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	json.Unmarshal(responseData, &Ap)
	SetData()
	// fmt.Println(Ap.Artists)
	// fmt.Println(Ap.Locations)
	// fmt.Println(Ap.Dates)
	// fmt.Println(Ap.Relations)
	// response2, err := http.Get(Ap.Artists)
	// responseData2, err := ioutil.ReadAll(response2.Body)
	// Ar := []ARTIST{}
	// json.Unmarshal(responseData2, &Ar)

	// for _, artist := range Ar {
	// 	fmt.Println(artist.Rlations)
	// 	fmt.Println()
	// }

	// response3, err := http.Get(Ap.Dates)
	// fmt.Println(Ap.Dates)
	// responseData3, err := ioutil.ReadAll(response3.Body)
	// Da := DATE{}
	// json.Unmarshal(responseData3, &Da)
	// for _, dates := range Da.Index {
	// 	fmt.Println(dates.Dates)
	// 	fmt.Println()
	// }
	// for _, dates := range Re.Index {
	// 	fmt.Println(dates.Dates_location)
	// 	fmt.Println()
	// response3, err := http.Get(Ap.Dates)
	// fmt.Println(Ap.Dates)
	// responseData3, err := ioutil.ReadAll(response3.Body)
	// Da := DATE{}
	// json.Unmarshal(responseData3, &Da)
	// for _, dates := range Da.Index {
	// 	fmt.Println(dates.Dates)
	// 	fmt.Println()
	// }

	// response4, err := http.Get(Ap.Locations)
	// fmt.Println(Ap.Dates)
	// responseData4, err := ioutil.ReadAll(response4.Body)
	// GL := GetLocation{}
	// json.Unmarshal(responseData4, &GL)
	// //fmt.Println(GL)
	// for _, G := range GL.Index {
	// 	fmt.Println(G.Locations)
	// 	fmt.Println()
	// }
}

type DATAS struct {
	Id            int
	Locations     []string
	Dates         string
	IdA           int
	Image         string
	Name          string
	Members       []string
	Creation_date int
	First_ablbum  string
	LocationsA    string
	Concert_dates string
	Rlations      string
	IdL           int
	LocationsL    []string
	DatesL        string
	Index         []struct {
		Id    int
		Dates []string
	}
	IndexM []struct {
		Id        int
		Locations []string
	}
	IndexJ []struct {
		Id             int
		DatesLocations struct {
			Location string
			Dates    []string
		}
	}
}
