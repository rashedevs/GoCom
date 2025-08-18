package cmd

import (
	"gocom/handlers"
	"gocom/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager){
	mux.Handle("GET /rashed", manager.With(
		http.HandlerFunc(handlers.Test),
	))
	mux.Handle("GET /route", manager.With(
		http.HandlerFunc(handlers.Test),
	))
	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(handlers.GetProducts),
		middleware.Arekta,
	))
	mux.Handle("GET /products/{productId}", manager.With(
		http.HandlerFunc(handlers.GetProductByID),
	))
	mux.Handle("POST /products", manager.With(
		http.HandlerFunc(handlers.CreateProduct),
	))
}