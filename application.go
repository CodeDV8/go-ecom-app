package EComApp

import (
	"github.com/codedv8/go-ecom-app/urihandler"
)

func NewApplication() *Application {
	app := &Application{
		Hooks:      make(map[string][]Hook),
		URIHandler: &urihandler.URIHandler{},
	}
	app.URIHandler.Init()
	return app
}
