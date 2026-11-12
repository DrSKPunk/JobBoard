package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DrSKPunk/JobBoard/internal/models"
	"github.com/DrSKPunk/JobBoard/internal/storage"
	"github.com/gorilla/mux"
)

func jsonError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func GetVakansii(w http.ResponseWriter, r *http.Request) {
	vakansii, err := storage.GetAllVacancies(r.Context())
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Ошибка получения вакансий из БД")
		return
	}
	jsonResponse(w, http.StatusOK, vakansii)
}

func GetVakansiiByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Неверный ID вакансии")
		return
	}

	v, err := storage.GetVacancyByID(r.Context(), id)
	if err != nil {
		jsonError(w, http.StatusNotFound, "Вакансия не найдена")
		return
	}
	jsonResponse(w, http.StatusOK, v)
}

func CreateVakansii(w http.ResponseWriter, r *http.Request) {
	var v models.Vakansii
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		jsonError(w, http.StatusBadRequest, "Некорректный JSON")
		return
	}

	id, err := storage.CreateVacancy(r.Context(), v)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Не удалось создать вакансию: "+err.Error())
		return
	}
	v.ID = id

	jsonResponse(w, http.StatusCreated, v)
}

func UpdateVakansii(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Неверный ID вакансии")
		return
	}

	var v models.Vakansii
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		jsonError(w, http.StatusBadRequest, "Некорректный JSON")
		return
	}

	if err := storage.UpdateVacancy(r.Context(), id, v); err != nil {
		jsonError(w, http.StatusInternalServerError, "Не удалось обновить вакансию")
		return
	}
	v.ID = id

	jsonResponse(w, http.StatusOK, v)
}

func DeleteVakansii(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Неверный ID вакансии")
		return
	}

	if err := storage.DeleteVacancy(r.Context(), id); err != nil {
		jsonError(w, http.StatusInternalServerError, "Не удалось удалить вакансию")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func ApplyForVakansii(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vacID, err := strconv.Atoi(vars["id"])
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Неверный ID вакансии")
		return
	}

	var o models.Otkliki
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		jsonError(w, http.StatusBadRequest, "Некорректный JSON")
		return
	}
	o.IDVakansii = vacID

	id, err := storage.CreateOtklik(r.Context(), o)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Не удалось создать отклик: "+err.Error())
		return
	}
	o.ID = id

	jsonResponse(w, http.StatusCreated, o)
}

func GetOtkliki(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vacID, err := strconv.Atoi(vars["id"])
	if err != nil {
		jsonError(w, http.StatusBadRequest, "Неверный ID вакансии")
		return
	}

	otkliki, err := storage.GetOtklikiByVacancyID(r.Context(), vacID)
	if err != nil {
		jsonError(w, http.StatusInternalServerError, "Ошибка получения откликов")
		return
	}
	if otkliki == nil {
		otkliki = []models.Otkliki{}
	}
	jsonResponse(w, http.StatusOK, otkliki)
}
