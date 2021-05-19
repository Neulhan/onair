package usecase

import (
	"onair/src/model"
	"onair/src/repository"
)

type BookUseCase interface {
	CreateBook(book model.Book) error
	UpdateBook(book model.Book) error
	DeleteBook(int) error
	GetAllBooks() ([]model.Book, error)
	GetBook(int) (model.Book, error)
}

type exampleBookUseCase struct {
	r repository.BookRepository
}

func (u exampleBookUseCase) CreateBook(book model.Book) error {
	return u.r.Create(book)
}

func (u exampleBookUseCase) UpdateBook(book model.Book) error {
	return u.r.Update(book)
}

func (u exampleBookUseCase) DeleteBook(id int) error {
	return u.r.Delete(id)
}

func (u exampleBookUseCase) GetAllBooks() ([]model.Book, error) {
	return u.r.GetAll()
}

func (u exampleBookUseCase) GetBook(id int) (model.Book, error) {
	return u.r.GetOne(id)
}

func NewBookUseCase(r repository.BookRepository) BookUseCase {
	return &exampleBookUseCase{r}
}
