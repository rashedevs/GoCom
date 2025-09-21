package handlers

import (
	"encoding/json"
	"fmt"
	"gocom/database"
	"gocom/util"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) { // POST and OPTIONS
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please send a valid json body", 400)
		return
	}
	createdProduct := database.Store(newProduct)
	util.SendData(w, createdProduct, 201)
}
