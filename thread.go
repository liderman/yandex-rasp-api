package yandex_rasp_api

type ThreadCard struct {
	ExceptDays       string           `json:"except_days"`
	ArrivalDate      string           `json:"arrival_date"`
	From             string           `json:"from"`
	Uid              string           `json:"uid"`
	Title            string           `json:"title"`
	DepartureDate    string           `json:"departure_date"`
	StartTime        string           `json:"start_time"`
	Number           string           `json:"number"`
	ShortTitle       string           `json:"short_title"`
	Days             string           `json:"days"`
	To               string           `json:"to"`
	Carrier          ThreadCarrier    `json:"carrier"`
	TransportType    string           `json:"transport_type"`
	Stops            []Stop           `json:"stops"`
	Vehicle          string           `json:"vehicle"`
	StartDate        string           `json:"start_date"`
	TransportSubtype TransportSubtype `json:"transport_subtype"`
	ExpressType      string           `json:"express_type"`
}

type TransportSubtype struct {
	Color string `json:"color"`
	Code  string `json:"code"`
	Title string `json:"title"`
}

type Stop struct {
	Arrival   string        `json:"arrival"`
	Departure string        `json:"departure"`
	Terminal  string        `json:"terminal"`
	Platform  string        `json:"platform"`
	Station   ThreadStation `json:"station"`
	StopTime  string        `json:"stop_time"`
	Duration  float32       `json:"duration"`
}

type ThreadStation struct {
	Codes        StationCodes `json:"codes"`
	Title        string       `json:"title"`
	PopularTitle string       `json:"popular_title"`
	ShortTitle   string       `json:"short_title"`
	Code         string       `json:"code"`
	Type         string       `json:"type"`
}

type StationCodes struct {
	Express string `json:"express"`
	Yandex  string `json:"yandex"`
	Esr     string `json:"esr"`
}

type ThreadCarrier struct {
	Code    int      `json:"code"`
	Offices []string `json:"offices"`
	Codes   Codes    `json:"codes"`
	Title   string   `json:"title"`
}

func (y *YandexRapsApi) Thread(params map[string]string) (response ThreadCard, err error) {
	err = y.getJson("v3.0/thread/", params, &response)
	return
}
