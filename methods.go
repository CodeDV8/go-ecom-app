package EComApp

import (
	"github.com/gin-gonic/gin"
	"log"
)

func (app *Application) Init() {

	app.Router = gin.Default()
	// Load system plugins
	app.LoadPlugins("./system/", &app.SystemModules)
	// Load user plugins
	app.LoadPlugins("./user/", &app.UserModules)
	// Initialize all
	for _, module := range app.SystemModules {
		module.Init(app)
	}
	for _, module := range app.UserModules {
		module.Init(app)
	}
}

func (app *Application) Run() {
	log.Printf("Run: %+v\n", app)
}

func (app *Application) Done() {
	for _, module := range app.UserModules {
		module.Done(app)
	}
	for _, module := range app.SystemModules {
		module.Done(app)
	}
}
