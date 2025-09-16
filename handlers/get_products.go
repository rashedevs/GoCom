package handlers

import (
	"gocom/database"
	"gocom/util"
	"log"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) { // GET and OPTIONS
	log.Println("ami getproducts handler")
	util.SendData(w, database.ProductList, 200)
}
