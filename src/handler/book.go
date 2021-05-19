package handler

import (
	"onair/src/model"
	"onair/src/usecase"
	"onair/src/utils"

	"github.com/gofiber/fiber/v2"
)

type BookHandler interface {
	CreateBook(c *fiber.Ctx) error
	UpdateBook(c *fiber.Ctx) error
	DeleteBook(c *fiber.Ctx) error
	GetAllBooks(c *fiber.Ctx) error
	GetBook(c *fiber.Ctx) error
}

type exampleBookHandler struct {
	u usecase.BookUseCase
}

func (h exampleBookHandler) CreateBook(c *fiber.Ctx) error {
	book := new(model.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(500).JSON(utils.NewErrorResponse(err))
	}
	errs := utils.Validate(book)
	if errs != nil {
		return c.Status(400).JSON(utils.NewValidationErrorResponse(errs))
	}
	return c.Status(200).JSON(fiber.Map{
		"err": h.u.CreateBook(*book),
	})
}

func (h exampleBookHandler) UpdateBook(c *fiber.Ctx) error {
	book := new(model.Book)
	if err := c.BodyParser(book); err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{
		"err": h.u.UpdateBook(*book),
	})
}

func (h exampleBookHandler) DeleteBook(c *fiber.Ctx) error {
	i, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{
		"err": h.u.DeleteBook(i),
	})
}

func (h exampleBookHandler) GetAllBooks(c *fiber.Ctx) error {
	books, err := h.u.GetAllBooks()
	return c.Status(200).JSON(fiber.Map{
		"books": books,
		"err":   err,
	})
}

func (h exampleBookHandler) GetBook(c *fiber.Ctx) error {
	i, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	book, err := h.u.GetBook(i)
	return c.Status(200).JSON(fiber.Map{
		"book": book,
		"err":  err,
	})
}

func NewBookHandler(u usecase.BookUseCase) BookHandler {
	return &exampleBookHandler{u}
}
