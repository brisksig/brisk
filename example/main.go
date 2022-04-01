package main

import (
	"brisk"

	"github.com/DomineCore/brisk/example/app01"
)

func LoadRouter(app *brisk.Brisk) {
	app.Router.Include("/app01/", app01.Router())
}

func main() {
	app := brisk.New()
	// load router
	LoadRouter(app)
	// load MiddleWare
	app.Router.Use(&brisk.LoggingMiddleware{})
	app.Router.Use(&brisk.CrosMiddleware{})
	// run
	app.Run(":8000")
}
