package main

import (
	"subscription/database"
	"subscription/middleware"
	"subscription/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
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

	app.Get("/subscription/list", func(c *fiber.Ctx) error {
		type SomeStruct struct {
			Subscription []Sub `json:"subscriptions"`
			Activate     bool  `json:"activate"`
		}

		data := SomeStruct{
			Subscription: SubJson,
			Activate:     false,
		}

		return c.Status(200).JSON(fiber.Map{"subscription": data})
	})

	app.Use(middleware.New(middleware.Config{}))

	app.Post("/subscription", func(c *fiber.Ctx) error {

		user := c.Locals("requestAuth")

		details, _ := user.(middleware.DeserializeUser)

		Subscription := new(model.Subscription)

		if err := c.BodyParser(Subscription); err != nil {
			return err
		}

		Subscription.Id = uuid.New()

		Subscription.UserId = details.ID

		var validate *validator.Validate

		validate = validator.New(validator.WithRequiredStructEnabled())

		err1 := validate.Struct(Subscription)

		if err1 != nil {

			return c.Status(400).JSON(fiber.Map{"message": "field validation error"})

		}

		db := database.DB

		OldSubscription := model.Subscription{UserId: details.ID}

		result_1 := db.Find(&OldSubscription)

		if result_1.RowsAffected > 0 {
			return c.Status(400).JSON(fiber.Map{"message": "ya existe"})
		}

		result := db.Create(&Subscription)

		if result.Error != nil {
			return c.Status(400).JSON(fiber.Map{"message": "error in creation"})
		}

		type SomeSomeStruct struct {
			Subscription model.SubscriptionWithoutId `json:"subscription"`
			Activate     bool                        `json:"activate"`
		}

		data := SomeSomeStruct{
			Subscription: Subscription.SubscriptionWithoutId,
			Activate:     true,
		}

		return c.Status(201).JSON(fiber.Map{"subscription": data})
	})

	app.Listen(":80")

}
