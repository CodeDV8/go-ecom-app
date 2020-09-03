package ecomapp

import (
	"github.com/codedv8/go-ecom-app/urihandler"
	ecomdb "github.com/codedv8/go-ecom-db"
)

// NewApplication - Create a new application object
func NewApplication() *Application {
	app := &Application{
		DB:         &ecomdb.DBConnector{},
		Hooks:      make(map[string][]Hook),
		URIHandler: &urihandler.URIHandler{},
	}
	app.URIHandler.Init()
	app.DB.Init()
	return app
}
