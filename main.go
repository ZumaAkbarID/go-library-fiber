package main

import (
	"github.com/ZumaAkbarID/go-library-fiber/internal/api"
	"github.com/ZumaAkbarID/go-library-fiber/internal/config"
	"github.com/ZumaAkbarID/go-library-fiber/internal/connection"
	"github.com/ZumaAkbarID/go-library-fiber/internal/repository"
	"github.com/ZumaAkbarID/go-library-fiber/internal/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	customerRepository := repository.NewCustomer(dbConnection)

	customerService := service.NewCustomer(customerRepository)

	app.Get("/", Developer)

	api.NewCustomer(app, customerService)

	app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)

}

func Developer(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON("data")
}
