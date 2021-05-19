package main

import (
	"log"
	"onair/src/config"
	"onair/src/database"
	"onair/src/handler"
	"onair/src/repository"
	"onair/src/usecase"
	"os"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	log.Printf("[%s] START SERVER ON %s", os.Getenv("MODE"), config.GetEnv("PORT"))
	db := database.GetNewConnection(config.DSN, &gorm.Config{})
	bookRepository := repository.NewBookRepository(db)
	bookUseCase := usecase.NewBookUseCase(bookRepository)
	bookHandler := handler.NewBookHandler(bookUseCase)

	bookRouter := app.Group("books")
	{
		bookRouter.Get("", bookHandler.GetAllBooks)
		bookRouter.Get("/:id", bookHandler.GetBook)
		bookRouter.Post("", bookHandler.CreateBook)
		bookRouter.Post("/:id", bookHandler.UpdateBook)
		bookRouter.Delete("/:id", bookHandler.DeleteBook)
	}

	log.Fatal(app.Listen(config.GetEnv("PORT")))
}
