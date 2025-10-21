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
	// fmt.Printf("%+v", cnf.DB)
	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	userRepo := repo.NewUserRepo(dbCon)
	productRepo := repo.NewProductRepo(dbCon)

	middlewares := middleware.NewMiddlewares(cnf)

	userHandler := user.NewHandler(cnf, userRepo)
	productHandler := product.NewHandler(middlewares, productRepo)

	server := rest.NewServer(cnf, userHandler, productHandler)
	server.Start()
}
