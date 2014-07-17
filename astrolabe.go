package astrolabe

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-martini/martini"
)

type route struct {
	Pattern string
	Method  string
}

const routesHTML string = `<html>
  <head>
    <title>Routes</title>
    <style type="text/css">
			html, body {
				font-size: 14px;
				font-family: sans-serif;
				color: #333333;
				background-color: #ea5343;
				margin: 0px;
			}
			h1 {
				color: #d04526;
				background-color: #ffffff;
				padding: 20px;
				border-bottom: 1px dashed #2b3848;
			}
			div {
				margin: 20px;
				padding: 20px;
				border: 2px solid #2b3848;
				background-color: #ffffff;
			}
			table {
				font-size: 14px;
				width: 50%;
				margin: 0 auto;
				border-collapse: collapse;
				text-align: center;
			}
			thead {
				width: 75%;
			}
			thead > tr {
				width: 40%;
			}
			thead > tr > th {
				width: inherit;
				padding-bottom: 10px;
			}
			tr {
				text-align: left;
			}
			td {
				padding: 5px 0;
			}
    </style>
  </head>
  <body>
    <h1>Routes</h1>

		<div>
			<table>
				<thead>
					<tr>
						<th>HTTP Verb</th>
						<th>Path</th>
					</tr>
				</thead>

				<tbody>
				{{range .}}
					<tr>
						<td>{{.Method}}</td>
						<td>{{.Pattern}}</td>
					</tr>
				{{end}}
				</tbody>
			</table>
		</div>
  </body>
</html>`

func exposeEndpoint(w http.ResponseWriter, r *http.Request, c martini.Context, routes martini.Routes) {
	arr := make([]route, len(routes.All())-1)

	for i, e := range routes.All() {
		// Remove /martini/routes endpoint from the array passed to the template.
		if i != (len(routes.All()) - 1) {
			arr[i].Method = e.Method()
			arr[i].Pattern = e.Pattern()
		}
	}

	w.Header().Set("Content-Type", "text/html")

	t := template.New("routes")
	t.Parse(routesHTML)
	t.Execute(w, arr)
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
