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
		user_id := "518031f7-bac1-43ba-b5fb-a6045b2e09de"
		result, err := image.ImagesGroupsByUser(db, user_id)
		err = checkers(err)
		if err != nil {
			return err
		}
		return c.JSON(result)
	})

	app.Get("/images/:image_id", func(c *fiber.Ctx) error {
		result, err := image.ImagesById(db, c.Params("image_id"))
		err = checkers(err)
		if err != nil {
			return err
		}
		return c.JSON(result)
	})

	app.Get("/orders", func(c *fiber.Ctx) error {
		user_id := "518031f7-bac1-43ba-b5fb-a6045b2e09de"
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
