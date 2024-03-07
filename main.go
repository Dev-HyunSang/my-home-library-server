package main

import (
	"context"
	"github.com/dev-hyunsang/my-home-library-server/db"
	"github.com/dev-hyunsang/my-home-library-server/logger"
	"github.com/dev-hyunsang/my-home-library-server/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	middleware.Middleware(app)

	mysqlClient, err := db.ConnectMySQL()
	if err != nil {
		logger.Error(err.Error())
	}

	defer func() {
		err = mysqlClient.Close()
		if err != nil {
			logger.Error(err.Error())
		}
	}()

	if err := mysqlClient.Schema.Create(context.Background()); err != nil {
		logger.Error(err.Error())
	}

	redisClient := db.ConnectRedis()
	if pong := redisClient.Ping(context.Background()); pong.String() != "ping: PONG" {
		logger.Error(err.Error())
	}

	logger.Info("Successful MySQL Connection")
	logger.Info("Successful Redis Connection")
	logger.Info("Starting Server on http://localhost:3000")

	if err := app.Listen(":3000"); err != nil {
		logger.Error(err.Error())
	}
}
