package handlers

import "net/http"

func AuthRequired(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := userIDFromRequest(r)
		if !ok {
			jsonError(w, http.StatusUnauthorized, "требуется авторизация")
			return
		}

		_ = userID
		next(w, r)
	}
}
