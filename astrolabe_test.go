package astrolabe

import (
	"net/http"
	"net/http/httptest"
	"strings"
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

// This function is used as a setup but with a route defined.
func setupWithRoute(method string, handler martini.Handler) *httptest.ResponseRecorder {
	m := martini.Classic()

	switch method {
	case "GET":
		m.Get("/posts", handler)
	case "POST":
		m.Post("/posts", handler)
	case "PUT":
		m.Put("/posts/:id", handler)
	case "PATCH":
		m.Patch("/posts/:id", handler)
	case "OPTIONS":
		m.Options("/posts", handler)
	case "HEAD":
		m.Head("/posts", handler)
	case "DELETE":
		m.Delete("/posts/:id", handler)
	}

	m.Use(ExposeEndpoint(m.Router))

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/martini/routes", nil)

	m.ServeHTTP(res, req)

	return res
}

// This checks the endpoint is not exposed when the MARTINI.ENV is anything other than development.
func TestNotExposedIfNotInDevelopment(t *testing.T) {
	res := setup()

	// Check that we are not in development.
	if martini.Env != martini.Dev {
		if res.Code != http.StatusNotFound {
			t.Error("/martini/routes endpoint should not be exposed when environment is not development.")
		}
	}
}

// Test GET routes are displayed.
func TestGetRoute(t *testing.T) {
	res := setupWithRoute("GET", func(w http.ResponseWriter, r *http.Request) string {
		return "a GET route"
	})

	if !strings.Contains(res.Body.String(), "GET") {
		t.Error("The GET route should be displayed.")
	}

	if !strings.Contains(res.Body.String(), "/posts") {
		t.Error("The route path should be displayed.")
	}
}

// Test POST routes are displayed.
func TestPostRoute(t *testing.T) {
	res := setupWithRoute("POST", func(w http.ResponseWriter, r *http.Request) string {
		return "a POST route"
	})

	if !strings.Contains(res.Body.String(), "POST") {
		t.Error("The POST route should be displayed.")
	}

	if !strings.Contains(res.Body.String(), "/posts") {
		t.Error("The route path should be displayed.")
	}
}
