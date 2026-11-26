package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DrSKPunk/JobBoard/internal/dto"
	"github.com/DrSKPunk/JobBoard/internal/services"
)

var Tokens = map[string]int{}

// Register godoc
// @Summary      Регистрация
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body dto.RegisterRequest true "Данные"
// @Success      201  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]string
// @Router       /auth/register [post]
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

// Login godoc
// @Summary      Авторизация
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input body dto.LoginRequest true "Данные"
// @Success      200  {object}  dto.AuthResponse
// @Failure      401  {object}  map[string]string
// @Router       /auth/login [post]
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

// Me godoc
// @Summary      Профиль
// @Tags         auth
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  dto.UserResponse
// @Router       /auth/me [get]
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
