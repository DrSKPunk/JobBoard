package handlers

import "github.com/DrSKPunk/JobBoard/internal/services"

type Handler struct {
	userService        *services.UserService
	vacancyService     *services.VacancyService
	applicationService *services.ApplicationService
}

func NewHandler(us *services.UserService, vs *services.VacancyService, as *services.ApplicationService) *Handler {
	return &Handler{
		userService:        us,
		vacancyService:     vs,
		applicationService: as,
	}
}
