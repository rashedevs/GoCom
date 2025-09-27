package rest

import (
	"gocom/rest/handlers"
	middleware "gocom/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("POST /products", manager.With(
		http.HandlerFunc(handlers.CreateProduct),
	))
	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(handlers.GetProducts),
		middleware.Arekta,
	))
	mux.Handle("GET /products/{id}", manager.With(
		http.HandlerFunc(handlers.GetProduct),
	))
	mux.Handle("PUT /products/{id}", manager.With(
		http.HandlerFunc(handlers.UpdateProduct),
	))
	mux.Handle("DELETE /products/{id}", manager.With(
		http.HandlerFunc(handlers.DeleteProduct),
	))
	mux.Handle("POST /users", manager.With(
		http.HandlerFunc(handlers.CreateUser),
	))
	mux.Handle("POST /users/login", manager.With(
		http.HandlerFunc(handlers.Login),
	))
}
