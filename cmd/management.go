package cmd

import (
	"log"

	"github.com/dev-hyunsang/my-home-library-server/auth"
	"github.com/gofiber/fiber/v2"
)

func UpdateBookHandler(ctx *fiber.Ctx) error {
	authToken := ctx.GetReqHeaders()["Authorization"][0]
	token, err := auth.ExtractTokenMetadata(authToken)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON("")
	}

	log.Println(token)

	return ctx.Status(fiber.StatusOK).JSON("")
}

func DeleteBookHandler(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON("")
}
