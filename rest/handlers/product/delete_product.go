package product

import (
	"gocom/database"
	"gocom/util"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product ID", 400)
		return
	}

	database.Delete(pId)
	util.SendData(w, "Product deleted successfully", 200)
}
