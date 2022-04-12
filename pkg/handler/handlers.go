package handler

import (
	"book_api/db"

	"github.com/jinzhu/gorm"
)

type handler struct {
	DB *gorm.DB
}

func NewHandler() *handler {
	initDB := db.Init()
	return &handler{DB: initDB}
}
