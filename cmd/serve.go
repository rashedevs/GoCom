package cmd

import (
	"gocom/config"
	"gocom/rest"
	"gocom/rest/handlers/product"
	"gocom/rest/handlers/user"
)

func Serve() {
	cnf := config.GetConfigs()
	userHandler := user.NewHandler()
	productHandler := product.NewHandler()

	server := rest.NewServer(userHandler, productHandler)
	server.Start(cnf)
}
