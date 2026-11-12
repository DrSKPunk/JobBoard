package handlers

import "net/http"

func AuthRequired(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, ok := userIDFromRequest(r)
		if !ok {
			jsonError(w, http.StatusUnauthorized, "требуется авторизация")
			return
		}
		next(w, r)
	}
}
