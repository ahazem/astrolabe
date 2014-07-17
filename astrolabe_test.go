package astrolabe

import (
	"net/http"
	"net/http/httptest"

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
