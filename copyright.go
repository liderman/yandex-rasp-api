package yandex_rasp_api

type CopyrightResponse struct {
	Copyright Copyright `json:"copyright"`
}

type Copyright struct {
	Url    string `json:"url"`
	LogoVm string `json:"logo_vm"`
	LogoVd string `json:"logo_vd"`
	LogoHy string `json:"logo_hy"`
	LogoHd string `json:"logo_hd"`
	LogoVy string `json:"logo_vy"`
	LogoHm string `json:"logo_hm"`
	Text   string `json:"text"`
}

func (y *YandexRapsApi) Copyright() (response CopyrightResponse, err error) {
	err = y.getJson("v3.0/copyright/", map[string]string{}, &response)
	return
}
