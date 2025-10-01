package cmd

import (
	"gocom/config"
	"gocom/rest"
	"gocom/rest/handlers/product"
	"gocom/rest/handlers/review"
	"gocom/rest/handlers/user"
	middleware "gocom/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfigs()
	middlewares := middleware.NewMiddlewares(cnf)
	userHandler := user.NewHandler()
	productHandler := product.NewHandler(middlewares)
	reviewHandler := review.NewHandler()

	server := rest.NewServer(cnf, userHandler, productHandler, reviewHandler)
	server.Start()
}
