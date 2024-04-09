package application

import (
	"net/http"

	"github.com/Abhishek-Jain-1925/Go-Micro-Services/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func loadRoutes() *chi.Mux{
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", HomeHandler)

	router.Route("/orders", loadOrderRoutes)

	return router
}

func HomeHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}

func loadOrderRoutes(router chi.Router){
	orderHandler := &handlers.Order{}

	router.Post("/", orderHandler.Create)
	router.Get("/",orderHandler.List)
	router.Get("/{id}",orderHandler.GetByID)
	router.Put("/{id}", orderHandler.UpdateByID)
	router.Delete("/{id}",orderHandler.DeleteByID)
}