package groupie

import (
	gpd "groupie/datas"
	"strconv"
	"strings"
)

func FilterDatas(Alldatas gpd.DATAS, buttonAll string, tabbutton []int, city string, intcreation int, intalbum int) gpd.DATAS {
	var Donnees gpd.DATAS
	var splitalbum []int
	name := []string{}

	for i := 0; i < (len(Alldatas.Artist)); i++ {
		splitalbumel := strings.Split(Alldatas.Artist[i].First_ablbum, "-")[2]
		splitalbumstr, _ := strconv.Atoi(splitalbumel)
		splitalbum = append(splitalbum, splitalbumstr)
	}
	for i := 0; i < len(Alldatas.Artist); i++ {
		for _, jsp := range Alldatas.Location[i].Locations {
			countries := strings.Split(jsp, "-")[1]
			for j := 0; j < 8; j++ {
				if string(buttonAll) != "All" && city != "All" {
					if j+1 == tabbutton[j] {
						if len(Alldatas.Artist[i].Members) == tabbutton[j] && Alldatas.Artist[i].Creation_date >= intcreation && int(splitalbum[i]) >= intalbum && countries == city && !Isin(Alldatas.Artist[i].Name, name) {
							Donnees.Artist = Displaydata(i, Donnees, Alldatas)
							name = append(name, Alldatas.Artist[i].Name)
						}
					}
				} else if buttonAll == "All" && city != "All" {
					if Alldatas.Artist[i].Creation_date >= intcreation && int(splitalbum[i]) >= intalbum && countries == city && !Isin(Alldatas.Artist[i].Name, name) {
						Donnees.Artist = Displaydata(i, Donnees, Alldatas)
						name = append(name, Alldatas.Artist[i].Name)
					}
				} else if buttonAll == "All" && city == "All" {
					if Alldatas.Artist[i].Creation_date >= intcreation && int(splitalbum[i]) >= intalbum && !Isin(Alldatas.Artist[i].Name, name) {
						Donnees.Artist = Displaydata(i, Donnees, Alldatas)
						name = append(name, Alldatas.Artist[i].Name)
					}
				} else if string(buttonAll) != "All" && city == "All" {
					if len(Alldatas.Artist[i].Members) == tabbutton[j] && Alldatas.Artist[i].Creation_date >= intcreation && int(splitalbum[i]) >= intalbum && !Isin(Alldatas.Artist[i].Name, name) {
						Donnees.Artist = Displaydata(i, Donnees, Alldatas)
						name = append(name, Alldatas.Artist[i].Name)
					}
				}
			}
		}
	}

	Donnees.All = Alldatas.All
	Donnees.Country = Alldatas.Country
	return Donnees
}
