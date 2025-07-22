package attendance

import (
	"payslips/internal/business"
	"payslips/internal/entity"
	"payslips/internal/response"
	"payslips/pkg/meta"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Attendance(c *fiber.Ctx) error
	ListAttendance(c *fiber.Ctx) error
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
			Errors("Failed to attendance", err.Error()).
			JSON(c, fiber.StatusBadRequest)
	}

	return response.NewResponse(Entity).
		Success("Attendance successfully", nil).
		JSON(c, fiber.StatusOK)
}

func (h *handler) ListAttendance(c *fiber.Ctx) error {

	var (
		Entity = "ListAttendance"
	)

	q := c.Queries()

	m := meta.NewParams(q)

	result, err := h.business.Attendance.List(c.UserContext(), &m)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("Failed to list attendance", err.Error()).
			JSON(c, fiber.StatusBadRequest)
	}

	return response.NewResponse(Entity).
		SuccessWithMeta("List Attendance successfully", result, m).
		JSON(c, fiber.StatusOK)
}
