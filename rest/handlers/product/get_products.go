package product

import (
	"gocom/util"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.svc.List()
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Internal Server Error")
		return
	}

	util.SendData(w, http.StatusOK, products)
}
