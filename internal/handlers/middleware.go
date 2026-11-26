package handlers

import (
	"context"
	"net/http"
	"strings"

	"github.com/DrSKPunk/JobBoard/internal/storage"
)

type contextKey string

const UserIDKey contextKey = "userID"

func AuthRequired(next http.HandlerFunc, userStorage *storage.UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := userIDFromRequest(r)
		if !ok {
			jsonError(w, http.StatusUnauthorized, "требуется авторизация")
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next(w, r.WithContext(ctx))
	}
}

func userIDFromRequest(r *http.Request) (int, bool) {
	auth := r.Header.Get("Authorization")
	const prefix = "Bearer "
	if !strings.HasPrefix(auth, prefix) {
		return 0, false
	}
	token := strings.TrimSpace(strings.TrimPrefix(auth, prefix))

	userID, ok := Tokens[token]
	return userID, ok
}
