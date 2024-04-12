package main

import (
	"subscription/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	type Sub struct {
		Plan  string `json:"plan"`
		Price int    `json:"price"`
	}

	var SubJson []Sub

	sub1 := Sub{Plan: "basico", Price: 0}
	sub2 := Sub{Plan: "intermedio", Price: 19}
	sub3 := Sub{Plan: "premium", Price: 39}

	SubJson = append(SubJson, sub1, sub2, sub3)

	database.Migrate()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Access-Control-Allow-Credentials",
		AllowOrigins: "*",
		// AllowCredentials: true,
		AllowMethods: "GET,POST",
	}))

	app.Use(logger.New())

	app.Get("/subscription/ping", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("ping")
	})

	app.Get("/subscription", func(c *fiber.Ctx) error {

		return c.Status(200).JSON(fiber.Map{"subscription": SubJson})
	})

	app.Listen(":80")

}
