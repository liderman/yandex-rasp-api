package yandex_rasp_api

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"testing"
)

type CopyrightDataFetcher struct{}

func (c *CopyrightDataFetcher) Fetch(url string) (resp io.ReadCloser, err error) {
	err = nil
	resp = ioutil.NopCloser(bytes.NewReader([]byte(
		`{"copyright":{"logo_vm":"<iframe frameborder=\"0\" style=\"overflow: hidden; border: 0; width: 240px` +
			`; height: 130px;\" src=\"//yandex.st/rasp/media/apicc/copyright_vert_mono.html\"></iframe>","url":"h` +
			`ttp://rasp.yandex.ru/","logo_vd":"<iframe frameborder=\"0\" style=\"overflow: hidden; border: 0; wid` +
			`th: 240px; height: 130px;\" src=\"//yandex.st/rasp/media/apicc/copyright_vert_dark.html\"></iframe>"` +
			`,"logo_hy":"<iframe frameborder=\"0\" style=\"overflow: hidden; border: 0; width: 740px; height: 51p` +
			`x;\" src=\"//yandex.st/rasp/media/apicc/copyright_horiz_yellow.html\"></iframe>","logo_hd":"<iframe ` +
			`frameborder=\"0\" style=\"overflow: hidden; border: 0; width: 740px; height: 51px;\" src=\"//yandex.` +
			`st/rasp/media/apicc/copyright_horiz_dark.html\"></iframe>","logo_vy":"<iframe frameborder=\"0\" styl` +
			`e=\"overflow: hidden; border: 0; width: 240px; height: 130px;\" src=\"//yandex.st/rasp/media/apicc/c` +
			`opyright_vert_yellow.html\"></iframe>","text":"Данные предоставлены сервисом Яндекс.Расписания","log` +
			`o_hm":"<iframe frameborder=\"0\" style=\"overflow: hidden; border: 0; width: 740px; height: 51px;\" ` +
			`src=\"//yandex.st/rasp/media/apicc/copyright_horiz_mono.html\"></iframe>"}}`,
	)))
	return
}

func TestCopyright(t *testing.T) {
	api := NewYandexRapsApi("TEST_TOKEN")
	api.Fetcher = &CopyrightDataFetcher{}
	reps, err := api.Copyright()
	if err != nil {
		t.Error(err)
	}

	data := CopyrightResponse{
		Copyright: Copyright{
			Url:    "http://rasp.yandex.ru/",
			LogoVm: `<iframe frameborder="0" style="overflow: hidden; border: 0; width: 240px; height: 130px;" src="//yandex.st/rasp/media/apicc/copyright_vert_mono.html"></iframe>`,
			LogoVd: `<iframe frameborder="0" style="overflow: hidden; border: 0; width: 240px; height: 130px;" src="//yandex.st/rasp/media/apicc/copyright_vert_dark.html"></iframe>`,
			LogoHy: `<iframe frameborder="0" style="overflow: hidden; border: 0; width: 740px; height: 51px;" src="//yandex.st/rasp/media/apicc/copyright_horiz_yellow.html"></iframe>`,
			LogoHd: `<iframe frameborder="0" style="overflow: hidden; border: 0; width: 740px; height: 51px;" src="//yandex.st/rasp/media/apicc/copyright_horiz_dark.html"></iframe>`,
			LogoVy: `<iframe frameborder="0" style="overflow: hidden; border: 0; width: 240px; height: 130px;" src="//yandex.st/rasp/media/apicc/copyright_vert_yellow.html"></iframe>`,
			LogoHm: `<iframe frameborder="0" style="overflow: hidden; border: 0; width: 740px; height: 51px;" src="//yandex.st/rasp/media/apicc/copyright_horiz_mono.html"></iframe>`,
			Text:   "Данные предоставлены сервисом Яндекс.Расписания",
		},
	}

	assert.Equal(t, reps, data, "CopyrightResponse is not Equal!")
}
