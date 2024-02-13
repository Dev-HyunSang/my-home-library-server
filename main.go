package main

import (
	"context"
	"log"

	"github.com/dev-hyunsang/my-home-library-server/db"
	"github.com/dev-hyunsang/my-home-library-server/middleware"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	middleware.Middleware(app)

	mysqlClient, err := db.ConnectMySQL()
	if err != nil {
		log.Fatalln(err)
	}

	defer mysqlClient.Close()

	if err := mysqlClient.Schema.Create(context.Background()); err != nil {
		log.Fatalln(err)
	}

	redisClient := db.ConnectRedis()
	if pong := redisClient.Ping(context.Background()); pong.String() != "ping: PONG" {
		log.Println(color.RedString("[ERROR]"), "Failed to Redis Connection")
	}

	log.Println(color.GreenString("[GOOD]"), "Successful MySQL Connection")
	log.Println(color.GreenString("[GOOD]"), "Successful Redis Connection")

	if err := app.Listen(":3000"); err != nil {
		log.Fatalln(err)
	}
}
