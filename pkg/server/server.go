package server

import (
	"fmt"
	"guevarez30/mapi/pkg/database"
	"guevarez30/mapi/pkg/image"
	"guevarez30/mapi/pkg/order"

	"github.com/gofiber/fiber/v2"
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

	app.Get("/images/groups", func(c *fiber.Ctx) error {
		user_id := "b5c6379a-ebf9-4845-841b-e187ece03d4d"
		result, err := image.ImageGroupsByUser(db, user_id)
		err = checkers(err)
		if err != nil {
			return err
		}
		return c.JSON(result)
	})

	app.Get("/images/:image_id", func(c *fiber.Ctx) error {
		result, err := image.ImageById(db, c.Params("image_id"))
		err = checkers(err)
		if err != nil {
			return err
		}
		return c.JSON(result)
	})

	app.Get("/orders", func(c *fiber.Ctx) error {
		user_id := "b5c6379a-ebf9-4845-841b-e187ece03d4d"
		result, err := order.OrdersByUser(db, user_id)
		err = checkers(err)
		if err != nil {
			return err
		}
		return c.JSON(result)
	})

	app.Get("/orders/:order_id", func(c *fiber.Ctx) error {
		result, err := order.OrderById(db, c.Params("order_id"))
		err = checkers(err)
		if err != nil {
			return err
		}
		return c.JSON(result)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Div Rhino!")
	})

	app.Listen(":3000")
}
