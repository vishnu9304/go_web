package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newMockApplication() *application {
	return &application{
		infoLog:  log.New(ioutil.Discard, "", 0),
		errorLog: log.New(ioutil.Discard, "", 0),
	}
}

func newTestHTTPServer(a *application) *httptest.Server {
	ts := httptest.NewTLSServer(a.loadRoutes())
	return ts
}

func mockGet(t *testing.T, ts *httptest.Server, url string) (status int, headers http.Header, body []byte) {
	rs, err := ts.Client().Get(ts.URL + url)
	status = rs.StatusCode
	headers = rs.Header

	defer rs.Body.Close()
	body, err = ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	return
}

func TestPing(t *testing.T) {
	// rr is an http.ResponseWriter object that records the response returned by the handler
	rr := httptest.NewRecorder()

	// create a request object
	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	a := application{}
	a.ping(rr, req)

	response := rr.Result()

	if response.StatusCode != http.StatusOK {
		t.Log(response.StatusCode)
		t.Errorf("non ok status code %d received", response.StatusCode)
	}

	t.Log(response.StatusCode)
}

// End 2 End testing example
func TestE2E(t *testing.T) {
	a := newMockApplication()

	ts := newTestHTTPServer(a)
	defer ts.Close()

	status, _, _ := mockGet(t, ts, "/ping")
	t.Log(status)
}
