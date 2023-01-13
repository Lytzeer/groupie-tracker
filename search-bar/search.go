package groupie

import (
	"fmt"
	gpd "groupie/datas"
	"strconv"
)

func GetAll(d gpd.DATAS) {
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

	fmt.Println(Names)
	fmt.Println(Members)
	fmt.Println(Positions)
	fmt.Println(FirstAlbum)
	fmt.Println(Creation)
}
