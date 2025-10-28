package product

import (
	"gocom/util"
	"net/http"
	"strconv"
	"sync"
)

var count int64

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	reqQuery := r.URL.Query()
	pageAsString := reqQuery.Get("page")
	limitAsString := reqQuery.Get("limit")

	page, _ := strconv.ParseInt(pageAsString, 10, 32)
	limit, _ := strconv.ParseInt(limitAsString, 10, 32)
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	products, err := h.svc.List(page, limit)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Internal Server Error")
		return
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		count1, err := h.svc.Count()
		if err != nil {
			util.SendError(w, http.StatusBadRequest, "Internal Server Error")
			return
		}
		count = count1
	}()

	wg.Wait()
	util.SendPage(w, products, page, limit, count)
}
