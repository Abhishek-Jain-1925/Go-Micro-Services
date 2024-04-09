package main

import (
	"context"
	"fmt"

	"github.com/Abhishek-Jain-1925/Go-Micro-Services/application"
)


func main(){
	fmt.Println("*** Welcome to Microservices in GO ***")

	app := application.New()

	fmt.Println("Server Running...")
	
	err := app.Start(context.Background())
	if err != nil{
		fmt.Println("Something went wrong !! Due to : ",err)
	}
}
