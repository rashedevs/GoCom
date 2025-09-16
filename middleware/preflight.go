package middleware

import (
	"log"
	"net/http"
)

func Preflight(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("ami preflight hoye jassi")
		// handle preflight request
		if r.Method == http.MethodOptions {
			log.Println("ami preflight handle korlam")
			w.WriteHeader(200)
			return
		}
		next.ServeHTTP(w, r)
	})
}
