package groupie

import (
	gpd "groupie/datas"
)

func SearchDatas(search string, intSearch int, Alldatas gpd.DATAS) gpd.DATAS {
	var sdatas gpd.DATAS
	members := []string{}
	if intSearch == 0 {
		for i := 0; i < (len(Alldatas.Artist)); i++ {
			for _, jsp := range Alldatas.Location[i].Locations {
				if (jsp == search || Alldatas.Artist[i].Name == search || Alldatas.Artist[i].First_ablbum == search) && !Isin(Alldatas.Artist[i].Name, members) {
					sdatas.Artist = Displaydata(i, sdatas, Alldatas)
					members = append(members, Alldatas.Artist[i].Name)
				}
			}
			for _, member := range Alldatas.Artist[i].Members {
				if member == search && !Isin(member, members) {
					sdatas.Artist = Displaydata(i, sdatas, Alldatas)
					members = append(members, Alldatas.Artist[i].Name)
				}
			}
		}
	}

	if intSearch != 0 {
		for i := 0; i < len(Alldatas.Artist); i++ {
			if Alldatas.Artist[i].Creation_date == intSearch && !Isin(Alldatas.Artist[i].Name, members) {
				sdatas.Artist = Displaydata(i, sdatas, Alldatas)
				members = append(members, Alldatas.Artist[i].Name)
			}
		}
	}

	sdatas.All = Alldatas.All
	sdatas.Country = Alldatas.Country
	return sdatas
}
