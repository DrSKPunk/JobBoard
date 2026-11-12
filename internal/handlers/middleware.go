package handlers

import (
	"context"
	"net/http"

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

		// Кладем ID пользователя в контекст (чтобы хендлеры могли его достать)
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next(w, r.WithContext(ctx))
	}
}
