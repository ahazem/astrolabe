package astrolabe

import (
	"net/http"

	"github.com/go-martini/martini"
)

func exposeEndpoint(w http.ResponseWriter, r *http.Request, c martini.Context, routes martini.Routes) {
	w.Header().Set("Content-Type", "text/html")
}

func ExposeEndpoint(r martini.Router) martini.Handler {
	r.Get("/martini/routes", exposeEndpoint)

	return exposeEndpoint
}
