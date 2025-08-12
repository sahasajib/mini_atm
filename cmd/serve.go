package cmd

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/sahasajib/mini_atm/global_routes"
	"github.com/sahasajib/mini_atm/handler"
	
)

func Serve(){
	route := http.NewServeMux()
	route.Handle("GET /atm", http.HandlerFunc(handler.GetAllTransactions))
	route.Handle("POST /users", http.HandlerFunc(handler.CreateUser))
	route.Handle("GET /users", http.HandlerFunc(handler.AllUser))

	globalHandler := global_routes.GlobalRouter(route)
	slog.Info("Starting server on port 8080")

	err := http.ListenAndServe(":8080", globalHandler)
	if err != nil{
		fmt.Println("Error starting server:", err)
	}
}