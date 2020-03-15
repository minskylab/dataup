package dataup

import (
	"github.com/gofiber/fiber"
)

func (data *DataService) configAPI() () {
	service := data.app.Group("/service")

	service.Post("/record", func(c *fiber.Ctx) {

	})
}