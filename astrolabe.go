package astrolabe

import (
	"log"
	"net/http"

	"github.com/go-martini/martini"
)

func exposeEndpoint(w http.ResponseWriter, r *http.Request, c martini.Context, routes martini.Routes) {
	w.Header().Set("Content-Type", "text/html")
}

func ExposeEndpoint(r martini.Router) martini.Handler {
	if martini.Env == martini.Dev {
		r.Get("/martini/routes", exposeEndpoint)

		return exposeEndpoint
	} else {
		// If used in any other environment, do nothing (but log that).
		return func(res http.ResponseWriter, req *http.Request, log *log.Logger) {
			log.Println("The astrolabe middleware does not work in environments other than development.")
		}
	}
}
