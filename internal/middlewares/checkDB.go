package middlewares

import (
	"cchallenge/internal/bd"
	"net/http"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Database connection lost", http.StatusInternalServerError)
			return
		}
		next(w, r)
	}
}
