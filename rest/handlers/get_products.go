package handlers

import (
	"gocom/database"
	"gocom/util"
	"log"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("ami getproducts handler")
	util.SendData(w, database.List(), 200)
}
