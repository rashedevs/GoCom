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

func (mnger *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next
	for _, middleware := range middlewares {
		n = middleware(n)
	}

	for _, globalMiddleware := range mnger.globalMiddlewares {
		n = globalMiddleware(n)
	}
	return n
}
