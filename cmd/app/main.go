package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DrSKPunk/JobBoard/internal/database"
	"github.com/DrSKPunk/JobBoard/internal/handlers"
	"github.com/DrSKPunk/JobBoard/internal/storage"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Файл .env не найден")
	}

	if err := database.Connect(); err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer database.DB.Close()

	vacancyStorage := storage.NewVacancyStorage(database.DB)
	otklikStorage := storage.NewOtklikiStorage(database.DB)
	userStorage := storage.NewUserStorage(database.DB)

	h := handlers.NewHandler(vacancyStorage, otklikStorage, userStorage)

	r := mux.NewRouter()

	r.HandleFunc("/auth/register", h.Register).Methods("POST")
	r.HandleFunc("/auth/login", h.Login).Methods("POST")
	r.HandleFunc("/auth/me", handlers.AuthRequired(h.Me, userStorage)).Methods("GET")

	r.HandleFunc("/vakansii", h.GetVakansii).Methods("GET")
	r.HandleFunc("/vakansii", handlers.AuthRequired(h.CreateVakansii, userStorage)).Methods("POST")
	r.HandleFunc("/vakansii/{id}", h.GetVakansiiByID).Methods("GET")
	r.HandleFunc("/vakansii/{id}", handlers.AuthRequired(h.UpdateVakansii, userStorage)).Methods("PUT")
	r.HandleFunc("/vakansii/{id}", handlers.AuthRequired(h.DeleteVakansii, userStorage)).Methods("DELETE")

	r.HandleFunc("/vakansii/{id}/otclick", handlers.AuthRequired(h.ApplyForVakansii, userStorage)).Methods("POST")
	r.HandleFunc("/vakansii/{id}/otclick", h.GetOtkliki).Methods("GET")

	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
