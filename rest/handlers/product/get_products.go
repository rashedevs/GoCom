package product

import (
	"gocom/util"
	"log"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("ami getproducts handler")
	products, err := h.productRepo.List()
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Internal Server Error")
		return
	}

	util.SendData(w, http.StatusOK, products)
}
