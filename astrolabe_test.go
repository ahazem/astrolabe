package astrolabe

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-martini/martini"
)

// This function acts as a setup for other tests.
func setup() *httptest.ResponseRecorder {
	m := martini.Classic()
	m.Use(ExposeEndpoint(m.Router))

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/martini/routes", nil)

	m.ServeHTTP(res, req)

	return res
}

// Test that /martini/routes endpoint is accessible and the content-type is set to text/html.
func TestExposeEndpoint(t *testing.T) {
	res := setup()

	if res.Code != 200 {
		t.Error("Response is not 200.")
	}

	if res.HeaderMap.Get("Content-Type") != "text/html" {
		t.Error("Expecting the Content-Type to be text/html.")
	}
}
