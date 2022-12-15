package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

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
	Id             int         `json:"id"`
	Dates_location interface{} `json:"datesLocations"`
}

func main() {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	Ap := API{}
	json.Unmarshal(responseData, &Ap)
	fmt.Println(Ap.Artists)
	fmt.Println(Ap.Locations)
	fmt.Println(Ap.Dates)
	fmt.Println(Ap.Relations)
	//response2, err := http.Get(Ap.Artists)
	//responseData2, err := ioutil.ReadAll(response2.Body)
	//Ar := []ARTIST{}
	//json.Unmarshal(responseData2, &Ar)
	//
	//for _, artist := range Ar {
	//	fmt.Println(artist.Rlations)
	//	fmt.Println()
	//}
	//
	response3, err := http.Get(Ap.Dates)
	fmt.Println(Ap.Dates)
	responseData3, err := ioutil.ReadAll(response3.Body)
	Da := DATE{}
	json.Unmarshal(responseData3, &Da)
	for _, dates := range Da.Index {
		fmt.Println(dates.Dates)
		fmt.Println()
	}

	response4, err := http.Get(Ap.Locations)
	fmt.Println(Ap.Dates)
	responseData4, err := ioutil.ReadAll(response4.Body)
	GL := GetLocation{}
	json.Unmarshal(responseData4, &GL)
	fmt.Println(GL)
	// for _, G := range GL.Index {
	// 	//fmt.Println(G.DatesLocations)
	// 	//fmt.Println()
	// }
}
