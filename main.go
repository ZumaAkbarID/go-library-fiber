package main

import (
	"github.com/ZumaAkbarID/go-library-fiber/internal/config"
	"github.com/ZumaAkbarID/go-library-fiber/internal/connection"
	"github.com/ZumaAkbarID/go-library-fiber/internal/repository"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	customerRepository := repository.NewCustomer(dbConnection)

	app.Get("/", Developer)

	app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)

}

func Developer(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("data")
}
