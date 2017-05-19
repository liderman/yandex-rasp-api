package yandex_rasp_api

type CarrierResponse struct {
	Carriers []Carrier `json:"carriers"`
}

type Carrier struct {
	Code     int      `json:"code"`
	Contacts string   `json:"contacts"`
	Url      string   `json:"url"`
	Title    string   `json:"title"`
	Phone    string   `json:"phone"`
	Codes    Codes    `json:"codes"`
	Offices  []string `json:"offices"`
	Address  string   `json:"address"`
	Logo     string   `json:"logo"`
	Email    string   `json:"email"`
}

type ShortCarrier struct {
	Code  int    `json:"code"`
	Codes Codes  `json:"codes"`
	Title string `json:"title"`
}

func (y *YandexRapsApi) Carrier(params map[string]string) (response CarrierResponse, err error) {
	err = y.getJson("v3.0/carrier/", params, &response)
	return
}
