package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/DrSKPunk/JobBoard/internal/models"
	"github.com/DrSKPunk/JobBoard/internal/storage"
	"github.com/gorilla/mux"
)

func otvetJSON(w http.ResponseWriter, status int, data string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(data))
}

func parsJSON(body io.Reader) string {
	bytes, _ := io.ReadAll(body)
	return string(bytes)
}

func GetVakansii(w http.ResponseWriter, r *http.Request) {
	var sb strings.Builder
	sb.WriteString("[")

	for i, v := range storage.Vakansii {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(fmt.Sprintf(
			`{"id_vakansii":%d,"nazvanie":"%s","opisanie":"%s","zarplata":%d,"gorod":"%s","data_publikacii":"%s"}`,
			v.ID, v.Nazvanie, v.Opisanie, v.Zarplata, v.Gorod, v.DataPub,
		))
	}
	sb.WriteString("]")
	otvetJSON(w, http.StatusOK, sb.String())
}

func GetVakansiiByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, v := range storage.Vakansii {
		if v.ID == id {
			res := fmt.Sprintf(
				`{"id_vakansii":%d,"nazvanie":"%s","opisanie":"%s","zarplata":%d,"gorod":"%s","data_publikacii":"%s"}`,
				v.ID, v.Nazvanie, v.Opisanie, v.Zarplata, v.Gorod, v.DataPub,
			)
			otvetJSON(w, http.StatusOK, res)
			return
		}
	}
	http.Error(w, "Вакансия не найдена", http.StatusNotFound)
}

func CreateVakansii(w http.ResponseWriter, r *http.Request) {
	body := parsJSON(r.Body)

	var newVac models.Vakansii
	fmt.Sscanf(body, `{"nazvanie":"%[^"]","opisanie":"%[^"]","zarplata":%d,"gorod":"%[^"]","data_publikacii":"%[^"]"}`,
		&newVac.Nazvanie, &newVac.Opisanie, &newVac.Zarplata, &newVac.Gorod, &newVac.DataPub)

	newVac.ID = storage.NextVakansiiID
	storage.NextVakansiiID++
	storage.Vakansii = append(storage.Vakansii, newVac)

	res := fmt.Sprintf(
		`{"id_vakansii":%d,"nazvanie":"%s","opisanie":"%s","zarplata":%d,"gorod":"%s","data_publikacii":"%s"}`,
		newVac.ID, newVac.Nazvanie, newVac.Opisanie, newVac.Zarplata, newVac.Gorod, newVac.DataPub,
	)
	otvetJSON(w, http.StatusCreated, res)
}

func UpdateVakansii(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	body := parsJSON(r.Body)

	var updatedVac models.Vakansii
	fmt.Sscanf(body, `{"nazvanie":"%[^"]","opisanie":"%[^"]","zarplata":%d,"gorod":"%[^"]","data_publikacii":"%[^"]"}`,
		&updatedVac.Nazvanie, &updatedVac.Opisanie, &updatedVac.Zarplata, &updatedVac.Gorod, &updatedVac.DataPub)

	for i, v := range storage.Vakansii {
		if v.ID == id {
			storage.Vakansii[i].Nazvanie = updatedVac.Nazvanie
			storage.Vakansii[i].Opisanie = updatedVac.Opisanie
			storage.Vakansii[i].Zarplata = updatedVac.Zarplata
			storage.Vakansii[i].Gorod = updatedVac.Gorod
			storage.Vakansii[i].DataPub = updatedVac.DataPub

			res := fmt.Sprintf(
				`{"id_vakansii":%d,"nazvanie":"%s","opisanie":"%s","zarplata":%d,"gorod":"%s","data_publikacii":"%s"}`,
				storage.Vakansii[i].ID, storage.Vakansii[i].Nazvanie, storage.Vakansii[i].Opisanie, storage.Vakansii[i].Zarplata, storage.Vakansii[i].Gorod, storage.Vakansii[i].DataPub,
			)
			otvetJSON(w, http.StatusOK, res)
			return
		}
	}
	http.Error(w, "Вакансия не найдена", http.StatusNotFound)
}

func DeleteVakansii(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for i, v := range storage.Vakansii {
		if v.ID == id {
			storage.Vakansii = append(storage.Vakansii[:i], storage.Vakansii[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Вакансия не найдена", http.StatusNotFound)
}

func ApplyForVakansii(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vacID, _ := strconv.Atoi(vars["id"])
	body := parsJSON(r.Body)

	var newOtk models.Otkliki
	fmt.Sscanf(body, `{"id_rezume":%d,"data_otklika":"%[^"]","status":"%[^"]"}`,
		&newOtk.IDRezume, &newOtk.DataOtk, &newOtk.Status)

	newOtk.ID = storage.NextOtklikID
	newOtk.IDVakansii = vacID
	storage.NextOtklikID++
	storage.Otkliki = append(storage.Otkliki, newOtk)

	res := fmt.Sprintf(
		`{"id_otklika":%d,"id_rezume":%d,"id_vakansii":%d,"data_otklika":"%s","status":"%s"}`,
		newOtk.ID, newOtk.IDRezume, newOtk.IDVakansii, newOtk.DataOtk, newOtk.Status,
	)
	otvetJSON(w, http.StatusCreated, res)
}

func GetOtkliki(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vacID, _ := strconv.Atoi(vars["id"])

	var sb strings.Builder
	sb.WriteString("[")

	first := true
	for _, otk := range storage.Otkliki {
		if otk.IDVakansii == vacID {
			if !first {
				sb.WriteString(",")
			}
			sb.WriteString(fmt.Sprintf(
				`{"id_otklika":%d,"id_rezume":%d,"id_vakansii":%d,"data_otklika":"%s","status":"%s"}`,
				otk.ID, otk.IDRezume, otk.IDVakansii, otk.DataOtk, otk.Status,
			))
			first = false
		}
	}
	sb.WriteString("]")
	otvetJSON(w, http.StatusOK, sb.String())
}
