package groupie

import (
	gpd "groupie/datas"
	gpm "groupie/map"
)

func InfoArtist(Alldatas gpd.DATAS, index int) gpd.ArtistInfos {
	dat := Alldatas.Date[index]
	donneesArtist := gpd.ArtistInfos{}
	donneesArtist.Artist = Alldatas.Artist[index]
	donneesArtist.Location = Alldatas.Location[index]
	donneesArtist.All = Alldatas.Locs[index]

	//Suppresion des * dans les dates
	for i := 0; i < len(dat.Dates); i++ {
		if string(dat.Dates[i][0]) == "*" {
			dat.Dates[i] = dat.Dates[i][1:]
		}
	}

	donneesArtist.Date = dat
	donneesArtist.Carte = gpm.CreateMap(Alldatas, index)

	return donneesArtist
}
