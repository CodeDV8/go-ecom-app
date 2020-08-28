package EComApp

func (app *Application) UseBasicAuth(contextName string) *BasicAuth {
	basic := &BasicAuth{
		App:         app,
		ContextName: contextName,
	}
	return basic
}
