package EComApp

// UseBasicAuth - Middleware registratio function for using basic authentication
func (app *Application) UseBasicAuth(contextName string) *BasicAuth {
	basic := &BasicAuth{
		App:         app,
		ContextName: contextName,
	}
	return basic
}
