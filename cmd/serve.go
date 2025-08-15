package cmd

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/sahasajib/mini_atm/global_routes"
	"github.com/sahasajib/mini_atm/middleware"
)

func Serve(){
	manager := middleware.NewManager()
	manager.Use(middleware.Logger)
	route := http.NewServeMux()
	InitRoute(route, manager)

	globalHandler := global_routes.GlobalRouter(route)
	slog.Info("Starting server on port 8080")

	err := http.ListenAndServe(":8080", globalHandler)
	if err != nil{
		fmt.Println("Error starting server:", err)
	}
}