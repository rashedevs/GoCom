package cmd

import (
	"gocom/config"
	"gocom/rest"
)

func Serve() {
	cnf := config.GetConfigs()

	rest.Server(cnf)
}
