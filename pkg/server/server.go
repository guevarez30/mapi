package server

import (
	"fmt"
	"guevarez30/mapi/pkg/database"
	"guevarez30/mapi/pkg/image"
	"guevarez30/mapi/pkg/order"
	"guevarez30/mapi/pkg/server/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func checkers(err error) error {
	switch err {
	case nil:
		return nil
	case gorm.ErrRecordNotFound:
		return fiber.NewError(fiber.StatusNotFound)
	default:
		fmt.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError)
	}
}

func Run() {

	db := database.InitDb()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	auth := auth.Initiatlize()

	authProtected := app.Group("/api", auth.Protect)

	authProtected.Get("/images/groups", func(c *fiber.Ctx) error {
		user_id := "b5c6379a-ebf9-4845-841b-e187ece03d4d"
		result, err := image.ImageGroupsByUser(db, user_id)
		err = checkers(err)
		if err != nil {
			return err
		}
		return c.JSON(result)
	})

	authProtected.Get("/images/:image_id", func(c *fiber.Ctx) error {
		result, err := image.ImageById(db, c.Params("image_id"))
		err = checkers(err)
		if err != nil {
			return err
		}
		return c.JSON(result)
	})

	authProtected.Get("/orders", func(c *fiber.Ctx) error {
		user_id := "b5c6379a-ebf9-4845-841b-e187ece03d4d"
		result, err := order.OrdersByUser(db, user_id)
		err = checkers(err)
		if err != nil {
			return err
		}
		return c.JSON(result)
	})

	authProtected.Get("/orders/:order_id", func(c *fiber.Ctx) error {
		result, err := order.OrderById(db, c.Params("order_id"))
		err = checkers(err)
		if err != nil {
			return err
		}
		return c.JSON(result)
	})

	authProtected.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Div Rhino!")
	})

	app.Listen(":3000")
}
