package application

import (
	"context"
	"fmt"
	"net/http"
)

// used to store application dependencies
type App struct {
	router http.Handler
}

// constructor
func New() *App {
	app := &App{
		router: loadRoutes(),
	}

	return app
}

// receiver of the method
func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":9000",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
