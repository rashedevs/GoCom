package handlers

import (
	"gocom/database"
	"gocom/util"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r  *http.Request){  // GET and OPTIONS
	util.SendData(w, database.ProductList, 200)
}

