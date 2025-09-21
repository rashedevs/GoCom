package rest

import (
	"fmt"
	"gocom/config"
	middleware "gocom/rest/middlewares"
	"net/http"
	"os"
	"strconv"
)

func Server(cnf config.Config) {
	manager := middleware.NewManager()

	manager.Use(
		middleware.Preflight, //3rd
		middleware.Cors,      //2nd
		middleware.Logger,    //1st
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)
	initRoutes(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort) // type casting (int to string)
	fmt.Println("Server is running on port", addr)

	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
