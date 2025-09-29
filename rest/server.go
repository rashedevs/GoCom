package rest

import (
	"fmt"
	"gocom/config"
	"gocom/rest/handlers/product"
	"gocom/rest/handlers/user"
	middleware "gocom/rest/middlewares"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	userHandler    *user.Handler
	productHandler *product.Handler
}

func NewServer(
	userHandler *user.Handler,
	productHandler *product.Handler,
) *Server {
	return &Server{
		userHandler:    userHandler,
		productHandler: productHandler,
	}
}

func (server *Server) Start(cnf config.Config) {
	manager := middleware.NewManager()

	manager.Use(
		middleware.Preflight, //3rd
		middleware.Cors,      //2nd
		middleware.Logger,    //1st
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	// initRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)
	server.productHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort) // type casting (int to string)
	fmt.Println("Server is running on port", addr)

	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
