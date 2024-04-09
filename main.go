package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)


func main(){
	fmt.Println("Hello World !!")

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/hello", myHandler);

	server := &http.Server{
		Addr: ":1925",
		Handler: router,
	}
	fmt.Println("Server Running on Port 1925...")
	err := server.ListenAndServe()
	if err != nil{
		fmt.Println("error while Running Server !")
	}
}

func myHandler(w http.ResponseWriter, r * http.Request){
	w.Write([]byte("Hello, World !"))
}