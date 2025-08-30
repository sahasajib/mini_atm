package middleware

import "net/http"

type Middleware func(next http.Handler) http.Handler

type Manager struct{
	globalMiddleware []Middleware
}

func NewManager() *Manager{
	return &Manager{
		globalMiddleware: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware){
	mngr.globalMiddleware = append(mngr.globalMiddleware, middlewares...)
}




func (mngr *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler{
	h := handler

	for _, middleware := range middlewares{
		h = middleware(h)
	}

	return h
}

func (mngr *Manager) WrapWith(handler http.Handler) http.Handler{
	h := handler

	for _, globalMiddleware := range mngr.globalMiddleware{
		h = globalMiddleware(h)
	}
	return h
}