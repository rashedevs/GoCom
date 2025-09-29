package product

import (
	"gocom/database"
	"gocom/util"
	"net/http"
	"strconv"
)

// http://localhost:8080/products/5  ---> 5 is called as request "path parameter"
func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) { // GET and OPTIONS
	productID := r.PathValue("id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product ID", 400)
		return
	}

	product := database.GetById(pId)
	if product == nil {
		util.SendError(w, "Product not found", 404)
		return
	}
	util.SendData(w, product, 200)
}
