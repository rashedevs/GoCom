package product

import (
	"gocom/util"
	"net/http"
	"strconv"
)

// http://localhost:8080/products/5  ---> 5 is called as request "path parameter"
func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) { // GET and OPTIONS
	productID := r.PathValue("id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.svc.Get(pId)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if product == nil {
		util.SendError(w, http.StatusNotFound, "Product not found")
		return
	}
	util.SendData(w, http.StatusOK, product)
}
