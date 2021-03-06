package yandex_rasp_api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type YandexRapsApi struct {
	apikey  string
	log     LoggerInterface
	Fetcher DataFetcher
}

type LoggerInterface interface {
	Debug(...interface{})
}

type DataFetcher interface {
	Fetch(url string) (resp io.ReadCloser, err error)
}

type HttpDataFetcher struct{}

func (d *HttpDataFetcher) Fetch(url string) (resp io.ReadCloser, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	client := &http.Client{}
	httpResp, err := client.Do(req)
	if err != nil {
		return
	}

	resp = httpResp.Body
	return
}

// NewYandexRapsApi creates a new instance YandexRapsApi.
func NewYandexRapsApi(apikey string) *YandexRapsApi {
	return &YandexRapsApi{
		apikey:  apikey,
		Fetcher: &HttpDataFetcher{},
	}
}

func (y *YandexRapsApi) SetLogger(logger LoggerInterface) {
	y.log = logger
}

func (y *YandexRapsApi) getJson(path string, args map[string]string, v interface{}) error {
	apiUrl, err := url.Parse("https://api.rasp.yandex.net/" + path)
	if err != nil {
		return err
	}
	params := url.Values{}
	params.Add("format", "json")
	params.Add("apikey", y.apikey)
	for k, v := range args {
		if v == "" {
			continue
		}
		params.Add(k, v)
	}

	apiUrl.RawQuery = params.Encode()

	if y.log != nil {
		y.log.Debug("API Send: " + apiUrl.String())
	}

	resp, err := y.Fetcher.Fetch(apiUrl.String())
	if err != nil {
		return err
	}
	defer resp.Close()
	return json.NewDecoder(resp).Decode(&v)
}
