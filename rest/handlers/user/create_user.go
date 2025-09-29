package user

import (
	"encoding/json"
	"fmt"
	"gocom/database"
	"gocom/util"
	"net/http"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}
	createdUser := newUser.Store()
	util.SendData(w, createdUser, http.StatusCreated)
}
