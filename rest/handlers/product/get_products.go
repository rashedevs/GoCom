package product

import (
	"gocom/database"
	"gocom/util"
	"log"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("ami getproducts handler")
	util.SendData(w, database.List(), 200)
}
