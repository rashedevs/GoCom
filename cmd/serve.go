package cmd

import (
	"fmt"
	"gocom/config"
	"gocom/infra/db"
	"gocom/repo"
	"gocom/rest"
	"gocom/rest/handlers/product"
	"gocom/rest/handlers/user"
	middleware "gocom/rest/middlewares"
	"os"
)

func Serve() {
	cnf := config.GetConfigs()
	dbCon, err := db.NewConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	userRepo := repo.NewUserRepo(dbCon)
	productRepo := repo.NewProductRepo()

	middlewares := middleware.NewMiddlewares(cnf)

	userHandler := user.NewHandler(cnf, userRepo)
	productHandler := product.NewHandler(middlewares, productRepo)

	server := rest.NewServer(cnf, userHandler, productHandler)
	server.Start()
}
