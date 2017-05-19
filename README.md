# yandex-rasp-api
Golang implementation Yandex Rasp Api v3.0 - https://tech.yandex.ru/rasp/doc/concepts/about-docpage/.

[![Build Status](https://travis-ci.org/liderman/yandex-rasp-api.svg?branch=master)](https://travis-ci.org/liderman/yandex-rasp-api)&nbsp;[![GoDoc](https://godoc.org/github.com/liderman/yandex-rasp-api?status.svg)](https://godoc.org/github.com/liderman/yandex-rasp-api)

Installation
-----------
	go get github.com/liderman/yandex-rasp-api

Usage
-----------
Getting a schedule of flights between stations:
```go
    yapi := NewYandexRapsApi("YOUR_APIKEY")
    reps, _ := yapi.Search(map[string]string{
        "from":  "c146",
        "to":    "c213",
        "lang":  "ru_RU",
        "page":  "1",
        "date":  "2017-05-18",
        "limit": "2",
    })
    fmt.Println(reps);
```

Getting information about the nearest town to the specified point:
```go
    yapi := NewYandexRapsApi("YOUR_APIKEY")
    reps, err := yapi.NearestSettlement(map[string]string{
        "lat":      "50.440046",
        "lng":      "40.4882367",
        "distance": "50",
        "lang":     "ru_RU",
    })
    fmt.Println(reps);
```

More examples can be found in the test files.

Features
--------

* Full support api version 3.0
* Code is covered by tests
* Without external dependencies

Requirements
-----------

* Need at least `go1.2` or newer.

Documentation
-----------

You can read package documentation [here](http:godoc.org/github.com/liderman/yandex-rasp-api).
Official docs: [here](https://tech.yandex.ru/rasp/doc/concepts/about-docpage/)

Testing
-----------
Unit-tests:
```bash
go test -v
```