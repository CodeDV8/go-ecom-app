package ecomapp

import (
	"path/filepath"
	"plugin"
)

// LoadPlugins - Load all plugins from the given path and add them to the supplied list of modules
func (app *Application) LoadPlugins(path string, modules *[]Module) {
	allPlugins, err := filepath.Glob(path + "*.so")
	if err != nil {
		panic(err)
	}
	for _, filename := range allPlugins {
		p, err := plugin.Open(filename)
		if err != nil {
			continue
		}

		// Find SysInit
		symbol, err := p.Lookup("SysInit")
		if err != nil {
			continue
		}

		sysInitFunc, okSysInit := symbol.(func(app *Application) error)
		if !okSysInit {
			continue
		}

		// Find Init
		symbol, err = p.Lookup("Init")
		if err != nil {
			continue
		}

		initFunc, okInit := symbol.(func(app *Application) error)
		if !okInit {
			continue
		}

		// Find Done
		symbol, err = p.Lookup("Done")
		if err != nil {
			continue
		}

		doneFunc, okDone := symbol.(func(app *Application) error)
		if !okDone {
			continue
		}

		module := &Module{
			SysInit:  sysInitFunc,
			Init:     initFunc,
			Done:     doneFunc,
			Plugin:   p,
			Filename: filename,
		}
		*modules = append(*modules, *module)
	}
}
