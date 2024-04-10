package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type App struct{
	router http.Handler
	rdb *redis.Client
}

func New() *App{
	app := &App{
		router: loadRoutes(),
		rdb: redis.NewClient(&redis.Options{}),
	}
	return app
}

func (a *App) Start(ctx context.Context) error{
	server:= &http.Server{
		Addr: ":1925",
		Handler: a.router,
	}

	err := a.rdb.Ping(ctx).Err()

	err = server.ListenAndServe()
	if err!=nil{
		return fmt.Errorf("failed to start the server: %w",err) 
	}
	return nil
}
