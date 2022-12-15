package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type API struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relations string `json:"relation"`
}

func main() {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api")
	f, err := os.Create("out.txt")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	f.WriteString(string(responseData))
	if err != nil {
		log.Fatal(err)
	}
	Ap := API{}
	jsonErr := json.Unmarshal(responseData, &Ap)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(string(responseData))
	fmt.Println(Ap.Artists)
	fmt.Println(Ap.Locations)
	fmt.Println(Ap.Dates)
	fmt.Println(Ap.Relations)
}
