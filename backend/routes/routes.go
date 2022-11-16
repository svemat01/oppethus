package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) error {
	err := SetupM1Routes(app)

	return err
}
