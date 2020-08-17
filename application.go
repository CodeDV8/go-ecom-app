package EComApp

import "github.com/gin-gonic/gin"

type Application struct {
	DB     interface{}
	Hooks  []Hook
	Router *gin.Engine
}

type Hook struct {
}

func NewApplication() *Application {
	app := &Application{}

	return app
}

func (app *Application) Init() {

}

func (app *Application) Run() {

}

func (app *Application) Done() {

}
