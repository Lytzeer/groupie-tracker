package groupie

import (
	"encoding/json"
	gpd "groupie/datas"
	"io/ioutil"
	"net/http"
	"strconv"
)

func CreateMap(Alldatas gpd.DATAS, index int) string {
	var Structcoordonnees []gpd.FeatureCollection
	carte := ""

	for _, loc := range Alldatas.Location[index].Locations {

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

		Structcoordonnees = append(Structcoordonnees, test)
	}

	carte = "https://api.mapbox.com/styles/v1/mapbox/streets-v12/static/"

	for i, l := range Structcoordonnees {
		coordonnees1 := strconv.FormatFloat(l.Features[0].Center[0], 'g', 9, 32)
		coordonnees2 := strconv.FormatFloat(l.Features[0].Center[1], 'g', 9, 32)
		if i == len(Structcoordonnees)-1 {
			carte += "pin-l-music+f74e4e(" + coordonnees1 + "," + coordonnees2 + ")" + "/20,0,0/500x500?access_token=pk.eyJ1IjoibWF0c3VlbGwiLCJhIjoiY2xkbjNoMTgzMGZseDN1bHgybjgwbmFnOCJ9.qUR-JuwsRM69PeuHEcVo4A"
		} else {
			carte += "pin-l-music+f74e4e(" + coordonnees1 + "," + coordonnees2 + "),"
		}
	}

	return carte
}
