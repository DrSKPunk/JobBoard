package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DrSKPunk/JobBoard/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/vakansii", handlers.GetVakansii).Methods("GET")
	r.HandleFunc("/vakansii", handlers.CreateVakansii).Methods("POST")
	r.HandleFunc("/vakansii/{id}", handlers.GetVakansiiByID).Methods("GET")
	r.HandleFunc("/vakansii/{id}", handlers.UpdateVakansii).Methods("PUT")
	r.HandleFunc("/vakansii/{id}", handlers.DeleteVakansii).Methods("DELETE")

	r.HandleFunc("/vakansii/{id}/otclick", handlers.ApplyForVakansii).Methods("POST")
	r.HandleFunc("/vakansii/{id}/otclick", handlers.GetOtkliki).Methods("GET")

	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
