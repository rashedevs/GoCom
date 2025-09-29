package user

import (
	"encoding/json"
	"fmt"
	"gocom/config"
	"gocom/database"
	"gocom/util"
	"net/http"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
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
	cnf := config.GetConfigs()
	accessToken, err := util.CreateJwt(cnf.SecretKey, util.Payload{
		Sub:         usr.ID,
		FirstName:   usr.FirstName,
		LastName:    usr.LirstName,
		Email:       usr.Email,
		IsShopOwner: usr.IsShopOwner,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	util.SendData(w, accessToken, http.StatusCreated)
}
