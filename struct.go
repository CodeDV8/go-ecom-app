package EComApp

import (
	"github.com/codedv8/go-ecom-app/urihandler"
	EComDB "github.com/codedv8/go-ecom-db"
	"github.com/gin-gonic/gin"
	"plugin"
)

type Application struct {
	DB            *EComDB.DBConnector
	Hooks         map[string][]Hook
	SystemModules []Module
	UserModules   []Module
	Router        *gin.Engine
	URIHandler    *urihandler.URIHandler
}

type HookCallback func(*func(interface{}) (bool, error)) (bool, error)

type Hook struct {
	Callback func(interface{}) (bool, error)
}

func (app *Application) ListenToHook(name string, callback func(interface{}) (bool, error)) {
	hook := &Hook{
		Callback: callback,
	}
	app.Hooks[name] = append(app.Hooks[name], *hook)
}

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

type Module struct {
	SysInit  func(app *Application) error
	Init     func(app *Application) error
	Done     func(app *Application) error
	Plugin   *plugin.Plugin
	Filename string
}
