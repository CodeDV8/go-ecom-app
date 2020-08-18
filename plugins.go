package EComApp

import (
	"path/filepath"
	"plugin"
)

func (app *Application) LoadPlugins(path string, modules *[]Module) {
	all_plugins, err := filepath.Glob(path + "*.so")
	if err != nil {
		panic(err)
	}
	for _, filename := range all_plugins {
		p, err := plugin.Open(filename)
		if err != nil {
			continue
		}

		// Find Init
		symbol, err := p.Lookup("Init")
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
			Init:     initFunc,
			Done:     doneFunc,
			Plugin:   p,
			Filename: filename,
		}
		*modules = append(*modules, *module)
	}
}
