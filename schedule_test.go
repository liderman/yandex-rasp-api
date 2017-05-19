package yandex_rasp_api

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"testing"
)

type ScheduleDataFetcher struct{}

func (s *ScheduleDataFetcher) Fetch(url string) (resp io.ReadCloser, err error) {
	err = nil
	resp = ioutil.NopCloser(bytes.NewReader([]byte(
		`{"pagination":{"total":43,"limit":2,"offset":0},"schedule":[{"except_days":"16, 17, 18, 21, 22, 23, ` +
			`24, 25, 28, 29, 30, 31 мая, 1 июня","direction":"на Москву","thread":{"uid":"7273_0_9600213_g17_4","` +
			`title":"аэропорт Шереметьево — Москва (Белорусский вокзал)","number":"7273","short_title":"а/п Шерем` +
			`етьево — М-Белорусск.","carrier":{"code":162,"codes":{"icao":null,"sirena":null,"iata":null},"title"` +
			`:"Аэроэкспресс"},"transport_type":"suburban","vehicle":null,"transport_subtype":{"color":"#FF7F44","` +
			`code":"suburban","title":"Пригородный поезд"},"express_type":"aeroexpress"},"arrival_time":null,"is_` +
			`fuzzy":false,"days":"ежедневно","stops":"без остановок","terminal":null,"platform":"","departure_tim` +
			`e":"00:01"},{"except_days":"16, 17, 18, 21, 22, 23, 24, 25, 28, 29, 30, 31 мая, 1 июня","direction":` +
			`"на Москву","thread":{"uid":"7275_0_9600213_g17_4","title":"аэропорт Шереметьево — Москва (Белорусск` +
			`ий вокзал)","number":"7275","short_title":"а/п Шереметьево — М-Белорусск.","carrier":{"code":162,"co` +
			`des":{"icao":null,"sirena":null,"iata":null},"title":"Аэроэкспресс"},"transport_type":"suburban","ve` +
			`hicle":null,"transport_subtype":{"color":"#FF7F44","code":"suburban","title":"Пригородный поезд"},"e` +
			`xpress_type":"aeroexpress"},"arrival_time":null,"is_fuzzy":false,"days":"ежедневно","stops":"без ост` +
			`ановок","terminal":null,"platform":"","departure_time":"00:30"}],"schedule_direction":{"code":"на Мо` +
			`скву","title":"на Москву"},"date":null,"station":{"code":"s9600213","title":"Шереметьево","station_t` +
			`ype":"аэропорт","popular_title":"","short_title":"","transport_type":"plane","type":"station"},"dire` +
			`ctions":[{"code":"arrival","title":"прибытие"},{"code":"на Москву","title":"на Москву"},{"code":"all` +
			`","title":"все направления"}]}`,
	)))
	return
}

func TestSchedule(t *testing.T) {
	api := NewYandexRapsApi("TEST_APIKEY")
	api.Fetcher = &ScheduleDataFetcher{}
	reps, err := api.Schedule(map[string]string{
		"station":         "s9600213",
		"transport_types": "suburban",
		"direction":       "на Москву",
		"limit":           "2",
	})
	if err != nil {
		t.Error(err)
	}

	data := ScheduleResponse{
		Pagination: Pagination{
			Total:  43,
			Limit:  2,
			Offset: 0,
		},
		Schedule: []Schedule{
			{
				ExceptDays: "16, 17, 18, 21, 22, 23, 24, 25, 28, 29, 30, 31 мая, 1 июня",
				Direction:  "на Москву",
				Thread: ScheduleThread{
					Carrier: ShortCarrier{
						Code: 162,
						Codes: Codes{
							Icao:   "",
							Sirena: "",
							Iata:   "",
						},
						Title: "Аэроэкспресс",
					},
					TransportType: "suburban",
					Uid:           "7273_0_9600213_g17_4",
					Title:         "аэропорт Шереметьево — Москва (Белорусский вокзал)",
					Vehicle:       "",
					Number:        "7273",
					ShortTitle:    "а/п Шереметьево — М-Белорусск.",
					TransportSubtype: TransportSubtype{
						Color: "#FF7F44",
						Code:  "suburban",
						Title: "Пригородный поезд",
					},
					ExpressType: "aeroexpress",
				},
				ArrivalTime:   "",
				Platform:      "",
				Days:          "ежедневно",
				Stops:         "без остановок",
				DepartureTime: "00:01",
				Terminal:      "",
				IsFuzzy:       false,
			},
			{
				ExceptDays: "16, 17, 18, 21, 22, 23, 24, 25, 28, 29, 30, 31 мая, 1 июня",
				Direction:  "на Москву",
				Thread: ScheduleThread{
					Carrier: ShortCarrier{
						Code: 162,
						Codes: Codes{
							Icao:   "",
							Sirena: "",
							Iata:   "",
						},
						Title: "Аэроэкспресс",
					},
					TransportType: "suburban",
					Uid:           "7275_0_9600213_g17_4",
					Title:         "аэропорт Шереметьево — Москва (Белорусский вокзал)",
					Vehicle:       "",
					Number:        "7275",
					ShortTitle:    "а/п Шереметьево — М-Белорусск.",
					TransportSubtype: TransportSubtype{
						Color: "#FF7F44",
						Code:  "suburban",
						Title: "Пригородный поезд",
					},
					ExpressType: "aeroexpress",
				},
				ArrivalTime:   "",
				Platform:      "",
				Days:          "ежедневно",
				Stops:         "без остановок",
				DepartureTime: "00:30",
				Terminal:      "",
				IsFuzzy:       false,
			},
		},
		ScheduleDirection: Direction{
			Code:  "на Москву",
			Title: "на Москву",
		},
		Directions: []Direction{{
			Code:  "arrival",
			Title: "прибытие",
		}, {
			Code:  "на Москву",
			Title: "на Москву",
		}, {
			Code:  "all",
			Title: "все направления",
		}},
		Station: Station{
			Code:          "s9600213",
			StationType:   "аэропорт",
			Title:         "Шереметьево",
			PopularTitle:  "",
			ShortTitle:    "",
			TransportType: "plane",
			Type:          "station",
		},
		Date: "",
	}

	assert.Equal(t, reps, data, "CopyrightResponse is not Equal!")
}
