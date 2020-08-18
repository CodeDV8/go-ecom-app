package EComApp

import (
	"github.com/gin-gonic/gin"
	"plugin"
)

type Application struct {
	DB            interface{}
	Hooks         []Hook
	SystemModules []Module
	UserModules   []Module
	Router        *gin.Engine
}

type Hook struct {
}

// type Plugin struct {
// }

type Module struct {
	Init     func(app *Application) error
	Done     func(app *Application) error
	Plugin   *plugin.Plugin
	Filename string
}
