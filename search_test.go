package yandex_rasp_api

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"testing"
)

type SearchDataFetcher struct{}

func (c *SearchDataFetcher) Fetch(url string) (resp io.ReadCloser, err error) {
	err = nil
	resp = ioutil.NopCloser(bytes.NewReader([]byte(
		`{"pagination":{"total":31,"limit":2,"offset":0},"segments":[{"arrival":"2017-05-18 03:30:00","from":` +
			`{"code":"s9600396","title":"Симферополь","station_type":"аэропорт","popular_title":"","short_title":` +
			`"","transport_type":"plane","type":"station"},"thread":{"uid":"NN-324_0_c9_12","title":"Симферополь ` +
			`— Москва","number":"NN 324","short_title":"Симферополь — Москва","thread_method_link":"api.rasp.yand` +
			`ex.net/v3/thread/?date=2017-05-18&uid=NN-324_0_c9_12","carrier":{"code":9,"contacts":"Телефон: +7 (4` +
			`95) 7830088","url":"http://www.vim-avia.com/","logo_svg":"//yastatic.net/rasp/media/data/company/svg` +
			`/vim_airlines.svg","title":"ВИМ-Авиа","phone":"","codes":{"icao":null,"sirena":"НН","iata":"NN"},"ad` +
			`dress":"г. Москва, ул. Новохохловская, д. 23, стр. 1","logo":"//yastatic.net/rasp/media/data/company` +
			`/logo/vim.jpg","email":""},"transport_type":"plane","vehicle":"Boeing 777-200","transport_subtype":{` +
			`"color":null,"code":null,"title":null},"express_type":null},"departure_platform":"","departure":"201` +
			`7-05-18 01:10:00","stops":"","departure_terminal":null,"to":{"code":"s9600216","title":"Домодедово",` +
			`"station_type":"аэропорт","popular_title":"","short_title":"","transport_type":"plane","type":"stati` +
			`on"},"has_transfers":false,"tickets_info":null,"duration":8400.0,"arrival_terminal":null,"start_date` +
			`":"2017-05-18","arrival_platform":""},{"arrival":"2017-05-18 05:00:00","from":{"code":"s9600396","ti` +
			`tle":"Симферополь","station_type":"аэропорт","popular_title":"","short_title":"","transport_type":"p` +
			`lane","type":"station"},"thread":{"uid":"SU-1627_0_c26_547","title":"Симферополь — Москва","number":` +
			`"SU 1627","short_title":"Симферополь — Москва","thread_method_link":"api.rasp.yandex.net/v3/thread/?` +
			`date=2017-05-18&uid=SU-1627_0_c26_547","carrier":{"code":26,"contacts":"Центр информации и бронирова` +
			`ния: +7 (495) 223-55-55,  +7 (800) 444-55-55. <br>\r\ne-mail: callcenter@aeroflot.ru","url":"http://` +
			`www.aeroflot.ru/","logo_svg":"//yastatic.net/rasp/media/data/company/svg/Aeroflot_1.svg","title":"Аэ` +
			`рофлот","phone":"","codes":{"icao":"AFL","sirena":"СУ","iata":"SU"},"address":"Москва, Ленинградский` +
			` пр., д.37, корп.9 ","logo":"//yastatic.net/rasp/media/data/company/logo/aeroflot_logo_ru.gif","emai` +
			`l":""},"transport_type":"plane","vehicle":"Airbus A320","transport_subtype":{"color":null,"code":nul` +
			`l,"title":null},"express_type":null},"departure_platform":"","departure":"2017-05-18 02:25:00","stop` +
			`s":"","departure_terminal":null,"to":{"code":"s9600213","title":"Шереметьево","station_type":"аэропо` +
			`рт","popular_title":"","short_title":"","transport_type":"plane","type":"station"},"has_transfers":f` +
			`alse,"tickets_info":null,"duration":9300.0,"arrival_terminal":null,"start_date":"2017-05-18","arriva` +
			`l_platform":""}],"search":{"date":"2017-05-18","to":{"code":"c213","type":"settlement","popular_titl` +
			`e":"Москва","short_title":"Москва","title":"Москва"},"from":{"code":"c146","type":"settlement","popu` +
			`lar_title":"Симферополь","short_title":"Симферополь","title":"Симферополь"}}}`,
	)))
	return
}

func TestSearch(t *testing.T) {
	api := NewYandexRapsApi("TEST_TOKEN")
	api.Fetcher = &SearchDataFetcher{}
	reps, err := api.Search(map[string]string{
		"from":  "c146",
		"to":    "c213",
		"lang":  "ru_RU",
		"page":  "1",
		"date":  "2017-05-18",
		"limit": "2",
	})
	if err != nil {
		t.Error(err)
	}

	data := SearchResponse{
		Pagination: Pagination{
			Total:  31,
			Limit:  2,
			Offset: 0,
		},
		Segments: []Segment{{
			Arrival: "2017-05-18 03:30:00",
			From: Station{
				Code:          "s9600396",
				StationType:   "аэропорт",
				Title:         "Симферополь",
				PopularTitle:  "",
				ShortTitle:    "",
				TransportType: "plane",
				Type:          "station",
			},
			Thread: Thread{
				Carrier: SearchCarrier{
					Title:    "ВИМ-Авиа",
					Code:     9,
					Contacts: "Телефон: +7 (495) 7830088",
					Url:      "http://www.vim-avia.com/",
					Phone:    "",
					Codes: Codes{
						Icao:   "",
						Sirena: "НН",
						Iata:   "NN",
					},
					Address: "г. Москва, ул. Новохохловская, д. 23, стр. 1",
					Logo:    "//yastatic.net/rasp/media/data/company/logo/vim.jpg",
					LogoSvg: "//yastatic.net/rasp/media/data/company/svg/vim_airlines.svg",
					Email:   "",
				},
				TransportType:    "plane",
				Uid:              "NN-324_0_c9_12",
				Title:            "Симферополь — Москва",
				Vehicle:          "Boeing 777-200",
				Number:           "NN 324",
				ShortTitle:       "Симферополь — Москва",
				ThreadMethodLink: "api.rasp.yandex.net/v3/thread/?date=2017-05-18&uid=NN-324_0_c9_12",
				ExpressType:      "",
				TransportSubtype: map[string]string{
					"color": "",
					"code":  "",
					"title": "",
				},
			},
			DeparturePlatform: "",
			Departure:         "2017-05-18 01:10:00",
			Stops:             "",
			DepartureTerminal: "",
			To: Station{
				Code:          "s9600216",
				StationType:   "аэропорт",
				Title:         "Домодедово",
				PopularTitle:  "",
				ShortTitle:    "",
				TransportType: "plane",
				Type:          "station",
			},
			HasTransfers:    false,
			TicketsInfo:     TicketsInfo{},
			Duration:        8400.0,
			ArrivalTerminal: "",
			StartDate:       "2017-05-18",
			ArrivalPlatform: "",
		}, {
			Arrival: "2017-05-18 05:00:00",
			From: Station{
				Code:          "s9600396",
				StationType:   "аэропорт",
				Title:         "Симферополь",
				PopularTitle:  "",
				ShortTitle:    "",
				TransportType: "plane",
				Type:          "station",
			},
			Thread: Thread{
				Carrier: SearchCarrier{
					Code:     26,
					Contacts: "Центр информации и бронирования: +7 (495) 223-55-55,  +7 (800) 444-55-55. <br>\r\ne-mail: callcenter@aeroflot.ru",
					Url:      "http://www.aeroflot.ru/",
					Title:    "Аэрофлот",
					Phone:    "",
					Codes: Codes{
						Icao:   "AFL",
						Sirena: "СУ",
						Iata:   "SU",
					},
					Address: "Москва, Ленинградский пр., д.37, корп.9 ",
					Logo:    "//yastatic.net/rasp/media/data/company/logo/aeroflot_logo_ru.gif",
					LogoSvg: "//yastatic.net/rasp/media/data/company/svg/Aeroflot_1.svg",
					Email:   "",
				},
				TransportType:    "plane",
				Uid:              "SU-1627_0_c26_547",
				Title:            "Симферополь — Москва",
				Vehicle:          "Airbus A320",
				Number:           "SU 1627",
				ShortTitle:       "Симферополь — Москва",
				ThreadMethodLink: "api.rasp.yandex.net/v3/thread/?date=2017-05-18&uid=SU-1627_0_c26_547",
				ExpressType:      "",
				TransportSubtype: map[string]string{
					"color": "",
					"code":  "",
					"title": "",
				},
			},
			DeparturePlatform: "",
			Departure:         "2017-05-18 02:25:00",
			Stops:             "",
			DepartureTerminal: "",
			To: Station{
				Code:          "s9600213",
				StationType:   "аэропорт",
				Title:         "Шереметьево",
				PopularTitle:  "",
				ShortTitle:    "",
				TransportType: "plane",
				Type:          "station",
			},
			HasTransfers:    false,
			TicketsInfo:     TicketsInfo{},
			Duration:        9300.0,
			ArrivalTerminal: "",
			StartDate:       "2017-05-18",
			ArrivalPlatform: "",
		}},
		Search: Search{
			Date: "2017-05-18",
			To: Location{
				Code:         "c213",
				Type:         "settlement",
				PopularTitle: "Москва",
				ShortTitle:   "Москва",
				Title:        "Москва",
			},
			From: Location{
				Code:         "c146",
				Type:         "settlement",
				PopularTitle: "Симферополь",
				ShortTitle:   "Симферополь",
				Title:        "Симферополь",
			},
		},
	}

	assert.Equal(t, reps, data, "SearchResponse is not Equal!")
}
