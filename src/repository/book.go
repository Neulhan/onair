package repository

import (
	"fmt"
	"onair/src/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book model.Book) error
	Update(model.Book) error
	Delete(int) error
	GetAll() ([]model.Book, error)
	GetOne(id int) (model.Book, error)
}

type exampleBookRepository struct {
	db *gorm.DB
}

func (r exampleBookRepository) Create(book model.Book) error {
	return r.db.Create(&book).Error
}

func (r exampleBookRepository) Update(book model.Book) error {
	return r.db.Model(&book).Updates(book).Error
}

func (r exampleBookRepository) Delete(id int) error {
	return r.db.Delete(&model.Book{}, id).Error
}

func (r exampleBookRepository) GetAll() (books []model.Book, err error) {
	err = r.db.Find(&books).Error
	return
}

func (r exampleBookRepository) GetOne(id int) (book model.Book, err error) {
	err = r.db.Find(&book, fmt.Sprintf("id = %d", id)).Error
	return
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &exampleBookRepository{db}
}
