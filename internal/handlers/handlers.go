package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DrSKPunk/JobBoard/internal/dto"
	"github.com/gorilla/mux"
)

// GetVakansii godoc
// @Summary      Список вакансий
// @Tags         vacancies
// @Produce      json
// @Success      200  {array}  dto.VacancyResponse
// @Router       /vakansii [get]
func (h *Handler) GetVakansii(w http.ResponseWriter, r *http.Request) {
	res, err := h.vacancyService.GetAll(r.Context())
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Ошибка получения вакансий")
		return
	}
	jsonResponse(w, http.StatusOK, res)
}

// GetVakansiiByID godoc
// @Summary      Вакансия по ID
// @Tags         vacancies
// @Produce      json
// @Param        id path int true "ID"
// @Success      200  {object}  dto.VacancyResponse
// @Router       /vakansii/{id} [get]
func (h *Handler) GetVakansiiByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Неверный ID")
		return
	}

	res, err := h.vacancyService.GetByID(r.Context(), id)
	if err != nil {
		jsonError(w, http.StatusNotFound, "Вакансия не найдена")
		return
	}
	jsonResponse(w, http.StatusOK, res)
}

// CreateVakansii godoc
// @Summary      Создать вакансию
// @Tags         vacancies
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        input body dto.VacancyRequest true "Данные"
// @Success      201  {object}  dto.VacancyResponse
// @Router       /vakansii [post]
func (h *Handler) CreateVakansii(w http.ResponseWriter, r *http.Request) {
	var req dto.VacancyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, http.StatusBadRequest, "Некорректный JSON")
		return
	}

	id, err := h.vacancyService.Create(r.Context(), req)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Не удалось создать")
		return
	}

	res := dto.VacancyResponse{
		ID: id, Title: req.Title, Description: req.Description,
		Salary: req.Salary, City: req.City, PublishedAt: req.PublishedAt,
	}
	jsonResponse(w, http.StatusCreated, res)
}

// UpdateVakansii godoc
// @Summary      Обновить вакансию
// @Tags         vacancies
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID"
// @Param        input body dto.VacancyRequest true "Данные"
// @Success      200  {object}  dto.VacancyResponse
// @Router       /vakansii/{id} [put]
func (h *Handler) UpdateVakansii(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Неверный ID")
		return
	}

	var req dto.VacancyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, http.StatusBadRequest, "Некорректный JSON")
		return
	}

	if err := h.vacancyService.Update(r.Context(), id, req); err != nil {
		jsonError(w, http.StatusInternalServerError, "Не удалось обновить")
		return
	}

	res := dto.VacancyResponse{
		ID: id, Title: req.Title, Description: req.Description,
		Salary: req.Salary, City: req.City, PublishedAt: req.PublishedAt,
	}
	jsonResponse(w, http.StatusOK, res)
}

// DeleteVakansii godoc
// @Summary      Удалить вакансию
// @Tags         vacancies
// @Security     BearerAuth
// @Param        id path int true "ID"
// @Success      204
// @Router       /vakansii/{id} [delete]
func (h *Handler) DeleteVakansii(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Неверный ID")
		return
	}

	if err := h.vacancyService.Delete(r.Context(), id); err != nil {
		jsonError(w, http.StatusInternalServerError, "Не удалось удалить")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// ApplyForVakansii godoc
// @Summary      Оставить отклик
// @Tags         applications
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path int true "ID вакансии"
// @Param        input body dto.ApplicationRequest true "Данные"
// @Success      201  {object}  dto.ApplicationResponse
// @Router       /vakansii/{id}/otclick [post]
func (h *Handler) ApplyForVakansii(w http.ResponseWriter, r *http.Request) {
	vacID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Неверный ID вакансии")
		return
	}

	var req dto.ApplicationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, http.StatusBadRequest, "Некорректный JSON")
		return
	}

	id, err := h.applicationService.Create(r.Context(), vacID, req)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Не удалось создать отклик")
		return
	}

	res := dto.ApplicationResponse{
		ID: id, ResumeID: req.ResumeID, VacancyID: vacID,
	}
	jsonResponse(w, http.StatusCreated, res)
}

// GetOtkliki godoc
// @Summary      Отклики на вакансию
// @Tags         applications
// @Produce      json
// @Param        id path int true "ID вакансии"
// @Success      200  {array}  dto.ApplicationResponse
// @Router       /vakansii/{id}/otclick [get]
func (h *Handler) GetOtkliki(w http.ResponseWriter, r *http.Request) {
	vacID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Неверный ID вакансии")
		return
	}

	res, err := h.applicationService.GetByVacancyID(r.Context(), vacID)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Ошибка получения откликов")
		return
	}
	if res == nil {
		res = []dto.ApplicationResponse{}
	}
	jsonResponse(w, http.StatusOK, res)
}
