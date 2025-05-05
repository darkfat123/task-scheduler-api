package create

import (
	"github.com/gofiber/fiber/v2"
)

func CreateTaskHandler(usecase CreateTaskUsecase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := new(CreateRequest)
		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		err := usecase.Execute(c.Context(), *req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(CreateResponse{Message: "Task created successfully"})
	}
}
