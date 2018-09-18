package meander

type Place struct {
	*googleGeometry `json:"geometry"`
	Name            string         `json:"name"`
	Icon            string         `json:"icon"`
	Photos          []*googlePhoto `json:"photos"`
	Vicinity        string         `json:"vicinity"`
}

type Query struct {
	Lat          float64
	Lng          float64
	Journey      []string
	Radius       int
	CostRangeStr string
}

type googleResponse struct {
	Result []*Place `json:"results"`
}

type googleGeometry struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type googlePhoto struct {
	PhotoRef string `json:"photo_reference"`
	URL      string `json:"url"`
}

func (p *Place) Public() interface{} {
	return map[string]interface{}{
		"name":     p.Name,
		"icon":     p.Icon,
		"photos":   p.Photos,
		"vicinity": p.Vicinity,
		"lat":      p.Lat,
		"lng":      p.Lng,
	}
}

var APIKey string

func (q *Query) find(types string) (*googleResponse, error) {
	u := "https://maps.googleapis.com/maps/api/place/nearbysearch/json"
}
