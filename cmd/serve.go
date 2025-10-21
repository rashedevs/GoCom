package cmd

import (
	"fmt"
	"gocom/config"
	"gocom/infra/db"
	"gocom/repo"
	"gocom/rest"
	prdctHandler "gocom/rest/handlers/product"
	usrHandler "gocom/rest/handlers/user"
	middleware "gocom/rest/middlewares"
	"gocom/user"
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

	//repos
	userRepo := repo.NewUserRepo(dbCon)
	productRepo := repo.NewProductRepo(dbCon)

	//domains
	usrSvc := user.NewService(userRepo)

	middlewares := middleware.NewMiddlewares(cnf)

	userHandler := usrHandler.NewHandler(cnf, usrSvc)
	productHandler := prdctHandler.NewHandler(middlewares, productRepo)

	server := rest.NewServer(cnf, userHandler, productHandler)
	server.Start()
}
