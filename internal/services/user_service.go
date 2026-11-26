package services

import (
	"context"
	"errors"
	"strings"

	"github.com/DrSKPunk/JobBoard/internal/dto"
	"github.com/DrSKPunk/JobBoard/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

var ErrUnauthorized = errors.New("неверный логин или пароль")

type UserService struct {
	storage *storage.UserStorage
}

func NewUserService(s *storage.UserStorage) *UserService {
	return &UserService{storage: s}
}

func (s *UserService) Register(ctx context.Context, req dto.RegisterRequest) (int, error) {
	req.Login = strings.TrimSpace(req.Login)
	if req.Login == "" || len(req.Password) < 6 {
		return 0, errors.New("неверные данные (пароль от 6 символов)")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	return s.storage.CreateUser(ctx, req.Login, string(hash))
}

func (s *UserService) Login(ctx context.Context, req dto.LoginRequest) (dto.AuthResponse, error) {
	u, err := s.storage.GetUserByLogin(ctx, req.Login)
	if err != nil {
		return dto.AuthResponse{}, ErrUnauthorized
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return dto.AuthResponse{}, ErrUnauthorized
	}

	return dto.AuthResponse{
		Token:  "fake-token-for-now",
		UserID: u.ID,
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, id int) (dto.UserResponse, error) {
	u, err := s.storage.GetUserByID(ctx, id)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return dto.UserResponse{ID: u.ID, Login: u.Login}, nil
}
