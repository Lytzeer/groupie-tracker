package groupie

import gpd "groupie/datas"

func Displaydata(i int, Donnees gpd.DATAS, Alldatas gpd.DATAS) []gpd.ARTIST {
	var Artist gpd.ARTIST
	Artist.Name = Alldatas.Artist[i].Name
	Artist.Image = Alldatas.Artist[i].Image
	Artist.Id = Alldatas.Artist[i].Id
	Donnees.Artist = append(Donnees.Artist, Artist)
	return (Donnees.Artist)
}
