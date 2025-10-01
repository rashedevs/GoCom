package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"log"
	"net/http"
	"strings"
)

func (m *Middlewares) AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		headerArr := strings.Split(header, " ")
		if len(headerArr) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		accssToken := headerArr[1]
		tokenParts := strings.Split(accssToken, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		signature := tokenParts[2]

		message := jwtHeader + "." + jwtPayload

		byteArrSecret := []byte(m.cnf.SecretKey)
		byteArrMessage := []byte(message)

		//HMAC - hash based message authentication code
		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrMessage)
		hash := h.Sum(nil)
		newSignature := base64UrlEncode(hash)
		if newSignature != signature {
			http.Error(w, "Unauthorized, hacker detected", http.StatusUnauthorized)
			return
		}
		log.Println("ami arekta middleware")
		next.ServeHTTP(w, r)
	})
}

func base64UrlEncode(arr []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(arr)
}
