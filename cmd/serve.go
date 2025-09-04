package cmd

import (
	"fmt"
	"gocom/middleware"
	"gocom/router"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()
	manager := middleware.NewManager()
	manager.Use(middleware.Logger, middleware.Hudai)

	initRoutes(mux, manager)

	globalRouter := router.GlobalRouter(mux)
	fmt.Println("Server is running on port :8080")

	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
