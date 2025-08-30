package rest

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/sahasajib/mini_atm/config"
	"github.com/sahasajib/mini_atm/rest/middleware"
)

func Start(cnf config.Config){
	route := http.NewServeMux()
	manager := middleware.NewManager()

	manager.Use(
		middleware.Cros,
		middleware.Preflight,
		middleware.Logger,
	)
	wrappedRoute := manager.WrapWith(route)
	
	InitRoute(route, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort)
	log.Println("Starting server on port:", addr)

	err := http.ListenAndServe(addr, wrappedRoute)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}