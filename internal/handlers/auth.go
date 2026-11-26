package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DrSKPunk/JobBoard/internal/dto"
	"github.com/DrSKPunk/JobBoard/internal/services"
)

var Tokens = map[string]int{}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, http.StatusBadRequest, "некорректный JSON")
		return
	}

	id, err := h.userService.Register(r.Context(), req)
	if err != nil {
		jsonError(w, http.StatusConflict, err.Error())
		return
	}

	jsonResponse(w, http.StatusCreated, map[string]interface{}{"id": id, "login": req.Login})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, http.StatusBadRequest, "некорректный JSON")
		return
	}

	res, err := h.userService.Login(r.Context(), req)
	if err != nil {
		if err == services.ErrUnauthorized {
			jsonError(w, http.StatusUnauthorized, "неверный логин или пароль")
			return
		}
		jsonError(w, http.StatusInternalServerError, "внутренняя ошибка")
		return
	}

	Tokens[res.Token] = res.UserID
	jsonResponse(w, http.StatusOK, res)
}

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(UserIDKey).(int)
	if !ok {
		jsonError(w, http.StatusUnauthorized, "требуется авторизация")
		return
	}

	u, err := h.userService.GetUser(r.Context(), userID)
	if err != nil {
		jsonError(w, http.StatusNotFound, "пользователь не найден")
		return
	}
	jsonResponse(w, http.StatusOK, u)
}
