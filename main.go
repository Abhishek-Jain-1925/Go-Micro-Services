package main

import (
	"fmt"
	"net/http"
)

func main(){
	fmt.Println("Hello World !!")

	server := &http.Server{
		Addr: ":1925",
		Handler: http.HandlerFunc(myHandler),
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