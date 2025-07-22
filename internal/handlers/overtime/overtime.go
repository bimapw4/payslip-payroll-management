package overtime

import (
	"payslips/internal/business"
	"payslips/internal/entity"
	"payslips/internal/response"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Create(c *fiber.Ctx) error
}

type handler struct {
	business business.Business
}

func NewHandler(business business.Business) Handler {
	return &handler{
		business: business,
	}
}

func (h *handler) Create(c *fiber.Ctx) error {

	var (
		Entity = "Overtime"
	)

	var payload entity.Overtime
	if err := c.BodyParser(&payload); err != nil {
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", err).
			JSON(c, fiber.StatusBadRequest)
	}

	if err := payload.Validation(); err != nil {
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", err).
			JSON(c, fiber.StatusBadRequest)
	}

	re, err := h.business.Overtime.Overtime(c.UserContext(), payload)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("Failed create overtime", err.Error()).
			JSON(c, fiber.StatusBadRequest)
	}

	return response.NewResponse(Entity).
		Success("Create Overtime Successfully", re).
		JSON(c, fiber.StatusOK)
}
