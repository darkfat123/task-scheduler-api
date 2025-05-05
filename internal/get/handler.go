package get

import (
	"github.com/gofiber/fiber/v2"
)

func GetTaskByCodeHandler(usecase GetTaskByCodeUsecase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		code := c.Params("code")

		task, err := usecase.Execute(c.Context(), code)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(task)
	}
}
