package EComApp

func NewApplication() *Application {
	app := &Application{
		Hooks: make(map[string][]Hook),
	}

	return app
}
