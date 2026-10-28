package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Vakansii struct {
	ID       int    `json:"id_vakansii"`
	Nazvanie string `json:"nazvanie"`
	Opisanie string `json:"opisanie"`
	Zarplata int    `json:"zarplata"`
	Gorod    string `json:"gorod"`
	DataPub  string `json:"data_publikacii"`
}

type Otkliki struct {
	ID         int    `json:"id_otklika"`
	IDRezume   int    `json:"id_rezume"`
	IDVakansii int    `json:"id_vakansii"`
	DataOtk    string `json:"data_otklika"`
	Status     string `json:"status"`
}

var vakansii = []Vakansii{
	{ID: 1, Nazvanie: "Программист", Opisanie: "Нужен программист для программирования", Zarplata: 150000, Gorod: "Барнаул", DataPub: "2026-10-28"},
	{ID: 2, Nazvanie: "Тракторист", Opisanie: "Нужно чистить дорогу", Zarplata: 100000, Gorod: "Сургут", DataPub: "2026-10-27"},
}
var otkliki = []Otkliki{}
var nextVakansiiID = 3
var nextOtklikID = 1

func main() {
	http.HandleFunc("GET /vakansii", getVakansii)
	http.HandleFunc("POST /vakansii", createVakansii)
	http.HandleFunc("GET /vakansii/{id}", getVakansiiByID)
	http.HandleFunc("PUT /vakansii/{id}", updateVakansii)
	http.HandleFunc("DELETE /vakansii/{id}", deleteVakansii)
	http.HandleFunc("POST /vakansii/{id}/otclick", applyForVakansii)
	http.HandleFunc("GET /vakansii/{id}/otclick", getOtkliki)
	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getVakansii(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vakansii)
}

func getVakansiiByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(r.PathValue("id"))

	for _, v := range vakansii {
		if v.ID == id {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	http.Error(w, "Вакансия не найдена", http.StatusNotFound)
}

func createVakansii(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newVac Vakansii

	if err := json.NewDecoder(r.Body).Decode(&newVac); err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	newVac.ID = nextVakansiiID
	nextVakansiiID++
	vakansii = append(vakansii, newVac)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newVac)
}

func updateVakansii(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := strconv.Atoi(r.PathValue("id"))

	var updatedVac Vakansii
	if err := json.NewDecoder(r.Body).Decode(&updatedVac); err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	for i, v := range vakansii {
		if v.ID == id {
			vakansii[i].Nazvanie = updatedVac.Nazvanie
			vakansii[i].Opisanie = updatedVac.Opisanie
			vakansii[i].Zarplata = updatedVac.Zarplata
			vakansii[i].Gorod = updatedVac.Gorod
			vakansii[i].DataPub = updatedVac.DataPub
			json.NewEncoder(w).Encode(vakansii[i])
			return
		}
	}
	http.Error(w, "Вакансия не найдена", http.StatusNotFound)
}

func deleteVakansii(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))

	for i, v := range vakansii {
		if v.ID == id {
			vakansii = append(vakansii[:i], vakansii[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Вакансия не найдена", http.StatusNotFound)
}

func applyForVakansii(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vacID, _ := strconv.Atoi(r.PathValue("id"))

	var newOtk Otkliki
	if err := json.NewDecoder(r.Body).Decode(&newOtk); err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	newOtk.ID = nextOtklikID
	newOtk.IDVakansii = vacID
	nextOtklikID++
	otkliki = append(otkliki, newOtk)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOtk)
}

func getOtkliki(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vacID, _ := strconv.Atoi(r.PathValue("id"))

	var result []Otkliki
	for _, otk := range otkliki {
		if otk.IDVakansii == vacID {
			result = append(result, otk)
		}
	}

	if result == nil {
		result = []Otkliki{}
	}
	json.NewEncoder(w).Encode(result)
}
