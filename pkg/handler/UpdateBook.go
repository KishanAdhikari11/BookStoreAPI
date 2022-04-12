package handler

import (
	"book_api/pkg/models"
	"book_api/pkg/responses"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	books, err := book.Update(h.DB, id)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, books)

}
