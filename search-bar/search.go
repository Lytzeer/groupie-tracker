package groupie

import (
	gpd "groupie/datas"
	"strconv"
)

func GetAll(d gpd.DATAS) []string {
	Names := []string{}
	Members := []string{}
	Positions := []string{}
	FirstAlbum := []string{}
	Creation := []string{}

	for _, artist := range d.Artist {
		Names = append(Names, artist.Name)
		Members = append(Members, artist.Members...)
		FirstAlbum = append(FirstAlbum, artist.First_ablbum)
		Creation = append(Creation, strconv.Itoa(artist.Creation_date))
	}

	for _, locations := range d.Location {
		for _, a := range locations.Locations {
			Positions = append(Positions, a)
		}

	}

	d.All = append(d.All, Names...)
	d.All = append(d.All, Members...)
	d.All = append(d.All, Positions...)
	d.All = append(d.All, FirstAlbum...)
	d.All = append(d.All, Creation...)

	return d.All
}
