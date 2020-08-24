package EComApp

import (
	"context"
	EComStructs "github.com/codedv8/go-ecom-structs"
	EComStructsAPI "github.com/codedv8/go-ecom-structs/API"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (app *Application) Ping() string {
	return "Pong"
}

func (app *Application) SysInit() {
	// Gin router
	app.Router = gin.New()
	app.Router.Use(gin.Logger())
	app.Router.Use(gin.Recovery())

	// Handle unhandled in router
	app.Router.NoRoute(func(c *gin.Context) {
		payload := &EComStructs.RouterWildcard{Context: c}
		ok, err := app.CallHook("ROUTER_WILDCARD", payload)
		if err != nil {
			// We had an error
			c.String(500, "Something went very wrong when processing the URL")
		} else if ok != false {
			// Since ok is true, we didn't have anyone handling this request
			c.String(404, "Sorry, couldn't find what you're looking for")
		}
	})

	// Load system plugins
	app.LoadPlugins("./system/", &app.SystemModules)

	// Load user plugins
	app.LoadPlugins("./user/", &app.UserModules)

	// Initialize all system modules
	for _, module := range app.SystemModules {
		module.SysInit(app)
	}

	// Initialize all user modules
	for _, module := range app.UserModules {
		module.SysInit(app)
	}

	app.ListenToHook("API_CALL", func(payload interface{}) (bool, error) {
		switch v := payload.(type) {
		case *EComStructsAPI.Root:
			log.Printf("API_CALL in App.Init: %+v\n", v)
			v.I = 42
		default:
			log.Print("Failed to detect struct")
		}
		return true, nil
	})

}

func (app *Application) Init() {
	// Initialize all system modules
	for _, module := range app.SystemModules {
		module.Init(app)
	}

	// Initialize all user modules
	for _, module := range app.UserModules {
		module.Init(app)
	}
}

func (app *Application) Run() {
	// Get the port to use
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	// Create http server
	srv := &http.Server{
		Addr:    ":" + PORT,
		Handler: app.Router,
	}

	// Create channel to handle shutdown
	shutDown := make(chan os.Signal, 1)
	signal.Notify(shutDown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Run the entire show from here
	go func() {
		// app.Router.Run(":" + PORT)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-shutDown

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %+v", err)
	}
	log.Print("Server exited properly")
}

func (app *Application) Done() {
	for _, module := range app.UserModules {
		module.Done(app)
	}
	for _, module := range app.SystemModules {
		module.Done(app)
	}
}
