package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/DrSKPunk/JobBoard/internal/models"
	"github.com/DrSKPunk/JobBoard/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

type credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type authResponse struct {
	Token  string `json:"token"`
	UserID int    `json:"user_id"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var creds credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		jsonError(w, http.StatusBadRequest, "некорректный JSON")
		return
	}

	creds.Login = strings.TrimSpace(creds.Login)
	if creds.Login == "" || creds.Password == "" {
		jsonError(w, http.StatusBadRequest, "логин и пароль обязательны")
		return
	}
	if len(creds.Password) < 6 {
		jsonError(w, http.StatusBadRequest, "пароль должен быть не короче 6 символов")
		return
	}

	for _, u := range storage.Users {
		if u.Login == creds.Login {
			jsonError(w, http.StatusConflict, "пользователь с таким логином уже существует")
			return
		}
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "ошибка хеширования пароля")
		return
	}

	user := models.User{
		ID:       storage.NextUserID,
		Login:    creds.Login,
		Password: string(hash),
	}
	storage.NextUserID++
	storage.Users = append(storage.Users, user)

	jsonResponse(w, http.StatusCreated, user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		jsonError(w, http.StatusBadRequest, "некорректный JSON")
		return
	}

	var found *models.User
	for i := range storage.Users {
		if storage.Users[i].Login == creds.Login {
			found = &storage.Users[i]
			break
		}
	}
	if found == nil {
		jsonError(w, http.StatusUnauthorized, "неверный логин или пароль")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(creds.Password)); err != nil {
		jsonError(w, http.StatusUnauthorized, "неверный логин или пароль")
		return
	}

	token, err := generateToken()
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "не удалось сгенерировать токен")
		return
	}
	storage.Tokens[token] = found.ID

	jsonResponse(w, http.StatusOK, authResponse{Token: token, UserID: found.ID})
}

func Me(w http.ResponseWriter, r *http.Request) {
	userID, ok := userIDFromRequest(r)
	if !ok {
		jsonError(w, http.StatusUnauthorized, "требуется авторизация")
		return
	}
	for _, u := range storage.Users {
		if u.ID == userID {
			jsonResponse(w, http.StatusOK, u)
			return
		}
	}
	jsonError(w, http.StatusUnauthorized, "пользователь не найден")
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
	userID, ok := storage.Tokens[token]
	return userID, ok
}
