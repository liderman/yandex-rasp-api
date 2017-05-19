package yandex_rasp_api

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"testing"
)

type NearestSettlementDataFetcher struct{}

func (n *NearestSettlementDataFetcher) Fetch(url string) (resp io.ReadCloser, err error) {
	err = nil
	resp = ioutil.NopCloser(bytes.NewReader([]byte(
		`{"distance":25.723805024479688,"code":"c10680","title":"Павловск","popular_title":"Павловск","short_` +
			`title":"Павловск","lat":50.44973,"lng":40.125282,"type":"settlement"}`,
	)))
	return
}

func TestNearestSettlement(t *testing.T) {
	api := NewYandexRapsApi("TEST_TOKEN")
	api.Fetcher = &NearestSettlementDataFetcher{}
	reps, err := api.NearestSettlement(map[string]string{
		"lat":      "50.440046",
		"lng":      "40.4882367",
		"distance": "50",
		"lang":     "ru_RU",
	})
	if err != nil {
		t.Error(err)
	}

	data := NearestSettlement{
		Distance:     25.723805024479688,
		Code:         "c10680",
		Title:        "Павловск",
		PopularTitle: "Павловск",
		ShortTitle:   "Павловск",
		Lat:          50.44973,
		Lng:          40.125282,
		Type:         "settlement",
	}
	assert.Equal(t, reps, data, "NearestSettlement is not Equal!")
}
