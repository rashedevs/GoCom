package cmd

import (
	"fmt"
	"gocom/handlers"
	"gocom/middleware"
	"gocom/util"
	"net/http"
)

func Serve(){
	mux := http.NewServeMux()

	mux.Handle("GET /route", middleware.Logger(http.HandlerFunc(http.HandlerFunc(handlers.Test))))

	mux.Handle("GET /products", middleware.Logger(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("GET /products/{productId}", middleware.Logger(http.HandlerFunc(handlers.GetProductByID)))
	mux.Handle("POST /products", middleware.Logger(http.HandlerFunc(handlers.CreateProduct)))

	globalRouter := util.GlobalRouter(mux)
	fmt.Println("Server is running on port :8080")

	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil{
		fmt.Println("Error starting the server", err)
	}
}