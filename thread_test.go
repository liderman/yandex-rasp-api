package yandex_rasp_api

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"testing"
)

type ThreadDataFetcher struct{}

func (c *ThreadDataFetcher) Fetch(url string) (resp io.ReadCloser, err error) {
	err = nil
	resp = ioutil.NopCloser(bytes.NewReader([]byte(
		`{"except_days":"","arrival_date":null,"from":null,"uid":"NN-324_0_c9_12","title":"Симферополь — Моск` +
			`ва","departure_date":null,"start_time":"01:10","number":"NN 324","short_title":"Симферополь — Москва` +
			`","days":"19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31 мая, 1, 2, 3, 6, 9, 10, 13 июня, …","to` +
			`":null,"carrier":{"code":9,"offices":[],"codes":{"icao":null,"sirena":"НН","iata":"NN"},"title":"ВИМ` +
			`-Авиа"},"transport_type":"plane","stops":[{"arrival":null,"departure":"2017-05-19 01:10:00","termina` +
			`l":null,"platform":"","station":{"code":"s9600396","station_type":"аэропорт","title":"Симферополь","` +
			`popular_title":"","short_title":"","codes":{"yandex":"s9600396"},"transport_type":"plane","type":"st` +
			`ation"},"stop_time":null,"duration":0.0},{"arrival":"2017-05-19 03:30:00","departure":null,"terminal` +
			`":null,"platform":"","station":{"code":"s9600216","station_type":"аэропорт","title":"Домодедово","po` +
			`pular_title":"","short_title":"","codes":{"yandex":"s9600216","esr":"193114"},"transport_type":"plan` +
			`e","type":"station"},"stop_time":null,"duration":8400.0}],"vehicle":"Boeing 777-200","start_date":"2` +
			`017-05-19","transport_subtype":{"color":null,"code":null,"title":null},"express_type":null}`,
	)))
	return
}

func TestThread(t *testing.T) {
	api := NewYandexRapsApi("TEST_APIKEY")
	api.Fetcher = &ThreadDataFetcher{}
	reps, err := api.Thread(map[string]string{
		"uid":          "NN-324_0_c9_12",
		"lang":         "ru_RU",
		"show_systems": "all",
	})
	if err != nil {
		t.Error(err)
	}

	data := ThreadCard{
		ExceptDays:    "",
		ArrivalDate:   "",
		From:          "",
		Uid:           "NN-324_0_c9_12",
		Title:         "Симферополь — Москва",
		DepartureDate: "",
		StartTime:     "01:10",
		Number:        "NN 324",
		ShortTitle:    "Симферополь — Москва",
		Days:          "19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31 мая, 1, 2, 3, 6, 9, 10, 13 июня, …",
		To:            "",
		Carrier: ThreadCarrier{
			Code:    9,
			Offices: []string{},
			Codes: Codes{
				Icao:   "",
				Sirena: "НН",
				Iata:   "NN",
			},
			Title: "ВИМ-Авиа",
		},
		TransportType: "plane",
		Stops: []Stop{{
			Arrival:   "",
			Departure: "2017-05-19 01:10:00",
			Terminal:  "",
			Platform:  "",
			Station: ThreadStation{
				Codes: StationCodes{
					Yandex: "s9600396",
				},
				Title:        "Симферополь",
				PopularTitle: "",
				ShortTitle:   "",
				Code:         "s9600396",
				Type:         "station",
			},
			StopTime: "",
			Duration: 0.0,
		}, {
			Arrival:   "2017-05-19 03:30:00",
			Departure: "",
			Terminal:  "",
			Platform:  "",
			Station: ThreadStation{
				Codes: StationCodes{
					Yandex: "s9600216",
					Esr:    "193114",
				},
				Title:        "Домодедово",
				PopularTitle: "",
				ShortTitle:   "",
				Code:         "s9600216",
				Type:         "station",
			},
			StopTime: "",
			Duration: 8400.0,
		}},
		Vehicle:   "Boeing 777-200",
		StartDate: "2017-05-19",
		TransportSubtype: TransportSubtype{
			Color: "",
			Code:  "",
			Title: "",
		},
		ExpressType: "",
	}

	assert.Equal(t, reps, data, "Thread is not Equal!")
}
