package middleware

import (
	"github.com/dev-hyunsang/my-home-library-server/cmd"
	"github.com/dev-hyunsang/my-home-library-server/logger"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Middleware(app *fiber.App) {
	app.Use(requestid.New())
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger.RequestLogger(),
	}))

	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/join", cmd.JoinUserHandler)
	auth.Post("/login", cmd.LoginUserHandler)
	auth.Post("/logout", cmd.LogOutUserHandler)

	books := api.Group("/books")
	books.Post("/create", cmd.CreateBookHandler)
	books.Get("/", cmd.GetBooksHandler)
	books.Get("/:id", cmd.GetBookHandler)
	books.Delete("/delete/:id", cmd.DeleteBookHandler)
}
