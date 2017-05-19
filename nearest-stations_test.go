package yandex_rasp_api

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"testing"
)

type NearestStationsDataFetcher struct{}

func (n *NearestStationsDataFetcher) Fetch(url string) (resp io.ReadCloser, err error) {
	err = nil
	resp = ioutil.NopCloser(bytes.NewReader([]byte(
		`{"pagination":{"total":35,"limit":2,"offset":0},"stations":[{"distance":7.463483472582138,"code":"s9` +
			`761931","station_type":"автобусная остановка","type_choices":{"schedule":{"desktop_url":"https://ras` +
			`p.yandex.ru/station/9761931/schedule","touch_url":"https://t.rasp.yandex.ru/station/9761931/schedule` +
			`"}},"title":"Семеновка, Калачеевский район","popular_title":"","short_title":"","majority":4,"transp` +
			`ort_type":"bus","lat":50.4083887285445,"lng":40.5811027251184,"type":"station"},{"distance":9.855640` +
			`047221277,"code":"s9761930","station_type":"автобусная остановка","type_choices":{"schedule":{"deskt` +
			`op_url":"https://rasp.yandex.ru/station/9761930/schedule","touch_url":"https://t.rasp.yandex.ru/stat` +
			`ion/9761930/schedule"}},"title":"Гаврильск","popular_title":null,"short_title":null,"majority":4,"tr` +
			`ansport_type":"bus","lat":50.4314878059252,"lng":40.3497710824011,"type":"station"}]}`,
	)))
	return
}

func TestNearestStations(t *testing.T) {
	api := NewYandexRapsApi("TEST_TOKEN")
	api.Fetcher = &NearestStationsDataFetcher{}
	reps, err := api.NearestStations(map[string]string{
		"lang":   "ru_RU",
		"code":   "TK",
		"system": "iata",
	})
	if err != nil {
		t.Error(err)
	}

	data := NearestStationsResponse{
		Pagination: Pagination{
			Total:  35,
			Limit:  2,
			Offset: 0,
		},
		Stations: []NearestStations{{
			Distance:    7.463483472582138,
			Code:        "s9761931",
			StationType: "автобусная остановка",
			TypeChoices: map[string]TypeChoiceLink{
				"schedule": {
					DesktopUrl: "https://rasp.yandex.ru/station/9761931/schedule",
					TouchUrl:   "https://t.rasp.yandex.ru/station/9761931/schedule",
				},
			},
			Title:         "Семеновка, Калачеевский район",
			PopularTitle:  "",
			ShortTitle:    "",
			Majority:      4,
			TransportType: "bus",
			Lat:           50.4083887285445,
			Lng:           40.5811027251184,
			Type:          "station",
		}, {
			Distance:    9.855640047221277,
			Code:        "s9761930",
			StationType: "автобусная остановка",
			TypeChoices: map[string]TypeChoiceLink{
				"schedule": {
					DesktopUrl: "https://rasp.yandex.ru/station/9761930/schedule",
					TouchUrl:   "https://t.rasp.yandex.ru/station/9761930/schedule",
				},
			},
			Title:         "Гаврильск",
			PopularTitle:  "",
			ShortTitle:    "",
			Majority:      4,
			TransportType: "bus",
			Lat:           50.4314878059252,
			Lng:           40.3497710824011,
			Type:          "station",
		}},
	}
	assert.Equal(t, reps, data, "NearestStationsResponse is not Equal!")
}
