Astrolabe
=========

Martini middleware/handler for easily finding all routes of your app.


## Why?

I started playing with Go a couple of days ago, so I figured this would be a good chance to create a package and share it with the world, I love how easy [Martini](https://github.com/go-martini/martini) makes web development in Go, so creating a middleware was naturally the first thought I had.

Fortunately for me, the functionality provided by Astrolabe was [asked for](https://github.com/go-martini/martini/issues/227) a while ago, so I didn't actually have to do much of work. It was a great exercise though.

## What's in a name?

I'm not really great at names, but my inspiration was the Rails' engine [Sextant](https://github.com/schneems/sextant).

## How to use

Astrolabe have only been tested with [martini.Classic()](http://godoc.org/github.com/go-martini/martini#Classic), if you use [martini.New()](http://godoc.org/github.com/go-martini/martini#New) and face any issues, please report the problem.

The route `/martini/routes` is only accessible in development. You can use the middleware in the following way:

```
// server.go
package main

import (
	"github.com/ahazem/astrolabe"
	"github.com/go-martini/martini"
)

func main() {
  m := martini.Classic()
  
  // Add some routes.
  m.Get("/posts", func() string {
    return "Hello World!"
  })
  
  // Use astrolabe (router is used to expose /martini/routes endpoint)
  m.Use(astrolabe.ExposeEndpoint(m.Router))

  m.Run()
}
```

Visit `http://localhost:3000/martini/routes` to see a list of all the routes in your app.

## License
See [LICENSE](https://github.com/ahazem/astrolabe/blob/master/LICENSE).

#### Disclaimer

Copyrights of some parts of the HTML code used belong to [Jeremy Saenz](https://github.com/codegangsta), and is licensed under the MIT license.