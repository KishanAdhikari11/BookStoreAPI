package handler

import (
	"book_api/pkg/models"
	"book_api/pkg/responses"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var book models.Book

	_, err := book.Delete(h.DB, id)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, book)
	responses.JSON(w, http.StatusOK, "Deleted")
}
