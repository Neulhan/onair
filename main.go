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
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New()
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		// c.Locals is added to the *websocket.Conn
		log.Println(c.Locals("allowed"))  // true
		log.Println(c.Params("id"))       // 123
		log.Println(c.Query("v"))         // 1.0
		log.Println(c.Cookies("session")) // ""

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Println("write:", err)
				break
			}
		}
	}))

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
