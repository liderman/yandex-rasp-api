package yandex_rasp_api

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"testing"
)

type CarrierDataFetcher struct{}

func (c *CarrierDataFetcher) Fetch(url string) (resp io.ReadCloser, err error) {
	err = nil
	resp = ioutil.NopCloser(bytes.NewReader([]byte(
		`{"carriers":[{"code":680,"contacts":"","url":"http://www.turkishairlines.com/","title":"Turkish Airl` +
			`ines","phone":"444 08 49","codes":{"icao":"THY","sirena":null,"iata":"TK"},"offices":[],"address":"B` +
			`akırköy/ İstanbul","logo":"//yastatic.net/rasp/media/data/company/logo/thy_kopya.jpg","email":""},{"` +
			`code":5483,"contacts":"","url":"http://www.anadolujet.com/","title":"AnadoluJet","phone":"444 25 38"` +
			`,"codes":{"icao":"AJA","sirena":null,"iata":"TK"},"offices":[],"address":"Bakırköy/ İstanbul","logo"` +
			`:"//yastatic.net/rasp/media/data/company/logo/anadolujet.png","email":"musteri@anadolujet.com"}]}`,
	)))
	return
}

func TestCarrier(t *testing.T) {
	api := NewYandexRapsApi("TEST_TOKEN")
	api.Fetcher = &CarrierDataFetcher{}
	reps, err := api.Carrier(map[string]string{
		"lang":   "ru_RU",
		"code":   "TK",
		"system": "iata",
	})
	if err != nil {
		t.Error(err)
	}

	data := CarrierResponse{
		Carriers: []Carrier{{
			Code:     680,
			Contacts: "",
			Url:      "http://www.turkishairlines.com/",
			Title:    "Turkish Airlines",
			Phone:    "444 08 49",
			Codes: Codes{
				Icao:   "THY",
				Sirena: "",
				Iata:   "TK",
			},
			Offices: []string{},
			Address: "Bakırköy/ İstanbul",
			Logo:    "//yastatic.net/rasp/media/data/company/logo/thy_kopya.jpg",
			Email:   "",
		}, {
			Code:     5483,
			Contacts: "",
			Url:      "http://www.anadolujet.com/",
			Title:    "AnadoluJet",
			Phone:    "444 25 38",
			Codes: Codes{
				Icao:   "AJA",
				Sirena: "",
				Iata:   "TK",
			},
			Offices: []string{},
			Address: "Bakırköy/ İstanbul",
			Logo:    "//yastatic.net/rasp/media/data/company/logo/anadolujet.png",
			Email:   "musteri@anadolujet.com",
		}},
	}
	assert.Equal(t, reps, data, "Carrier is not Equal!")
}
