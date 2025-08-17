package cmd

import (
	"net/http"

	"github.com/sahasajib/mini_atm/handler"
	"github.com/sahasajib/mini_atm/middleware"
)

func InitRoute(route *http.ServeMux, manager *middleware.Manager){
	route.Handle("GET /atm", manager.With(http.HandlerFunc(handler.GetAllTransactions)))
	// create user
	route.Handle("POST /users", manager.With(http.HandlerFunc(handler.CreateUser)))
	route.Handle("GET /users", manager.With(http.HandlerFunc(handler.AllUser)))
	route.Handle("GET /users/{id}", manager.With(http.HandlerFunc(handler.GetUser)))
	route.Handle("DELETE /users/delete/{id}", manager.With(http.HandlerFunc(handler.DelteUser)))
	route.Handle("PUT /users/update/{id}", manager.With(http.HandlerFunc(handler.UpdateUser)))
	// login user
	route.Handle("POST /user", manager.With(http.HandlerFunc(handler.Login)))
	route.Handle("GET /user/me/balance", manager.With(middleware.JWTMiddleware(http.HandlerFunc(handler.Balance))))
	route.Handle("POST /user/me/deposit", manager.With(middleware.JWTMiddleware(http.HandlerFunc(handler.Deposit))))

	route.Handle("POST /user/me/logout", manager.With(http.HandlerFunc(handler.Logout)))
}