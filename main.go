package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main(){
	fmt.Println("Hello World !!")

	router := chi.NewRouter()

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