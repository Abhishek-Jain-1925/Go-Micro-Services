package handlers

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create an Order")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request){
	fmt.Println("List an Orders")
}

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get an Order by ID")
}

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get an Order by ID")
}

func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get an Order by ID")
}