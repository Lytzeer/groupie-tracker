package groupie

type FeatureCollection struct {
	Features []struct {
		Center []float64 `json:"center"`
	} `json:"features"`
}
