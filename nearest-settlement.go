package yandex_rasp_api

type NearestSettlement struct {
	Distance     float64 `json:"distance"`
	Code         string  `json:"code"`
	Title        string  `json:"title"`
	PopularTitle string  `json:"popular_title"`
	ShortTitle   string  `json:"short_title"`
	Lat          float64 `json:"lat"`
	Lng          float64 `json:"lng"`
	Type         string  `json:"type"`
}

func (y *YandexRapsApi) NearestSettlement(params map[string]string) (response NearestSettlement, err error) {
	err = y.getJson("v3.0/nearest_settlement/", params, &response)
	return
}
