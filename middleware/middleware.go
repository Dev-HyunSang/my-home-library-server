package middleware

import (
	"github.com/dev-hyunsang/my-home-library-server/cmd"
	"github.com/dev-hyunsang/my-home-library-server/config"
	"github.com/dev-hyunsang/my-home-library-server/logger"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"strings"
)

func Middleware(app *fiber.App) {
	app.Use(requestid.New())
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger.RequestLogger(),
	}))

	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return config.GetEnv("ENVIRONMENT") == "development"
		},
		AllowOrigins: "http://localhost:8080, http://127.0.0.1:8080",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodPut,
			fiber.MethodDelete,
		}, ""),
		AllowCredentials: true,
	}))

	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/data", cmd.CheckinUserDataHandler)
	auth.Post("/join", cmd.JoinUserHandler)
	auth.Post("/login", cmd.LoginUserHandler)
	auth.Post("/logout", cmd.LogOutUserHandler)

	books := api.Group("/books")
	books.Post("/", cmd.CreateBookHandler)
	books.Get("/", cmd.GetBooksHandler)
	books.Get("/:id", cmd.GetBookHandler)
	books.Post("/:id", cmd.EditBookHandler)
	books.Delete("/:id", cmd.DeleteBookHandler)
}
