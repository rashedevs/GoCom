package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         int    `json:"sub"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJwt(secret string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	byteArrHeader, err := json.Marshal(header)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	headerBase64 := base64UrlEncode(byteArrHeader)

	byteArrData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	payloadBase64 := base64UrlEncode(byteArrData)
	message := headerBase64 + "." + payloadBase64

	byteArrSecret := []byte(secret)
	byteArrMessage := []byte(message)

	//HMAC - hash based message authentication code
	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)
	signature := h.Sum(nil)
	signatureBase64 := base64UrlEncode(signature)

	jwt := headerBase64 + "." + payloadBase64 + "." + signatureBase64
	return jwt, nil
}

func base64UrlEncode(arr []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(arr)
}
