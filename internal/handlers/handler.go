package handlers

import "github.com/DrSKPunk/JobBoard/internal/storage"

type Handler struct {
	vacancyStorage *storage.VacancyStorage
	otklikStorage  *storage.OtklikiStorage
	userStorage    *storage.UserStorage
}

func NewHandler(vs *storage.VacancyStorage, os *storage.OtklikiStorage, us *storage.UserStorage) *Handler {
	return &Handler{
		vacancyStorage: vs,
		otklikStorage:  os,
		userStorage:    us,
	}
}
