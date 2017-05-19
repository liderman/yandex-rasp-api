package yandex_rasp_api

type NearestStationsResponse struct {
	Pagination Pagination        `json:"pagination"`
	Stations   []NearestStations `json:"stations"`
}

type TypeChoiceLink struct {
	DesktopUrl string `json:"desktop_url"`
	TouchUrl   string `json:"touch_url"`
}

type NearestStations struct {
	Distance      float64                   `json:"distance"`
	Code          string                    `json:"code"`
	StationType   string                    `json:"station_type"`
	TypeChoices   map[string]TypeChoiceLink `json:"type_choices"`
	Title         string                    `json:"title"`
	PopularTitle  string                    `json:"popular_title"`
	ShortTitle    string                    `json:"short_title"`
	Majority      int                       `json:"majority"`
	TransportType string                    `json:"transport_type"`
	Lat           float64                   `json:"lat"`
	Lng           float64                   `json:"lng"`
	Type          string                    `json:"type"`
}

func (y *YandexRapsApi) NearestStations(params map[string]string) (response NearestStationsResponse, err error) {
	err = y.getJson("v3.0/nearest_stations/", params, &response)
	return
}
