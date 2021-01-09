package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
// if you tests are growing you can move the functions newMockApplication, newTestHTTPServer and mockGet to a separate file

func newMockApplication() *application {
	return &application{
		infoLog:  log.New(ioutil.Discard, "", 0),
		errorLog: log.New(ioutil.Discard, "", 0),
	}
}

func newTestHTTPServer(a *application) *httptest.Server {
	// NewTLSServer function takes an "http.Handler" object and returns an test server instance
	ts := httptest.NewTLSServer(a.loadRoutes())
	return ts
}

func mockGet(t *testing.T, ts *httptest.Server, url string) (status int, headers http.Header, body []byte) {

	// ts.Client().Get() method to make requests against our test server
	// we don't know under which port our test server is running. It picks up a random port.

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
*/

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
	a := newTestApplication()

	ts := newTestHTTPServer(a.loadRoutes())
	// Make sure you close the test server instance by end of the tests.
	defer ts.Close()

	statusCode, body, headers := testHTTPGet(t, "/ping", ts)
	t.Log(statusCode)
	t.Log(string(body))
	t.Log(headers)
}
