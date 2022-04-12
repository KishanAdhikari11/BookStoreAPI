package main

import (
	"book_api/pkg/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	h := handler.NewHandler()

	router := mux.NewRouter()

	router.HandleFunc("/api/books", h.GetAllBooks).Methods("GET")
	router.HandleFunc("/api/books", h.AddBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", h.GetBook).Methods("GET")
	router.HandleFunc("/api/book/{id}", h.UpdateBook).Methods("PUT")

	router.HandleFunc("/api/books/{id}", h.DeleteBook).Methods("DELETE")

	log.Println("API is running")
	http.ListenAndServe(":3002", router)

}
