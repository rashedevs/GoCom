package rest

import (
	"gocom/rest/handlers"
	middleware "gocom/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("POST /products", manager.With(
		http.HandlerFunc(handlers.CreateProduct),
		middleware.AuthenticateJWT,
	))
	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(handlers.GetProducts),
	))
	mux.Handle("GET /products/{id}", manager.With(
		http.HandlerFunc(handlers.GetProduct),
	))
	mux.Handle("PUT /products/{id}", manager.With(
		http.HandlerFunc(handlers.UpdateProduct),
		middleware.AuthenticateJWT,
	))
	mux.Handle("DELETE /products/{id}", manager.With(
		http.HandlerFunc(handlers.DeleteProduct),
		middleware.AuthenticateJWT,
	))
	mux.Handle("POST /users", manager.With(
		http.HandlerFunc(handlers.CreateUser),
	))
	mux.Handle("POST /users/login", manager.With(
		http.HandlerFunc(handlers.Login),
	))
}
