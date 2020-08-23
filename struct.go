package EComApp

import (
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

func (app *Application) CallHook(name string, payload interface{}) (bool, error) {
	for _, hook := range app.Hooks[name] {
		ok, err := hook.Callback(payload)
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
	SysInit  func(app *Application) error
	Init     func(app *Application) error
	Done     func(app *Application) error
	Plugin   *plugin.Plugin
	Filename string
}
