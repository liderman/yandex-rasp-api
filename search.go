package yandex_rasp_api

type SearchResponse struct {
	Pagination Pagination `json:"pagination"`
	Segments   []Segment  `json:"segments"`
	Search     Search     `json:"search"`
}

type Segment struct {
	Arrival           string      `json:"arrival"`
	From              Station     `json:"from"`
	Thread            Thread      `json:"thread"`
	DeparturePlatform string      `json:"departure_platform"`
	Departure         string      `json:"departure"`
	Stops             string      `json:"stops"`
	DepartureTerminal string      `json:"departure_terminal"`
	To                Station     `json:"to"`
	HasTransfers      bool        `json:"has_transfers"`
	TicketsInfo       TicketsInfo `json:"tickets_info"`
	Duration          float32     `json:"duration"`
	ArrivalTerminal   string      `json:"arrival_terminal"`
	StartDate         string      `json:"start_date"`
	ArrivalPlatform   string      `json:"arrival_platform"`
}

type Thread struct {
	Carrier          SearchCarrier    `json:"carrier"`
	TransportType    string           `json:"transport_type"`
	Uid              string           `json:"uid"`
	Title            string           `json:"title"`
	Vehicle          string           `json:"vehicle"`
	Number           string           `json:"number"`
	ShortTitle       string           `json:"short_title"`
	ExpressType      string           `json:"express_type"`
	TransportSubtype TransportSubtype `json:"transport_subtype"`
	ThreadMethodLink string           `json:"thread_method_link"`
}

type SearchCarrier struct {
	Code     int    `json:"code"`
	Contacts string `json:"contacts"`
	Url      string `json:"url"`
	Title    string `json:"title"`
	Phone    string `json:"phone"`
	Codes    Codes  `json:"codes"`
	Address  string `json:"address"`
	LogoSvg  string `json:"logo_svg"`
	Logo     string `json:"logo"`
	Email    string `json:"email"`
}

type TicketsInfo struct {
	EtMarker bool    `json:"et_marker"`
	Places   []Place `json:"places"`
}

type Place struct {
	Currency string `json:"currency"`
	Price    Price  `json:"price"`
	Name     string `json:"name"`
}

type Price struct {
	Cents int      `json:"cents"`
	Whole Location `json:"whole"`
}

type Search struct {
	Date string   `json:"date"`
	To   Location `json:"to"`
	From Location `json:"from"`
}

type Location struct {
	Code         string `json:"code"`
	Type         string `json:"type"`
	PopularTitle string `json:"popular_title"`
	ShortTitle   string `json:"short_title"`
	Title        string `json:"title"`
}

func (y *YandexRapsApi) Search(params map[string]string) (response SearchResponse, err error) {
	err = y.getJson("v3.0/search/", params, &response)
	return
}
