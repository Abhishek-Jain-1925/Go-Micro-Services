package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

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
	if err != nil{
		return fmt.Errorf("failed to connect to redis DB : %w",err)
	}
	fmt.Println("Connected to DB...")

	defer func(){
		if err := a.rdb.Close(); err !=nil{
			fmt.Println("Error while closing DB")
		}
	}()

	fmt.Println("Server Running...")
	ch := make(chan error, 1)

	go func(){
		err = server.ListenAndServe()
		if err!=nil{
			ch <- fmt.Errorf("failed to start the server: %w",err) 
		}
		close(ch)
	}()

	select{
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
