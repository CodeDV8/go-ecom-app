package EComApp

import (
	"plugin"

	"github.com/codedv8/go-ecom-app/urihandler"
	EComDB "github.com/codedv8/go-ecom-db"
	"github.com/gin-gonic/gin"
)

// Application - Struct that defines the application
type Application struct {
	DB            *EComDB.DBConnector
	Hooks         map[string][]Hook
	SystemModules []Module
	UserModules   []Module
	Router        *gin.Engine
	URIHandler    *urihandler.URIHandler
}

// HookCallback - type definition for a HookCallback
type HookCallback func(*func(interface{}) (bool, error)) (bool, error)

// Hook - Struct to define a hook
type Hook struct {
	Callback func(interface{}) (bool, error)
}

// ListenToHook - Register a hook listener by name and with its callback function
func (app *Application) ListenToHook(name string, callback func(interface{}) (bool, error)) {
	hook := &Hook{
		Callback: callback,
	}
	app.Hooks[name] = append(app.Hooks[name], *hook)
}

// CallHook - Call and handle hokks by the given name and execute them with the supplied payload
func (app *Application) CallHook(name string, payload interface{}) (bool, bool, error) {
	handled := false
	for _, hook := range app.Hooks[name] {
		next, err := hook.Callback(payload)
		if err != nil {
			return handled, next, err
		}
		handled = true
		if next == false {
			return handled, next, err
		}
	}
	return handled, true, nil
}

// type Plugin struct {
// }

// Module - Struct that defined a module (plugin)
type Module struct {
	SysInit  func(app *Application) error
	Init     func(app *Application) error
	Done     func(app *Application) error
	Plugin   *plugin.Plugin
	Filename string
}
