package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
)

func newTestRequest(method, path, body string) *httptest.ResponseRecorder {
	var request *http.Request
	var err error
	switch method {
	case "GET":
		request, err = http.NewRequest("GET", path, nil)
		if err != nil {
			log.Fatal(err)
		}
	case "POST":
		request, err = http.NewRequest("POST", path, strings.NewReader(body))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("unsupported method: %s", method)
	}
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	return recorder
}
