package getall

import (
	"github.com/gofiber/fiber/v2"
)

func GetAllTaskHandler(usecase GetAllTaskUsecase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		task, err := usecase.Execute(c.Context())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(task)
	}
}
