package product

import (
	middleware "gocom/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("POST /products", manager.With(
		http.HandlerFunc(h.CreateProduct),
		middleware.AuthenticateJWT,
	))
	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(h.GetProducts),
	))
	mux.Handle("GET /products/{id}", manager.With(
		http.HandlerFunc(h.GetProduct),
	))
	mux.Handle("PUT /products/{id}", manager.With(
		http.HandlerFunc(h.UpdateProduct),
		middleware.AuthenticateJWT,
	))
	mux.Handle("DELETE /products/{id}", manager.With(
		http.HandlerFunc(h.DeleteProduct),
		middleware.AuthenticateJWT,
	))
}
