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
	Carrier          ShortCarrier     `json:"carrier"`
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

type Direction struct {
	Code  string `json:"code"`
	Title string `json:"title"`
}

func (y *YandexRapsApi) Schedule(params map[string]string) (response ScheduleResponse, err error) {
	err = y.getJson("v3.0/schedule/", params, &response)
	return
}
