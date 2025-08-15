package middleware

import "net/http"

type Middleware func(next http.Handler) http.Handler

type Manager struct{
	globalMiddleware []Middleware
}

func (mngr *Manager) Use(middlewares ...Middleware){
	mngr.globalMiddleware = append(mngr.globalMiddleware, middlewares...)
}

func NewManager() *Manager{
	return &Manager{
		globalMiddleware: make([]Middleware, 0),
	}
}


func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler{
	n := next

	for _, middleware := range middlewares{
		n = middleware(n)
	}

	for _, globalMiddleware := range mngr.globalMiddleware{
		n = globalMiddleware(n)
	}
	return n
}