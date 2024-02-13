package middleware

import (
	"github.com/dev-hyunsang/my-home-library-server/cmd"
	"github.com/gofiber/fiber/v2"
)

func Middleware(app *fiber.App) {
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/join", cmd.JoinUserHandler)
	auth.Post("/login", cmd.LoginUserHandler)
	auth.Post("/logout", cmd.LogOutUserHandler)

	books := api.Group("/books")
	books.Post("/update", cmd.UpdateBookHandler)
}
