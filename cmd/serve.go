package cmd

import (
	"fmt"
	"gocom/handlers"
	"gocom/util"
	"net/http"
)

func Serve(){
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("GET /products/{productId}", http.HandlerFunc(handlers.GetProductByID))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))

	globalRouter := util.GlobalRouter(mux)
	fmt.Println("Server is running on port :8080")

	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil{
		fmt.Println("Error starting the server", err)
	}
}