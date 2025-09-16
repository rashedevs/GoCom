package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler // signature (short form of func)

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

// this is called builder pattern
func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
}

func (mnger *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {
	h := handler
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

// func (mnger *Manager) WrapMux(handler http.Handler, middlewares ...Middleware) http.Handler {
func (mnger *Manager) WrapMux(handler http.Handler) http.Handler {

	h := handler //h = logger(hudai(corsWithPreflight()))
	for _, middleware := range mnger.globalMiddlewares {
		h = middleware(h)
	}

	return h
}
