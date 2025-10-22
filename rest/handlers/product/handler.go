package product

import (
	middleware "gocom/rest/middlewares"
)

type Handler struct {
	middlewares *middleware.Middlewares
	svc         Service
}

func NewHandler(middlewares *middleware.Middlewares, svc Service) *Handler {
	return &Handler{
		middlewares: middlewares,
		svc:         svc,
	}
}
