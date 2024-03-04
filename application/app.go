package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

// used to store application dependencies
type App struct {
	router http.Handler
	rdb    *redis.Client
}

// constructor
func New() *App {
	app := &App{
		rdb: redis.NewClient(&redis.Options{}),
	}

	app.loadRoutes()

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
