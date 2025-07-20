package attendance

import (
	"payslips/internal/business"
	"payslips/internal/entity"
	"payslips/internal/response"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Attendance(c *fiber.Ctx) error
}

type handler struct {
	business business.Business
}

func NewHandler(business business.Business) Handler {
	return &handler{
		business: business,
	}
}

func (h *handler) Attendance(c *fiber.Ctx) error {

	var (
		Entity = "Attendance"
	)

	var payload entity.AttendanceInput
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

	err := h.business.Attendance.Attendance(c.UserContext(), payload)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("Failed to login", err.Error()).
			JSON(c, fiber.StatusBadRequest)
	}

	return response.NewResponse(Entity).
		Success("Attendance successfully", nil).
		JSON(c, fiber.StatusOK)
}
