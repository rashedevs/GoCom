package cmd

import (
	"gocom/config"
	"gocom/repo"
	"gocom/rest"
	"gocom/rest/handlers/product"
	"gocom/rest/handlers/user"
	middleware "gocom/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfigs()

	userRepo := repo.NewUserRepo()
	productRepo := repo.NewProductRepo()

	middlewares := middleware.NewMiddlewares(cnf)

	userHandler := user.NewHandler(cnf, userRepo)
	productHandler := product.NewHandler(middlewares, productRepo)

	server := rest.NewServer(cnf, userHandler, productHandler)
	server.Start()
}
