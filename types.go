package yandex_rasp_api

type Pagination struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Codes struct {
	Icao   string `json:"icao"`
	Sirena string `json:"sirena"`
	Iata   string `json:"iata"`
}

type Station struct {
	Code          string `json:"code"`
	StationType   string `json:"station_type"`
	Title         string `json:"title"`
	PopularTitle  string `json:"popular_title"`
	ShortTitle    string `json:"short_title"`
	TransportType string `json:"transport_type"`
	Type          string `json:"type"`
}

type TransportSubtype struct {
	Color string `json:"color"`
	Code  string `json:"code"`
	Title string `json:"title"`
}
