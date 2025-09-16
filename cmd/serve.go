package cmd

import (
	"fmt"
	"gocom/middleware"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()

	manager.Use(
		middleware.Preflight, //3rd
		middleware.Cors,      //2nd
		middleware.Logger,    //1st
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)
	initRoutes(mux, manager)

	fmt.Println("Server is running on port :8080")

	err := http.ListenAndServe(":8080", wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
