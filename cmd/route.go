package cmd

import (
	"net/http"

	"github.com/sahasajib/mini_atm/handler"
	"github.com/sahasajib/mini_atm/middleware"
)

func InitRoute(route *http.ServeMux, manager *middleware.Manager){
	route.Handle("GET /atm", manager.With(http.HandlerFunc(handler.GetAllTransactions)))
	route.Handle("POST /users", manager.With(http.HandlerFunc(handler.CreateUser)))
	route.Handle("GET /users", manager.With(http.HandlerFunc(handler.AllUser)))
	route.Handle("GET /users/{id}", manager.With(http.HandlerFunc(handler.GetUser)))
	route.Handle("DELETE /users/delete/{id}", manager.With(http.HandlerFunc(handler.DelteUser)))
	route.Handle("PUT /users/update/{id}", manager.With(http.HandlerFunc(handler.UpdateUser)))
	route.Handle("POST /user", manager.With(http.HandlerFunc(handler.Login)))
}