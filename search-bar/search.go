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
	A := []string{}
	All := []string{}

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

	A = append(A, Names...)
	A = append(A, Members...)
	A = append(A, Positions...)
	A = append(A, FirstAlbum...)
	A = append(A, Creation...)
	for _, Ele := range A {
		if !Isin(Ele, All) {
			All = append(All, Ele)
		}
	}

	return All
}

func Isin(ele string, tab []string) bool {

	for _, element := range tab {
		if element == ele {
			return true
		}
	}
	return false
}
