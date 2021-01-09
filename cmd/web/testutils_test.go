package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestHTTPServer(h http.Handler) *httptest.Server {
	return httptest.NewTLSServer(h)
}

func newTestApplication() *application {
	a := &application{
		infoLog:  log.New(ioutil.Discard, "", 0),
		errorLog: log.New(ioutil.Discard, "", 0),
	}
	return a
}

func testHTTPGet(t *testing.T, url string, ts *httptest.Server) (statusCode int, body []byte, headers http.Header) {
	resp, err := ts.Client().Get(ts.URL + url)
	if err != nil {
		t.Fatal(err)
	}

	statusCode = resp.StatusCode

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	headers = resp.Header

	return statusCode, body, headers
}
