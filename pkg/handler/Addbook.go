package handler

import (
	"book_api/pkg/models"
	"book_api/pkg/responses"
	"encoding/json"
	"net/http"
)

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	// book.Id = rand.Intn(100)
	err := book.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	book.Prepare()
	book, err = book.Save(h.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, book)
}
