package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Abhishek-Jain-1925/Go-Micro-Services/application"
)


func main(){
	fmt.Println("*** Welcome to Microservices in GO ***")

	app := application.New()
	
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	
	err := app.Start(ctx)
	if err != nil{
		fmt.Println("Something went wrong !! Due to : ",err)
	}
}
