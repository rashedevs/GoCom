package review

import (
	"gocom/database"
	"gocom/util"
	"log"
	"net/http"
)

func (h *Handler) GetReviews(w http.ResponseWriter, r *http.Request) {
	log.Println("ami getreviews handler")
	util.SendData(w, database.List(), 200)
}
