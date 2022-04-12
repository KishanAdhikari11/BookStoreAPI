package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Id     int    `gorm:"AUTO_INCREMENT" `
	Title  string `json:"title"`
	Author string `json:"Author"`
	Desc   string `json:"desc"`
}

func (b *Book) Validate() error {
	if b.Title == "" {
		return errors.New("title is required")
	}
	if b.Author == "" {
		return errors.New("author is required")
	}
	return nil
}

func (b *Book) Prepare() {
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
}

func (b *Book) Save(db *gorm.DB) (Book, error) {
	err := db.Debug().Create(&b).Error
	if err != nil {
		return Book{}, err
	}
	return *b, nil
}

func (b *Book) FindAll(db *gorm.DB) ([]Book, error) {
	var books []Book
	err := db.Debug().Find(&books).Error
	if err != nil {
		return books, err
	}
	return books, nil
}

func (b *Book) FindById(db *gorm.DB, id int) (Book, error) {
	var book Book
	err := db.Debug().First(&book, id).Error
	if err != nil {
		return book, err
	}

	return book, err
}

func (b *Book) Update(db *gorm.DB, id int) (Book, error) {
	var book = Book{}

	if b.Author != "" {
		book.Author = b.Author
	}

	if b.Title != "" {
		book.Title = b.Title
	}

	if b.Desc != "" {
		book.Desc = b.Desc
	}
	js, _ := json.Marshal(book)
	fmt.Println(string(js))
	err := db.Debug().Model(&Book{}).Where("id=?", id).Updates(book).Error
	if err != nil {
		fmt.Println("Helloooooooooo", err)
		return Book{}, err
	}

	return book, err

}

func (b *Book) Delete(db *gorm.DB, id int) (int64, error) {

	db = db.Debug().Model(&Book{}).Where("id=?", id).Take(&Book{}).Delete(&Book{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
