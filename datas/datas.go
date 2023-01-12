package groupie

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
	Relation []struct {
		Id             int                 "json:\"id\""
		DatesLocations map[string][]string `json:"datesLocations"`
	}
}
