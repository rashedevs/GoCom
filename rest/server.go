package rest

import (
	"fmt"
	"gocom/config"
	"gocom/rest/handlers/product"
	"gocom/rest/handlers/review"
	"gocom/rest/handlers/user"
	middleware "gocom/rest/middlewares"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	cnf            *config.Config
	userHandler    *user.Handler
	productHandler *product.Handler
	reviewHandler  *review.Handler
}

func NewServer(
	cnf *config.Config,
	userHandler *user.Handler,
	productHandler *product.Handler,
	reviewHandler *review.Handler,
) *Server {
	return &Server{
		cnf:            cnf,
		userHandler:    userHandler,
		productHandler: productHandler,
		reviewHandler:  reviewHandler,
	}
}

func (server *Server) Start() {
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
	server.reviewHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort) // type casting (int to string)
	fmt.Println("Server is running on port", addr)

	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
