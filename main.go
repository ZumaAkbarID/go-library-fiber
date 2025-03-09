package main

import "github.com/gofiber/fiber/v2"

func main() {

	app := fiber.New()

	app.Get("/", Developer)

	app.Listen(":3000")

}

func Developer(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("data")
}
