package handlers

import (
	"encoding/json"
	"fmt"
	"gocom/database"
	"gocom/util"
	"net/http"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginReq
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginReq)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}
	usr := database.Find(loginReq.Email, loginReq.Password)
	if usr == nil {
		// http.Error(w, "Invalid Credentials", http.StatusBadRequest)
		util.SendError(w, "Invalid Credentials", http.StatusBadRequest) //testing with sendError
		return
	}
	util.SendData(w, usr, http.StatusCreated)
}
