package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var Tokens = map[string]int{}

type credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type authResponse struct {
	Token  string `json:"token"`
	UserID int    `json:"user_id"`
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var creds credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		jsonError(w, http.StatusBadRequest, "некорректный JSON")
		return
	}

	creds.Login = strings.TrimSpace(creds.Login)
	if creds.Login == "" || len(creds.Password) < 6 {
		jsonError(w, http.StatusBadRequest, "неверные данные (пароль от 6 символов)")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "ошибка хеширования")
		return
	}

	id, err := h.userStorage.CreateUser(r.Context(), creds.Login, string(hash))
	if err != nil {
		jsonError(w, http.StatusConflict, "пользователь уже существует или ошибка БД")
		return
	}

	jsonResponse(w, http.StatusCreated, map[string]interface{}{"id": id, "login": creds.Login})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var creds credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		jsonError(w, http.StatusBadRequest, "некорректный JSON")
		return
	}

	u, err := h.userStorage.GetUserByLogin(r.Context(), creds.Login)
	if err != nil {
		jsonError(w, http.StatusUnauthorized, "неверный логин или пароль")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(creds.Password)); err != nil {
		jsonError(w, http.StatusUnauthorized, "неверный логин или пароль")
		return
	}

	token, _ := generateToken()
	Tokens[token] = u.ID

	jsonResponse(w, http.StatusOK, authResponse{Token: token, UserID: u.ID})
}

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	userID, ok := userIDFromRequest(r)
	if !ok {
		jsonError(w, http.StatusUnauthorized, "требуется авторизация")
		return
	}

	u, err := h.userStorage.GetUserByID(r.Context(), userID)
	if err != nil {
		jsonError(w, http.StatusNotFound, "пользователь не найден")
		return
	}
	jsonResponse(w, http.StatusOK, u)
}

func generateToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
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
