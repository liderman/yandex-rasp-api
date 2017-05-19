package yandex_rasp_api

type ScheduleResponse struct {
	Pagination        Pagination  `json:"pagination"`
	Schedule          []Schedule  `json:"schedule"`
	ScheduleDirection Direction   `json:"schedule_direction"`
	Directions        []Direction `json:"directions"`
	Station           Station     `json:"station"`
	Date              string      `json:"date"`
}

type Schedule struct {
	ExceptDays    string         `json:"except_days"`
	Arrival       string         `json:"arrival"`
	ArrivalTime   string         `json:"arrival_time"`
	Direction     string         `json:"direction"`
	Thread        ScheduleThread `json:"thread"`
	Platform      string         `json:"platform"`
	Days          string         `json:"days"`
	Stops         string         `json:"stops"`
	DepartureTime string         `json:"departure_time"`
	Terminal      string         `json:"terminal"`
	IsFuzzy       bool           `json:"is_fuzzy"`
}

type ScheduleThread struct {
	Carrier          CarrierItem       `json:"carrier"`
	TransportType    string            `json:"transport_type"`
	Uid              string            `json:"uid"`
	Title            string            `json:"title"`
	Vehicle          string            `json:"vehicle"`
	Number           string            `json:"number"`
	ShortTitle       string            `json:"short_title"`
	ExpressType      string            `json:"express_type"`
	TransportSubtype map[string]string `json:"transport_subtype"`
	ThreadMethodLink string            `json:"thread_method_link"`
}

type ThreadItem struct {
	Carrier          Carrier           `json:"carrier"`
	TransportType    string            `json:"transport_type"`
	Uid              string            `json:"uid"`
	Title            string            `json:"title"`
	Vehicle          string            `json:"vehicle"`
	Number           string            `json:"number"`
	ShortTitle       string            `json:"short_title"`
	ExpressType      string            `json:"express_type"`
	TransportSubtype map[string]string `json:"transport_subtype"`
	ThreadMethodLink string            `json:"thread_method_link"`
}

type CarrierItem struct {
	Code  int    `json:"code"`
	Codes Codes  `json:"codes"`
	Title string `json:"title"`
}

type Codes struct {
	Icao   string `json:"icao"`
	Sirena string `json:"sirena"`
	Iata   string `json:"iata"`
}

type Direction struct {
	Code  string `json:"code"`
	Title string `json:"title"`
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

func (y *YandexRapsApi) Schedule(params map[string]string) (response ScheduleResponse, err error) {
	err = y.getJson("v3.0/schedule/", params, &response)
	return
}
