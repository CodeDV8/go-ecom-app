package EComApp

import (
	EComBase "github.com/codedv8/go-ecom-base"
	"github.com/gin-gonic/gin"
	"plugin"
)

type Application struct {
	DB            interface{}
	Hooks         map[string][]Hook
	SystemModules []Module
	UserModules   []Module
	Router        *gin.Engine
}

type Hook struct {
	Callback EComBase.HookCallback
}

func (app *Application) AddHook(name string, callback func(interface{}) (bool, error)) {
	hook := &Hook{
		Callback: callback,
	}
	app.Hooks[name] = append(app.Hooks[name], *hook)
}

func (app *Application) CallHook(name string, args *interface{}) (bool, error) {
	for _, hook := range app.Hooks[name] {
		ok, err := hook.Callback(args)
		if err != nil {
			return ok, err
		}
		if ok == false {
			return false, nil
		}
	}
	return true, nil
}

// type Plugin struct {
// }

type Module struct {
	Init     func(app *Application) error
	Done     func(app *Application) error
	Plugin   *plugin.Plugin
	Filename string
}
